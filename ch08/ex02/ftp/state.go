package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

type state struct {
	name      string
	printType string
	mode      string
	stru      string
	cwd       Cwd
}

func newState() state {
	return state{
		name:      "anonymous",
		printType: "ASCII Non-print",
		mode:      "stream",
		stru:      "file",
		cwd:       newCwd(),
	}
}

func (s *state) handle(conn *ftpConn, req request) response {
	fmt.Printf("%s >>> %s\n", s.name, req.String())

	switch req.command {
	case "USER":
		s.name = req.message
		return newResponse(needPassword)

	case "PASS":
		return newResponse(userLoggedIn)

	case "PORT":
		hostPort := strings.Split(req.message, ",")
		if len(hostPort) != 6 {
			return newResponse(wrongArguments)
		}
		large, err := strconv.Atoi(hostPort[4])
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}
		small, err := strconv.Atoi(hostPort[5])
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}
		host := strings.Join(hostPort[:4], ".")
		port := int64(large*256 + small)
		address := fmt.Sprintf("%s:%d", host, port)
		dataConn, err := net.Dial("tcp", address)
		if err != nil {
			return newResponse(cantOpenConnection)
		}
		conn.dataConn = dataConn
		return newResponse(ok)

	case "LIST", "NLST":
		files, err := s.cwd.Ls(req.message)
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}

		lines := []string{}
		for _, file := range files {
			lines = append(lines, fmt.Sprintf("%s", file.Name()))
		}
		res := newResponse(fileStatusOk)
		res.SetData(strings.Join(lines, "\r\n"))
		return res

	case "CWD":
		err := s.cwd.Cd(req.message)
		if err != nil {
			return newResponse(wrongArguments)
		}
		return newResponse(fileActionOk)

	case "PWD":
		return newResponse(fmt.Sprintf(created, s.cwd.Pwd()))

	case "SIZE":
		return newResponse(notImplementedAtThisSite)

	case "SYST":
		return newResponse(fmt.Sprintf(nameSystemType, "UNIX"))

	case "RETR":
		data, err := s.cwd.Get(req.message)
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}
		res := newResponse(fileStatusOk)
		res.SetData(string(data))
		return res

	case "STOR":
		conn.Reply(*s, newResponse(fileStatusOk))
		data, err := ioutil.ReadAll(conn.dataConn)
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}
		err = s.cwd.Put(req.message, data)
		if err != nil {
			log.Println(err)
			return newResponse(wrongArguments)
		}
		return newResponse(closingDataConnection)

	case "NOOP":
		return newResponse(ok)

	case "QUIT":
		res := newResponse(closingControlConnection)
		res.closing = true
		return res

	case "FEAT", "EPSV", "PASV":
		return newResponse(notImplemented)

	default:
		return newResponse(notImplementedAtThisSite)
	}
}

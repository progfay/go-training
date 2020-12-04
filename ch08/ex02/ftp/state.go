package ftp

import (
	"fmt"
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

	case "LIST":
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
		s.cwd.Cd(req.message)
		return newResponse(fileActionOk)

	case "PWD":
		return newResponse(fmt.Sprintf(created, s.cwd.Pwd()))

	case "SIZE":
		return newResponse(notImplementedAtThisSite)

	case "SYST":
		return newResponse(fmt.Sprintf(nameSystemType, "UNIX"))

		// case "TYPE":
		// case "MODE":
		// case "STRU":
		// case "RETR":
		// case "STOR":
		// case "NOOP":

	case "QUIT":
		return newResponse(closingControlConnection)

	case "FEAT", "EPSV", "PASV":
		return newResponse(notImplemented)

	default:
		return newResponse(notImplementedAtThisSite)
	}
}

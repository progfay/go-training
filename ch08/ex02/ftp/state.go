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
}

func newState() state {
	return state{
		name:      "anonymous",
		printType: "ASCII Non-print",
		mode:      "stream",
		stru:      "file",
	}
}

func (s *state) handle(conn ftpConn, req request) response {
	switch req.command {
	case "USER":
		s.name = req.message
		return newResponse(needPassword)

	case "PASS":
		return newResponse(userLoggedIn)

	// case "QUIT":

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
		return newResponse(ok, "Okay")

	case "LIST":
		return newResponse(fileStatusOk, "hoge", "fuga")

	case "SYST":
		return newResponse(nameSystemType, "UNIX")

	// case "TYPE":
	// case "MODE":
	// case "STRU":
	// case "RETR":
	// case "STOR":
	// case "NOOP":

	case "FEAT", "EPSV", "PASV":
		return newResponse(notImplemented)

	default:
		return newResponse(notImplementedAtThisSite)
	}
}

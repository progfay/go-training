package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleUSER(conn *ftpConn, req request) response {
	conn.state.name = req.message
	return newResponse(needPassword)
}

func handlePASS(conn *ftpConn, req request) response {
	return newResponse(userLoggedIn)
}

func handlePORT(conn *ftpConn, req request) response {
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
}

func handleLIST(conn *ftpConn, req request) response {
	files, err := conn.state.cwd.Ls(req.message)
	if err != nil {
		log.Println(err)
		return newResponse(wrongArguments)
	}

	lines := []string{}
	for _, file := range files {
		fileType := "file"
		if file.IsDir() {
			fileType = "dir "
		}
		lines = append(lines, fmt.Sprintf("%s\t%d\t%s\t%s", file.Mode(), file.Size(), fileType, file.Name()))
	}
	res := newResponse(fileStatusOk)
	res.SetData(strings.Join(lines, "\r\n"))
	return res
}

func handleNLST(conn *ftpConn, req request) response {
	files, err := conn.state.cwd.Ls(req.message)
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
}

func handleCWD(conn *ftpConn, req request) response {
	err := conn.state.cwd.Cd(req.message)
	if err != nil {
		return newResponse(wrongArguments)
	}
	return newResponse(fileActionOk)
}

func handlePWD(conn *ftpConn, req request) response {
	return newResponse(fmt.Sprintf(created, conn.state.cwd.Pwd()))
}

func handleSIZE(conn *ftpConn, req request) response {
	return newResponse(notImplementedAtThisSite)
}

func handleSYST(conn *ftpConn, req request) response {
	return newResponse(fmt.Sprintf(nameSystemType, "UNIX"))
}

func handleRETR(conn *ftpConn, req request) response {
	data, err := conn.state.cwd.Get(req.message)
	if err != nil {
		log.Println(err)
		return newResponse(wrongArguments)
	}
	res := newResponse(fileStatusOk)
	res.SetData(string(data))
	return res
}

func handleSTOR(conn *ftpConn, req request) response {
	conn.Reply(newResponse(fileStatusOk))
	data, err := ioutil.ReadAll(conn.dataConn)
	if err != nil {
		log.Println(err)
		return newResponse(wrongArguments)
	}
	err = conn.state.cwd.Put(req.message, data)
	if err != nil {
		log.Println(err)
		return newResponse(wrongArguments)
	}
	return newResponse(closingDataConnection)
}

func handleNOOP(conn *ftpConn, req request) response {
	return newResponse(ok)
}

func handleQUIT(conn *ftpConn, req request) response {
	res := newResponse(closingControlConnection)
	res.closing = true
	return res
}

func handleFEAT(conn *ftpConn, req request) response {
	return newResponse(notImplemented)
}

func handleEPSV(conn *ftpConn, req request) response {
	return newResponse(notImplemented)
}

func handlePASV(conn *ftpConn, req request) response {
	return newResponse(notImplemented)
}

package ftp

import (
	"fmt"
	"net"
)

type state struct {
	name string
	cwd  Cwd
}

type ftpConn struct {
	ctrlConn net.Conn
	dataConn net.Conn
	state    state
}

func newftpConn(ctrlConn net.Conn) ftpConn {
	return ftpConn{
		ctrlConn: ctrlConn,
		state: state{
			name: "anonymous",
			cwd:  newCwd(),
		},
	}
}

var commandHanderMap = map[string]func(*ftpConn, request) response{
	"USER": handleUSER,
	"PASS": handlePASS,
	"PORT": handlePORT,
	"LIST": handleLIST,
	"NLST": handleNLST,
	"CWD":  handleCWD,
	"PWD":  handlePWD,
	"SIZE": handleSIZE,
	"SYST": handleSYST,
	"RETR": handleRETR,
	"STOR": handleSTOR,
	"NOOP": handleNOOP,
	"QUIT": handleQUIT,
	"FEAT": handleFEAT,
	"EPSV": handleEPSV,
	"PASV": handlePASV,
}

func (conn *ftpConn) handle(req request) response {
	fmt.Printf("%s >>> %s\n", conn.state.name, req.String())
	handler, ok := commandHanderMap[req.command]
	if !ok {
		return newResponse(notImplementedAtThisSite)
	}

	return handler(conn, req)
}

func (conn *ftpConn) Reply(res response) {
	fmt.Printf("%s <<< %s\n", conn.state.name, res.String())
	fmt.Fprintf(conn.ctrlConn, "%s\n", res.message)

	if res.hasData {
		fmt.Fprintf(conn.dataConn, "%s\r\n", res.data)
		conn.dataConn.Close()
		fmt.Printf("%s <<< %s\n", conn.state.name, closingControlConnection)
		fmt.Fprintf(conn.ctrlConn, "%s\n", closingDataConnection)
	}
}

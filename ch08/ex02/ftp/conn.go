package ftp

import (
	"fmt"
	"net"
)

type ftpConn struct {
	ctrlConn net.Conn
	dataConn net.Conn
}

func newftpConn(ctrlConn net.Conn) ftpConn {
	return ftpConn{
		ctrlConn: ctrlConn,
	}
}

func (conn *ftpConn) Write(msg string) {
	fmt.Fprintf(conn.ctrlConn, msg)
}

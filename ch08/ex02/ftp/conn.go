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

func (conn *ftpConn) Reply(s state, res response) {
	fmt.Printf("%s <<< %s\n", s.name, res.String())
	fmt.Fprintf(conn.ctrlConn, "%s\n", res.message)

	if res.hasData {
		fmt.Fprintf(conn.dataConn, "%s\r\n", res.data)
		conn.dataConn.Close()
		fmt.Printf("%s <<< %s\n", s.name, closingControlConnection)
		fmt.Fprintf(conn.ctrlConn, "%s\n", closingDataConnection)
	}
}

package connection

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Connection struct {
	Name string
	conn net.Conn
}

func New(kv string) (*Connection, error) {
	i := strings.IndexRune(kv, '=')
	if i == -1 {
		return nil, fmt.Errorf("invalid Connection format: %q", kv)
	}

	c := Connection{}
	c.Name = kv[:i]
	value := kv[i+1:]

	conn, err := net.Dial("tcp", value)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	return &c, nil
}

func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) Copy(dst io.Writer) {
	if _, err := io.Copy(dst, c.conn); err != nil {
		log.Fatal(err)
	}
}

func (c *Connection) Read(b []byte) (int, error) {
	return c.conn.Read(b)
}

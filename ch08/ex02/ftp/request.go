package ftp

import (
	"fmt"
	"strings"
)

type request struct {
	command string
	message string
}

func parse(text string) request {
	s := strings.SplitN(text, " ", 2)

	switch len(s) {
	case 0:
		return request{}

	case 1:
		return request{
			command: strings.ToUpper(s[0]),
		}

	default:
		return request{
			command: strings.ToUpper(s[0]),
			message: s[1],
		}
	}
}

func (req *request) String() string {
	if req.message == "" {
		return req.command
	}
	return fmt.Sprintf("%s %s", req.command, req.message)
}

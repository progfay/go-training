package counter

import (
	"bufio"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

package counter

import (
	"bufio"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

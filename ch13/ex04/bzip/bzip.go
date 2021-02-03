package bzip

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

func init() {
	_, err := exec.Command("which", "bzip2").Output()
	if err != nil {
		panic(fmt.Sprintf("bzip2 command not found: %v", err))
	}
}

type writer struct {
	mu  sync.Mutex
	cmd *exec.Cmd
	wc  io.WriteCloser
}

func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")
	cmd.Stdout = out
	wc, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cmd.Start()
	return &writer{
		cmd: cmd,
		wc:  wc,
	}
}

func (w *writer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.wc.Write(p)
}

func (w *writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wc.Close()
	return w.cmd.Wait()
}

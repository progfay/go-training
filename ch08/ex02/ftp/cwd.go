package ftp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Cwd struct {
	path string
}

func newCwd() Cwd {
	return Cwd{path: os.Getenv("HOME")}
}

func (cwd *Cwd) Pwd() string {
	return cwd.path
}

func (cwd *Cwd) Cd(path string) error {
	var p string
	if filepath.IsAbs(path) {
		p = path
	} else {
		p = filepath.Join(cwd.path, path)
	}

	info, err := os.Stat(p)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("not directory: %q", p)
	}

	cwd.path = p
	return nil
}

func (cwd *Cwd) Stat(path string) (os.FileInfo, error) {
	if filepath.IsAbs(path) {
		return os.Stat(path)
	}

	return os.Stat(filepath.Join(cwd.path, path))
}

func (cwd *Cwd) Ls(path string) ([]os.FileInfo, error) {
	var p string
	switch {
	case filepath.IsAbs(path):
		p = path

	case path == "":
		p = cwd.path

	default:
		p = filepath.Join(cwd.path, path)
	}

	info, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return []os.FileInfo{info}, nil
	}

	return ioutil.ReadDir(p)
}

func (cwd *Cwd) Get(path string) ([]byte, error) {
	var p string
	if filepath.IsAbs(path) {
		p = path
	} else {
		p = filepath.Join(cwd.path, path)
	}

	info, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("not file: %q", p)
	}

	return ioutil.ReadFile(p)
}

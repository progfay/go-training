package ftp

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Cwd struct {
	path string
}

func newCwd() Cwd {
	return Cwd{path: os.Getenv("HOME")}
}

func (cwd *Cwd) Cd(path string) {
	if filepath.IsAbs(path) {
		cwd.path = path
		return
	}

	cwd.path = filepath.Join(cwd.path, path)
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

	log.Println(p)

	info, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return []os.FileInfo{info}, nil
	}

	return ioutil.ReadDir(p)
}

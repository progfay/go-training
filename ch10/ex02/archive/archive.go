package archive

import (
	"errors"
	"sync"
)

type Loader struct {
	Name string

	Match func(string) bool
	Load  func(string) (interface{}, error)
}

var (
	mu      sync.RWMutex
	loaders []Loader

	ErrUnregisteredFormat = errors.New("unregisterd format")
)

func Register(l Loader) {
	mu.Lock()
	defer mu.Unlock()

	loaders = append(loaders, l)
}

func Load(infile string) (interface{}, error) {
	mu.RLock()
	defer mu.RUnlock()

	for _, loader := range loaders {
		if loader.Match(infile) {
			return loader.Load(infile)
		}
	}

	return nil, ErrUnregisteredFormat
}

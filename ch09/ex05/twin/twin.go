package twin

import (
	"sync/atomic"
)

type unit struct {
	ch chan bool
}

type twin struct {
	running bool
	count   *int64
	left    unit
	right   unit
}

func New() *twin {
	var count int64
	return &twin{
		running: false,
		count:   &count,
		left: unit{
			ch: make(chan bool),
		},
		right: unit{
			ch: make(chan bool),
		},
	}
}

func (t *twin) Run() {
	if t.running {
		return
	}
	t.running = true

	go func() {
		for {
			select {
			case <-t.left.ch:
				atomic.AddInt64(t.count, 1)
				go func() {
					t.right.ch <- false
				}()

			case <-t.right.ch:
				atomic.AddInt64(t.count, 1)
				go func() {
					t.left.ch <- true
				}()
			}
		}
	}()

	t.left.ch <- true
}

func (t *twin) Count() int64 {
	return atomic.LoadInt64(t.count)
}

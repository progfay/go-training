package pipeline

type leaf struct {
	name string
	ch   chan func(string)
}

func newLeaf(name string, ch chan func(string)) *leaf {
	go func() {
		for {
			select {
			case f := <-ch:
				f(name)
			}
		}
	}()

	return &leaf{
		name: name,
		ch:   ch,
	}
}

type Pipeline struct {
	leaves []*leaf
}

func New(len int, naming func(int) string) *Pipeline {
	leaves := make([]*leaf, len)

	for i := 0; i < len; i++ {
		ch := make(chan func(string))
		leaves[i] = newLeaf(naming(i), ch)
	}

	return &Pipeline{
		leaves: leaves,
	}
}

func (p *Pipeline) Send(f func(string)) {
	for _, l := range p.leaves {
		go func(l *leaf) { l.ch <- f }(l)
	}
}

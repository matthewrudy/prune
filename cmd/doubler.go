package main

import (
	"github.com/matthewrudy/prune"
)

func main() {
	prune.Run(&NStream{})
}

type NStream struct {
	i int
}

func (s *NStream) Read() Task {
	if s.i >= 1000 {
		return nil
	}

	s.i = s.i + 1

	return &DoubleTask{
		n: s.i,
	}
}

type DoubleTask struct {
	n int
}

func (t *DoubleTask) Process() int {
	return t.n * 2
}

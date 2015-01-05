package main

import (
	"github.com/matthewrudy/prune"
	"strconv"
)

func main() {
	prune.Run(&NStream{}, &Summer{})
}

type NStream struct {
	i int
}

func (s *NStream) Read() prune.Input {
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

func (t *DoubleTask) Map() int {
	return t.n * 2
}

type Summer struct {
	count int
}

func (s *Summer) Reduce(out chan int) {
	var sum int

	for n := range out {
		sum += n
	}

	s.count = sum
}

func (s *Summer) String() string {
	return "sum: " + strconv.Itoa(s.count)
}

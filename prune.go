package prune

import (
	"fmt"
	"sync"
)

type Stream interface {
	Read() Input
}

type Input interface {
	Map() int
}

type Reducer interface {
	Reduce(chan int)
	String() string
}

func Run(stream Stream, reducer Reducer) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	in := make(chan Input)

	go func() {
		for {
			if task := stream.Read(); task != nil {
				in <- task
			} else {
				break
			}
		}
		close(in)
		wg.Done()
	}()

	out := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			wg.Add(1)
			for task := range in {
				out <- task.Map()
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	reducer.Reduce(out)
	fmt.Println(reducer.String())
}

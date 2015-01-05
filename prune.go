package prune

import (
	"fmt"
	"sync"
)

type TaskStream interface {
	Read() Task
}

type Task interface {
	Process() int
}

func Run(stream TaskStream) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	in := make(chan Task)

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
				out <- task.Process()
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var sum uint64

	for n := range out {
		sum += uint64(n)
	}
	fmt.Println("sum: ", sum)
}

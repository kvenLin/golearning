package main

import (
	"fmt"
	"sync"
)

func doWorker(i int, w worker) {
	for n := range w.in {
		//n, ok := <- c
		//if !ok {
		//	break
		//}

		fmt.Printf("createWorker %d received %c \n", i, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(i, w)
	return w
}

func channelDemo() {
	//1. wait group
	var wg sync.WaitGroup
	wg.Add(20) // 大小写总计20个任务

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//2. wait for all of them
	//for _,worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
	wg.Wait()
}

func main() {
	channelDemo()
}

package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {
	for n := range c {
		//n, ok := <- c
		//if !ok {
		//	break
		//}

		fmt.Printf("createWorker %d received %d \n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //创建带有3个缓冲区的bufferedChannel,对提升性能有一定的优势
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	fmt.Println("Channel as first-class citizen ")
	//channelDemo()
	fmt.Println("Channel as bufferedChannel")
	//bufferedChannel()
	fmt.Println("Channel close and range received")
	channelClose()
}

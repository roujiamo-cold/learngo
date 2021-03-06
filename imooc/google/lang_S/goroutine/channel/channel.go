package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func chanDemo() {
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
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)

}

// 发送方来closechannel
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// channel由发送方close后, 接收方仍会收到数据
	// 接收到的数据为具体类型的零值
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}

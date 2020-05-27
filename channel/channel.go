package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
}

func createWorker(c int) chan<- int {
	ch := make(chan int)
	go func() {
		for {
			n := <-ch
			fmt.Printf("worker %d received %c\n", c, n)
		}
	}()
	return ch
}
func chanDemo() {
	//go worker(c)
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
	go worker1(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int, 3)
	go worker2(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	// close之后会收到空串
	time.Sleep(time.Millisecond)
}

func worker1(c chan int) {
	func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
}

func worker2(c chan int) {
	//func() {
	//	for {
	//		n, ok := <-c
	//		if !ok {
	//			break
	//		}
	//		fmt.Println(n)
	//	}
	//}()
	// range写法
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}

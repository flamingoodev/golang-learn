package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(j int) {
			for {
				// printf是io阻塞性操作
				fmt.Printf("Hello from goroutine %d\n", i)
				//a[j]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	// found 1 data race(s)
	//fmt.Println(a)
}

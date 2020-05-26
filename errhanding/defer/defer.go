package main

import (
	"bufio"
	"fmt"
	fib2 "golang-learn/functional/fib"
	"os"
)

func tryDefer() {
	// 多个defer会存入一个栈中 先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
	panic("error")
	return
}

func writerFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib2.Fibonacci()
	for i := 0; i < 20; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	//tryDefer()
	writerFile("fib.txt")
}

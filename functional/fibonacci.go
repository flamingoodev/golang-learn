package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 使用闭包完成斐波那契数
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 使用闭包完成斐波那契数
func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
func main() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	//
	fib := fib()
	printFileContents(fib)
}

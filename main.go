package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var name = "s"
var address string

// 全局变量可以不使用，局部变量不使用会报错
var address1 string
var name1, name2, name3 string = "1", "2", "3"
var (
	age  = 18
	isOk = true
)

func varZeroValue() {
	var name string = "zhangsan"
	isOk = false
	fmt.Println(name1, name2, name3)
	fmt.Println(name)
}

func euler() {
	// 欧拉公式 e的i pi次方 + 1 = 0
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf(string(c))
}

func constDemo() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(fileName, c)
}

func enumDemo() {
	const (
		cpp = iota
		_
		java
		python
		golang
		js
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, js)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//非全局变量声明后必须使用
	//var address2© string
	varZeroValue()
	fmt.Print(name)
	fmt.Println(age)
	fmt.Println(isOk)
	fmt.Print("Hello world\n")
	fmt.Println(address)
	fmt.Println(address1)
	euler()
	triangle()
	constDemo()
	enumDemo()
}

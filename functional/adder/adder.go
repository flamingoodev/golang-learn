package main

import "fmt"

// 函数式编程实现累加器
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}

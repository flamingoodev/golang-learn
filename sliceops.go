package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap())
}
func main() {
	// 定义slice
	// zero value for slice is nil
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}

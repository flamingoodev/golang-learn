package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice=%d, len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	// 定义slice
	// zero value for slice is nil
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	// 创建slice
	s5 := []int{2, 4, 6, 8}
	printSlice(s5)
	s6 := make([]int, 10)
	printSlice(s6)
	s7 := make([]int, 10, 30)
	printSlice(s7)
	// copy
	fmt.Println("Copying slice")
	copy(s6, s7)
	printSlice(s6)
	// delete
	fmt.Println("Deleting elements form slice")
	s7 = append(s7[:3], s6[4:]...)
	printSlice(s7)
	// pop
	fmt.Println("Popping from front")
	front := s7[0]
	s7 = s7[1:]
	fmt.Printf("pop %d\n", front)
	printSlice(s7)

	fmt.Println("Popping from back")
	back := s7[len(s7)-1]
	s7 = s7[:len(s7)-1]
	fmt.Printf("pop %d\n", back)
	printSlice(s7)
}

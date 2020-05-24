package main

import "fmt"

// 数组是值类型传递，故调用func f(arr [10]int)会拷贝数组（其他语言数组都是引用传递）
// 以上函数可使用指针改写
// 在go语言中一般不直接使用指针，一般使用切片

// []int 数组切片 [5]int 数组
func printArray(arr [5]int) {
	// 改变第0个的值，验证数组的值传递
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 使用指针数组，避免数组拷贝
func printArray1(arr *[5]int) {
	// 改变第0个的值，验证数组的值传递
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	// 空值数组
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// 动态数量
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 二维数组
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// range遍历
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	// range遍历 下划线省略返回值
	for _, v := range arr3 {
		fmt.Println(v)
	}
	// [5]int 和[3]int 会被认为是两个不同的数据类型，以下代码编译不通过
	// printArray(arr2)
	printArray(arr3)
	// arr3的第0个值未变成100，确定是值传递
	fmt.Println(arr3)
	// 指针数组
	printArray1(&arr3)
	fmt.Println(arr3)
}

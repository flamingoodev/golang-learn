package main

import "fmt"

// slice本身是没有数据的，是对底层array的一个视图（view）
// s[i]不可以超越len(s)，slice向后扩展不可以超越底层数组cap(s)
// 切片添加元素时如果超过cap，系统会分配更大的底层数组，并拷贝原来的底层数组
// 由于值传递的关系，必须接收append的返回值
// s = append(s, val)
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 半开半闭
	s := [...]int{2: 6}
	fmt.Println(s)
	// 以下都是slice切片 arr[2:6]是arr的一个视图(view)
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	// 数组取得它的slice：arr[:]
	fmt.Println("arr[:] = ", arr[:])
	// 利用切片修改数组的值
	s1 := arr[2:6]
	s2 := arr[:6]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	updateSlice(s1)
	fmt.Println("After updateSlice(s1)")
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println("After updateSlice(s2)")
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	// slice扩展 slice可以向后扩展，不可以向前扩展。
	fmt.Println("Extending slice")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := arr1[2:6]
	// 以下会报错，数组越界，但是使用切片是可以取到值的
	//fmt.Println(s3[4])
	s4 := s3[3:5]
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(cap(s3))
	s3 = append(s3, 10)
	fmt.Println(s3)
}

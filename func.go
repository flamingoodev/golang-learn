package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go函数
// 返回一个值
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

// 可以有多个返回值
// 带余除法 13/3 = 4...1
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 给返回值命名，调用方法之后，返回值的参数名将会是参数值的命名值
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

// 给返回值赋值，直接return
// 不建议使用，函数体较长的时候，可读性差
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// 多个返回值的函数，可以选择只接受部分返回值
func div3(a, b int) (q int) {
	i, _ := div(a, b)
	return i
}

// 返回error，返回error要放在最后
func div4(a, b int, op string) (q int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 函数式编程，函数的参数也可以是一个函数
func apply(op func(int, int) int, a, b int) int {
	// 根据反射获取函数的指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 重写math.pow方法
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 值传递
func swap(a, b int) {
	b, a = a, b
}

// 利用返回值
func swap2(a, b int) (int, int) {
	return b, a
}

// 值传递
func swap1(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	fmt.Println(eval(10, 5, "*"))
	fmt.Println(eval(10, 5, "/"))
	fmt.Println(eval(10, 5, "+"))
	fmt.Println(eval(10, 5, "-"))
	//fmt.Println(eval(10, 5, "++"))
	fmt.Println(div(13, 3))
	q, r := div1(13, 3)
	i, i2 := div(13, 3)
	fmt.Println(i, i2)
	fmt.Println(q, r)
	fmt.Println(div2(13, 3))
	fmt.Println(div4(13, 3, "++"))
	// 判断返回值的错误
	if result, err := div4(13, 3, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(apply(pow, 13, 3))
	// 匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 13, 3))
	fmt.Println(sum(1, 2, 3, 4))
	// go中参数传递的方式是值传递，位置交换失败
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)
	i3, i4 := swap2(a, b)
	fmt.Println(i3, i4)
	// 指针
	c, d := 3, 4
	swap1(&c, &d)
	fmt.Println(c, d)
}

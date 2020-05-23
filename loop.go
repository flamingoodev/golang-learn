package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("adb")
	}
}
func main() {
	fmt.Println(convertToBin(5))  // 101
	fmt.Println(convertToBin(13)) // 1101
	fmt.Println(convertToBin(0))
	fmt.Println(convertToBin(23333333))
	printFile("abc.txt")
	//forever()
}

package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "abc.txt"

func readFile1() {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}
}

func readFile2() {
	if file, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"

	case score < 90:
		g = "B"

	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	readFile1()
	readFile2()
	grade(10)
	fmt.Println(grade(0),
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100))
}

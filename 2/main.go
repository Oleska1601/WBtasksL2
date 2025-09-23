package main

import "fmt"

func test() (x int) { // x - именованный возвращаемый параметр
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	//fmt.Println(test())
	fmt.Println(anotherTest())
}

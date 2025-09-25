package main

import (
	"fmt"
	"log"
	"wbtasksl2/9/unpack"
)

func main() {
	t := "a4bc2d5e"
	res, err := unpack.UnpackString(t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

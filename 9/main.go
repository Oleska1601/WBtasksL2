package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func unpackString(input string) (string, error) {
	var res strings.Builder
	var char rune
	var numBuilder strings.Builder
	for _, v := range input {
		if v >= '0' && v <= '9' {
			numBuilder.WriteRune(v)
		} else {
			strNum := numBuilder.String()
			if strNum == "" {
				res.WriteString(string(v))
				char = v
				continue
			}
			if char == 0 {
				return "", errors.New("incorrect input")
			}
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return "", fmt.Errorf("strconv.Atoi: %w", err)
			}

			if num == 0 {

			}
			rep := strings.Repeat(string(char), num-1)
			res.WriteString(rep)
			char = v
			numBuilder.Reset()
		}
	}
	strNum := numBuilder.String()
	if strNum != "" {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return "", fmt.Errorf("strconv.Atoi: %w", err)
		}
		if num == 0 {

		}
		rep := strings.Repeat(string(char), num-1)
		res.WriteString(rep)
	}
	return res.String(), nil
}

func main() {
	t1 := "a4bc2d5e"
	t2 := "abcd"
	t7 := "a1bccccdefff3"
	t3 := "45"
	t4 := ""
	t5 := "a4bc2d5e10"
	t6 := "a1"
	//t7 := "a0"
	t8 := "1a0"
	fmt.Println(unpackString(t1))
	fmt.Println(unpackString(t2))
	fmt.Println(unpackString(t3))
	fmt.Println(unpackString(t4))
	fmt.Println(unpackString(t5))
	fmt.Println(unpackString(t6))
	//fmt.Println(unpackString(t7))
	fmt.Println(unpackString(t8))
}

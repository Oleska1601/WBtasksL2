package unpack

import (
	"fmt"
	"strconv"
	"strings"
)

func UnpackString(input string) (string, error) {
	var res []rune

	var numBuilder strings.Builder
	isEscape := false // если был \ -> true
	if len(input) == 0 {
		return "", nil
	}
	if input[0] >= '0' && input[0] <= '9' {
		return "", fmt.Errorf("incorrect input")
	}

	for _, v := range input {
		if v == '\\' {
			if isEscape {
				res = append(res, v)
				isEscape = false
			} else {
				if len(numBuilder.String()) > 0 {
					if num, err := strconv.Atoi(numBuilder.String()); err == nil {
						if num == 0 {
							res = res[:len(res)-1]
						} else {
							rep := strings.Repeat(string(res[len(res)-1]), num-1)
							res = append(res, []rune(rep)...)
						}

						numBuilder.Reset()
					}
				}
				isEscape = true
			}
			continue

		}
		if v >= '0' && v <= '9' && !isEscape {
			numBuilder.WriteRune(v)
		} else {

			strNum := numBuilder.String()
			num, err := strconv.Atoi(strNum)
			if err != nil { // значит значения еще нет и переходим на след элемент
				isEscape = false
				res = append(res, v)
				continue
			}

			if num == 0 {
				res = res[:len(res)-1]
			} else {
				rep := strings.Repeat(string(res[len(res)-1]), num-1)
				res = append(res, []rune(rep)...)
				res = append(res, v)
			}

			numBuilder.Reset()
			isEscape = false

		}
	}

	strNum := numBuilder.String()
	num, err := strconv.Atoi(strNum)
	if err == nil { // значит какое-то значение еще лежит
		if num == 0 {
			res = res[:len(res)-1]
		} else {
			rep := strings.Repeat(string(res[len(res)-1]), num-1)
			res = append(res, []rune(rep)...)
		}
	}
	return string(res), nil
}

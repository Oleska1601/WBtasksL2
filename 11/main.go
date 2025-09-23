package main

import (
	"fmt"
	"slices"
)

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r) //m log m
	return string(r)

}

func countAnagrams(sl []string) map[string][]string {
	if len(sl) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for i := 0; i < len(sl); i++ {
		s := sortString(sl[i])
		m[s] = append(m[s], sl[i])
	}
	result := make(map[string][]string)
	for _, v := range m {
		if len(v) > 1 {
			key := v[0]
			result[key] = v
			slices.Sort(result[key])
		}
	}
	return result
}

func main() {
	sl := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	fmt.Println(countAnagrams(sl))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCombinations("723"))
}

func letterCombinations(digits string) []string {
	var numMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if digits == "" {
		return []string{}
	}
	res := []string{""}
	for _, num := range strings.Split(digits, "") {
		var tmp  []string
		for _, suf := range strings.Split(numMap[num], "") {
			for _, t := range res {
				tmp = append(tmp, t+suf)
			}
		}
		res = tmp
	}
	return res
}

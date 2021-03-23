package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(S))
}

func partitionLabels(S string) []int {
	sl := strings.Split(S, "")
	fmt.Println(sl)
	return []int{}
}

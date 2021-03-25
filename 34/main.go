package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(searchRange(a, 3))
}

func searchRange(nums []int, target int) []int {
	start := -1
	newest := -1
	for index, val := range nums {
		if val == target {
			if start == -1 {
				start = index
			}
			newest = index
		}
	}
	return []int{start, newest}
}

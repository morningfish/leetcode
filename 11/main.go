package main

import "fmt"

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	result := 0
	for start < end {
		result = max((end-start)*min(height[start], height[end]), result)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

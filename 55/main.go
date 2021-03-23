package main

import "fmt"

func main() {
	a := []int{2, 0}
	fmt.Println(canJump(a))
}

func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

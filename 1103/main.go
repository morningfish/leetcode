package main

import (
	"fmt"
	"math"
)

func main() {
	res := distributeCandies(10, 3)
	fmt.Println(res)
}

func distributeCandies(candies int, num_people int) []int {
	var num = 0
	var res = make([]int, num_people)
	for candies > 0 {
		num += 1
		res[num%num_people] += int(math.Min(float64(num), float64(candies)))
		candies -= num
	}
	return res
}

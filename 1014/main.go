package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(a))
}

/*
(A[i] + [j]) + (i-j) => max(A[i] + i) && max(A[j]-j)
找到A[i] + i 和 A[j] - j 的最大值 i < j

j - i
score (A[i] + [j]) + (i-j)
A[i] + i && A[j] -j i > j
*/

func maxScoreSightseeingPair(A []int) int {
	maxScore := math.MinInt32
	var i int = A[0]
	for j := 1; j < len(A); j++ {
		if i+A[j]-j > maxScore {
			maxScore = i + A[j] - j
		}
		if A[j]+j >= i {
			i = A[j] + j
		}
	}
	return maxScore
}

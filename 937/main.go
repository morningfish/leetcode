package main

import "fmt"

func prevPermOpt1(A []int) []int {
	if len(A) == 0 {
		return A
	}
	for i := len(A) - 1; i > 0; i-- {
		if A[i] >= A[i-1] {
			continue
		} else {
			max := 0
			k := 0
			for j := i; j < len(A); j++ {
				if A[j] > max && A[j] < A[i-1] {
					max = A[j]
					k = j
				}
			}
			A[i-1],A[k]=A[k],A[i-1]
		}
		return A
	}
	return A
}
func main() {
	exmp := []int{3, 2, 1}
	res := prevPermOpt1(exmp)
	fmt.Println(res)
}

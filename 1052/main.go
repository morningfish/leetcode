package main

import "fmt"

func main() {
	c := []int{1, 0, 1, 2, 1, 1, 7, 5}
	g := []int{0, 1, 0, 1, 0, 1, 0, 1}
	x := 3
	fmt.Print(maxSatisfied1(c, g, x))
}

func maxSatisfied1(customers []int, grumpy []int, X int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase
	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		maxIncrease = max(maxIncrease, increase)
	}
	return total + maxIncrease
}
func maxSatisfied(customers []int, grumpy []int, X int) int {
	l := len(customers)
	max := 0
	for i := 0; i+X < l; i++ {

	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

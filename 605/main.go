package main

import "fmt"

func main() {
	l := []int{0, 0, 1, 0, 0}
	res := canPlaceFlowers(l, 2)
	fmt.Println(res)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	for i := 1; i <= length; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count += 1
		}
	}
	if count == n {
		return true
	} else {
		return false
	}
}

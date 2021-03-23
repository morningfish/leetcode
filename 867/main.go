package main

import "fmt"

func main() {
	a := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(transpose(a))
}

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	if m == n {
		// 原地转置
		for i := 0; i < n; i++ {
			for j := i + 1; j < m; j++ {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		return matrix
	} else {
		// 构造新数组转置
		ans := make([][]int, 0, n)
		for i := 0; i < n; i++ {
			ans = append(ans, make([]int, 0, m))
			for j := 0; j < m; j++ {
				ans[i] = append(ans[i], matrix[j][i])
			}
		}
		return ans
	}
}

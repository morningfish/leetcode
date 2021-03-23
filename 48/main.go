package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		start := i
		end := len(matrix) - i - 1
		for j := 0; j < end-start; j++ {
			temp := matrix[start][start+j]
			matrix[start][start+j] = matrix[end-j][start]
			matrix[end-j][start] = matrix[end][end-j]
			matrix[end][end-j] = matrix[start+j][end]
			matrix[start+j][end] = temp
		}
	}
}

package main

import (
	"fmt"
)

func createMatrix(rows, columns int) [][] int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i ++ {
		row := make([]int, 0)
		for j := 0; j < columns; j ++ {
			row = append(row, i * j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func main() {
	nByN := createMatrix(10, 10)
	for i := 0; i < len(nByN); i ++ {
		fmt.Println(nByN[i])
	}
}
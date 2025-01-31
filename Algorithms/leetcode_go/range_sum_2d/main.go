package main

import "fmt"

type NumMatrix struct {
	prefix [][]int
}

// Constructor calc prefix for arrays
//
//	The sum of elements from index 0 to i is called the prefix sum
//	prefix = from the beginning, prefix sum = sum from the beginning, i.e. index 0
//	if we have the prefix sum for each index, we can find the sum of any subarray in constant time.
func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return numMatrix
	}

	rows, cols := len(matrix), len(matrix[0])
	numMatrix.prefix = make([][]int, rows+1)
	for i := range numMatrix.prefix {
		numMatrix.prefix[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			numMatrix.prefix[r+1][c+1] = matrix[r][c] +
				numMatrix.prefix[r][c+1] +
				numMatrix.prefix[r+1][c] -
				numMatrix.prefix[r][c]
		}
	}

	return numMatrix
}

// SumRegion calc sums region from matrix
func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.prefix[row2+1][col2+1] -
		nm.prefix[row1][col2+1] -
		nm.prefix[row2+1][col1] +
		nm.prefix[row1][col1]
}

// implement - https://leetcode.com/problems/range-sum-query-2d-immutable
// with prefix sum

func main() {
	matrix := [][]int{{
		3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)

	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
}

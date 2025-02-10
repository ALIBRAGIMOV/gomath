package main

import "fmt"

// https://algo.monster/liteproblems/370
// https://doocs.github.io/leetcode/en/lc/370/

// use prefix-sum

// rangeAddition
func rangeAddition(updates [][]int, length int) []int {
	res := make([]int, length) // [0, 0, 0, 0, 0]

	for _, update := range updates {
		start := update[0]
		end := update[1]
		inc := update[2]

		res[start] += inc

		if (end + 1) < length {
			res[end+1] -= inc
		}
	}

	// prefix sum accumulate
	for i := 1; i < length; i++ {
		res[i] += res[i-1]
	}

	return res
}

func main() {
	length := 5
	updates := [][]int{
		{1, 3, 2},
		{2, 4, 3},
		{0, 2, -2},
	}

	// Input: length = 5, updates = [[1,3,2],[2,4,3],[0,2,-2]]
	// Output: [-2,0,3,5,3]
	fmt.Println(rangeAddition(updates, length))

	length = 10
	updates = [][]int{
		{2, 4, 6},
		{5, 6, 8},
		{1, 9, -4},
	}

	// Input: length = 10, updates = [[2,4,6],[5,6,8],[1,9,-4]]
	// Output: [0,-4,2,2,2,4,4,-4,-4,-4]
	fmt.Println(rangeAddition(updates, length))
}

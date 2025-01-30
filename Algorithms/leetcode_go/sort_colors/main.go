package main

import (
	"fmt"
)

// implementation for https://leetcode.com/problems/sort-colors/description/
//  before [2 0 2 1 1 0] -> [0 0 2 1 1 2]
//	before [0 0 2 1 1 2] -> [0 0 2 1 1 2]
//	before [0 0 2 1 1 2] -> [0 0 2 1 1 2]
//	before [0 0 2 1 1 2] -> [0 0 1 1 2 2]

func sortColors(nums []int) []int {
	l, r := 0, len(nums)-1
	i := 0

	for i <= r {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}

	return nums
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	fmt.Println(sortColors(nums))
}

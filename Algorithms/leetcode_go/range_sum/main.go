package main

import "fmt"

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, 0, len(nums))
	currSum := 0

	for _, num := range nums {
		currSum += num
		prefixSums = append(prefixSums, currSum)
	}

	return NumArray{
		prefixSums: prefixSums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSums[right]
	}

	return this.prefixSums[right] - this.prefixSums[left-1]
}

// implement - https://leetcode.com/problems/range-sum-query-immutable/
// with prefix sum

func main() {
	nums := []int{3, 0, 1, 4, 2}
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(1, 3))
	fmt.Println(obj.SumRange(0, 3))
}

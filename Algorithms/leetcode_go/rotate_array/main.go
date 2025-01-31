package main

// rotate implement - O(n), O(1) - https://leetcode.com/problems/rotate-array/
//  1. [7 6 5 4 3 2 1]
//  2. [5 6 7 4 3 2 1]
//  3. [5 6 7 1 2 3 4]
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)
}

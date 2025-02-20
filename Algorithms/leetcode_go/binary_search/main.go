package main

import "fmt"

// https://leetcode.com/problems/binary-search

// BinarySearch performs a binary search on a sorted slice of integers.
// It returns the index of the target if found, otherwise -1.
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right { // <= here because left and right could point to the same element, < would miss it

		// Prevent potential integer overflow when calculating mid.
		// Using (left + right) / 2 directly can overflow if left and right are large.
		// and analog use `(right - left) / 2` to prevent `left + right` potential overflow
		// Instead, we convert to uint and use >> "bitwise right shift"  for safety.

		// The >> operator in Go is a bitwise right shift operation.
		// It moves the bits of a number to the right by a specified number of positions,
		// effectively dividing the number by 2^n.
		// 	x := 8           // 8 in binary: 1000
		//	y := x >> 1      // Shift right by 1: 0100 (4 in decimal) - 2^1
		//	z := x >> 2      // Shift right by 2: 0010 (2 in decimal) - 2^2

		mid := int(uint(left+right) >> 1)

		switch {
		case nums[mid] == target: // Target found, return index
			return mid
		case nums[mid] < target: // middle less than target, discard left half by making left search boundary `mid + 1`
			left = mid + 1
		default: // middle greater than target, discard right half by making right search boundary `mid - 1`
			right = mid - 1
		}
	}

	// Target not found
	return -1
}

func main() {
	arr := []int{1, 2, 3, 6, 9, 10}

	fmt.Println(binarySearch(arr, 6)) // 3

	arr = []int{
		1, 2, 3, 6, 9, 10, 11, 22, 55, 100, 200,
		500, 600, 1000, 9999,
	}

	fmt.Println(binarySearch(arr, 600)) // 12
}

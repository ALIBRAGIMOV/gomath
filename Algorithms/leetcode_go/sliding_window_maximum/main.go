package main

import (
	"fmt"
)

// maxSlidingWindow returns the maximum value in each sliding window of size k.
//
//	using monotonic deque pattern - https://algo.monster/problems/sliding_window_maximum
//	deque operations have in https://pkg.go.dev/github.com/gammazero/deque
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	deque := make([]int, 0, k) // Stores indices of elements in decreasing order
	result := make([]int, 0, len(nums)-k+1)

	for i, num := range nums {
		// Remove elements from the back of deque if they are smaller than the current element
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1] // Pop from back
		}

		// Add the current element index to the deque
		deque = append(deque, i)

		// Remove the front element if it's outside the current window
		if deque[0] == i-k {
			deque = deque[1:] // Pop from front
		}

		// Start adding results when the first window of size k is complete
		if i >= k-1 {
			result = append(result, nums[deque[0]]) // The front of the deque holds the max element index
		}
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{-1485, -1486, -4530, -1636, -2088, 1295}, 3))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

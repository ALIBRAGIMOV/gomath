```go
package main

import "fmt"

// singleNonDuplicate finds the single element in a sorted array where every other element appears exactly twice.
// Binary Search Approach: O(log n)
// Best understanding comments included for visualizing execution clearly.
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		// Calculate mid index safely to avoid integer overflow
		mid := left + (right-left)/2

		// Adjust mid to always point to the first element of the pair (even index)
		if mid%2 == 1 {
			mid--
		}

		// Visualize clearly:
		// If nums[mid] equals nums[mid+1], the unique element is further right.
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			// Otherwise, unique element is at mid or further left.
			right = mid
		}
	}

	// When left == right, we've isolated the unique element.
	return nums[left]
}

func main() {
	tests := [][]int{
		{1, 1, 2, 3, 3, 4, 4, 8, 8},  // expected result: 2
		{3, 3, 7, 7, 10, 11, 11},      // expected result: 10
	}

	for _, test := range tests {
		fmt.Printf("The single non-duplicate element in %v is %d\n", test, singleNonDuplicate(test))
	}
}
```

### Detailed Execution Visualization:

For example `[1,1,2,3,3,4,4,8,8]`:

- **Initial state**: `left = 0`, `right = 8`
- **First iteration**:
  - `mid = 4` (value = 3), adjusted to even index if odd.
  - `nums[mid] (3)` != `nums[mid+1] (4)` ⇒ unique element is at mid or to the left.
  - Update: `right = mid`
- **Second iteration**:
  - `mid = 2` (value = 2), `nums[mid]` != `nums[mid+1]` ⇒ unique element is at mid or left.
  - Update: `right = mid`
- **Third iteration**:
  - `mid = 0`, `nums[mid] (1)` == `nums[mid+1] (1)` ⇒ unique element is right.
  - Update: `left = mid + 2`
- Now `left == right`, pointing to `nums[2] = 2`, the unique element.

This clear step-by-step visualization ensures a deep understanding and clear mental model for solving similar binary search problems.

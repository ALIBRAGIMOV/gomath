## Optimized Product Except Self Solution in Go (Best Practices)

### 1. Problem Statement
Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all elements in `nums` **except `nums[i]`**.

#### 🔹 Constraints:
- Runs in **O(n) time**.
- **No division operation** allowed.
- The product of any prefix or suffix fits in a **32-bit integer**.

#### 🔹 Best Practices:
- **Avoid extra memory allocation** by reusing the `answer` array.
- **Use two passes** (left and right accumulations) to achieve O(n) complexity.
- **Leverage prefix & suffix products** instead of division.
- **Ensure correct initialization** for edge cases (e.g., single-element arrays).

---

### 2. Optimized Solution

```go
package main

import "fmt"

// productExceptSelf returns an array where each index i contains the product of all elements except nums[i].
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil // Handle edge case
	}

	answer := make([]int, n)
	answer[0] = 1 // Initialize left product

	// Compute left products
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	right := 1 // Accumulate right products
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Product Except Self:", productExceptSelf(nums))
}
```

---

### 3. Complexity Analysis
| Operation      | Time Complexity |
|---------------|----------------|
| **Compute left products** | O(n) |
| **Compute right products** | O(n) |
| **Overall Complexity** | O(n) |

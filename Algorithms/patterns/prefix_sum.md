## Prefix Sum Pattern in Go (Best Practices)

### 1. Overview
The **Prefix Sum Pattern** (or **Cumulative Sum**) is a common technique in algorithmic problem-solving. It allows efficient range sum queries by precomputing cumulative sums, reducing the complexity of each query from **O(n) to O(1)** after **O(n) preprocessing**.

#### 🔹 Best Practices:
- **Use slices over arrays** for dynamic sizing.
- **Precompute prefix sums** in a separate slice to avoid redundant calculations.
- **Use zero-based indexing** for consistency.
- **Avoid unnecessary allocations** by using `make([]int, n+1)`.
- **Use const for fixed values** to enhance readability and performance.

---

### 2. Implementation

```go
// subarraySum - small template ofr prefixSums
func subarraySum(arr []int, target int) []int {
    prefixSums := map[int]int{0: 0}
    currentSum := 0

    for i, num := range arr {
        currentSum += num

        if idx, exists := prefixSums[currentSum-target]; exists {
            return []int{idx, i+1}
        }

        if _, exists := prefixSums[currentSum]; !exists {
            prefixSums[currentSum] = i+1
        }
    }

    return nil
}
```

```go
package main

import (
	"fmt"
)

// ComputePrefixSum computes the prefix sum for a given array.
func ComputePrefixSum(arr []int) []int {
	n := len(arr)
	prefixSum := make([]int, n+1) // Extra space for easier calculations

	// Compute prefix sums
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}

	return prefixSum
}

// RangeSum computes the sum of elements in the range [left, right] efficiently.
func RangeSum(prefixSum []int, left, right int) int {
	if left > right {
		return 0 // Edge case: invalid range
	}
	return prefixSum[right+1] - prefixSum[left]
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	prefixSum := ComputePrefixSum(arr)

	// Query example: sum of elements from index 2 to 5
	left, right := 2, 5
	fmt.Printf("Sum of range [%d, %d]: %d\n", left, right, RangeSum(prefixSum, left, right))
}
```

---

### 3. Complexity Analysis
| Operation     | Time Complexity |
|--------------|---------------|
| **Precompute Prefix Sum** | O(n) |
| **Range Query** | O(1) |

---

### 4. Practical Use Cases
- **Subarray Sum Queries**
- **Sliding Window Optimization**
- **2D Grid Prefix Sum (Cumulative Frequency Counting)**
- **Difference Array for Range Updates**
- **Dynamic Programming Optimization**

---

### 5. Key Takeaways
✅ **Use `make([]int, n+1)` to avoid index confusion**
✅ **Precompute prefix sums to speed up range queries**
✅ **Use `prefixSum[right+1] - prefixSum[left]` for O(1) queries**
✅ **Ensure edge case handling (e.g., `left > right`)**

---

**🔥 Expert Tip:** In high-load applications, using prefix sums can significantly reduce computational overhead by replacing multiple iterations with a single O(1) lookup.

🚀 **Next Step:** Try extending this concept to **2D prefix sums** for grid-based problems!

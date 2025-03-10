## Optimized Subarray Sum Solution in Go

### 1. Problem Statement
Given an array of integers `nums` and an integer `k`, return the total number of **contiguous subarrays** whose sum equals `k`.

---

### 2. Optimized Solution

```go
package main

import "fmt"

// subarraySum computes the number of subarrays that sum up to target.
func subarraySum(arr []int, target int) int {
	prefixSum := make(map[int]int, len(arr)) // Optimized map allocation
	prefixSum[0] = 1 // Base case: empty subarray sum
	
	currSum, count := 0, 0
	
	for _, num := range arr {
		currSum += num
		if occurrences, exists := prefixSum[currSum-target]; exists {
			count += occurrences
		}
		prefixSum[currSum]++
	}

	return count
}

func main() {
	arr := []int{1, 2, 3, -3, 1, 2, 3}
	target := 3
	fmt.Printf("Total subarrays with sum %d: %d\n", target, subarraySum(arr, target))
}
```

---

### 3. Complexity Analysis
| Operation      | Time Complexity |
|---------------|----------------|
| **Iterate over array** | O(n) |
| **Map lookups and updates** | O(1) per operation |
| **Overall Complexity** | O(n) |

---

### 4. Improvements from Initial Code
✅ **Optimized map allocation (`make(map[int]int, len(arr))`)** to avoid unnecessary growth.
✅ **Safer map lookup (`if occurrences, exists := prefixSum[currSum-target]; exists`)** to avoid zero-value issues.
✅ **Code clarity improvements** following **Uber Go Style Guide**.
✅ **Ensured correct initialization of `prefixSum[0] = 1`** for handling subarrays starting at index `0`.

---

### 5. Key Takeaways
🚀 **Using a prefix sum hashmap transforms a naive O(n²) solution into an efficient O(n) approach.**

🔹 **This technique is widely used in**:
- **Sliding Window Problems**
- **Dynamic Programming**
- **Counting Prefix-Based Patterns in High-Load Systems**

🔥 **Next Step:** Try extending this to **subarrays with at most K distinct elements** for advanced optimization!

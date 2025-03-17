## Minimum Window Substring in Go (Best Practices)

### 1. Problem Statement
Given two strings, `original` and `check`, return the **minimum substring** of `original` such that **each character in `check` (including duplicates)** is included in this substring.

#### 🔹 Constraints:
- If multiple substrings have the same length, return the **lexicographically smallest** one.
- `1 ≤ len(check), len(original) ≤ 10⁵`
- Only uppercase and lowercase English letters.
- **Case-sensitive** comparison.

#### 🔹 Best Practices:
- **Use the sliding window technique** for O(n) complexity.
- **Use a frequency map** to track occurrences of `check` characters.
- **Dynamically shrink the window** while ensuring it still contains all required characters.
- **Efficient map lookups** using `map[rune]int`.

---

### 2. Optimized Solution

```go
package main

import (
	"fmt"
	"strings"
)

// minWindowSubstring finds the smallest substring of `original` that contains all characters of `check`.
func minWindowSubstring(original, check string) string {
	if len(original) < len(check) {
		return ""
	}

	req := make(map[rune]int) // Frequency map for check
	for _, ch := range check {
		req[ch]++
	}

	left, right, matched := 0, 0, 0
	minStart, minLen := 0, len(original)+1
	window := make(map[rune]int)

	for right < len(original) {
		charRight := rune(original[right])
		window[charRight]++

		if req[charRight] > 0 && window[charRight] <= req[charRight] {
			matched++
		}

		// Try to shrink window while it still contains all characters
		for matched == len(check) {
			// Update minimum window if necessary
			if right-left+1 < minLen || (right-left+1 == minLen && original[left:right+1] < original[minStart:minStart+minLen]) {
				minStart, minLen = left, right-left+1
			}

			charLeft := rune(original[left])
			window[charLeft]--

			if req[charLeft] > 0 && window[charLeft] < req[charLeft] {
				matched--
			}

			left++
		}

		right++
	}

	if minLen > len(original) {
		return ""
	}
	return original[minStart : minStart+minLen]
}

func main() {
	original := "cdbaebaecd"
	check := "abc"
	fmt.Println("Smallest substring containing all characters:", minWindowSubstring(original, check))
}
```

---

### 3. Complexity Analysis
| Operation      | Time Complexity |
|---------------|----------------|
| **Sliding window traversal** | O(n) |
| **Map operations (lookup/update)** | O(1) avg |
| **Overall Complexity** | O(n) |

---
### 4. Interview Expectations & Time Limits

The expected time to solve this problem varies by interview level:

| Level         | Expected Time (minutes) | Breakdown |
|--------------|------------------------|-----------|
| **Junior (L3 - Entry Level)** | 40 - 50 mins | More hints expected, emphasis on understanding sliding window |
| **Mid-Level (L4 - Software Engineer)** | 25 - 35 mins | Expected to implement sliding window efficiently with minor debugging |
| **Senior (L5 - Senior Engineer)** | 20 - 30 mins | Should recognize the pattern quickly, write bug-free code with edge case handling |
| **Staff/Principal (L6/L7)** | 15 - 25 mins | Expected to solve fast, optimize for readability & performance |

### **Recommended Breakdown for a Strong Interview Approach**
- **0 - 5 mins:** Clarify the problem, constraints, and edge cases.
- **5 - 10 mins:** Identify brute force (O(n²)) and optimize to sliding window (O(n)).
- **10 - 15 mins:** Implement the solution, ensuring correctness.
- **15 - 20 mins:** Optimize, check for edge cases (`""`, `aaaa`, `aabb`).
- **20 - 25 mins:** Discuss improvements, trade-offs, and test cases.

🔥 **At top-tier companies (FAANG, Big Tech, Hedge Funds), Staff+ engineers are expected to solve it in ~15-20 minutes** with full correctness and clarity.

# **Efficient Binary Search in Go: Understanding `mid - 1` and `mid + 1`**

## **Overview**
Binary search is a fundamental algorithm used to efficiently locate a target value in a sorted dataset. This document explains why we adjust the search boundaries with `mid - 1` and `mid + 1` and provides an optimized Go implementation following best practices.

## **Understanding the Boundary Adjustments**
### **Problem Statement**
Given a number picked from `1` to `n`, we need to guess the number using a predefined function:

```go
func guess(num int) int
```

- Returns `0` if `num` is correct.
- Returns `-1` if `num` is higher than the picked number.
- Returns `1` if `num` is lower than the picked number.

### **Key Concept: Why Adjust the Search Space?**

When performing binary search, we repeatedly divide the search space by half and update either the lower or upper bound based on the response from `guess(num)`:

- **If `guess(mid) == 1` (Target is Higher):**
    - The picked number **must be greater** than `mid`, so we discard the left half, including `mid`.
    - **Update:** `low = mid + 1`

- **If `guess(mid) == -1` (Target is Lower):**
    - The picked number **must be smaller** than `mid`, so we discard the right half, including `mid`.
    - **Update:** `high = mid - 1`

This ensures that we always **eliminate half of the search space** in each step, achieving an optimal time complexity of **O(log n)**.

## **Illustration of Binary Search Execution**

Let's assume `n = 10`, and the correct number is `7`. The following table illustrates each iteration:

| Step | `low` | `high` | `mid` | `guess(mid)` | Action |
|------|------|-------|------|-------------|--------|
| 1    | 1    | 10    | 5    | `1` (higher)  | `low = mid + 1 → 6` |
| 2    | 6    | 10    | 8    | `-1` (lower)  | `high = mid - 1 → 7` |
| 3    | 6    | 7     | 6    | `1` (higher)  | `low = mid + 1 → 7` |
| 4    | 7    | 7     | 7    | `0` (correct) | Return `mid = 7` |

## **Optimized Go Implementation**

```go
func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2 // Prevents integer overflow
		res := guess(mid)         // Store result to avoid redundant API calls

		switch res {
		case 0:  // Found the number
			return mid
		case -1: // The target number is lower
			high = mid - 1
		case 1:  // The target number is higher
			low = mid + 1
		}
	}

	// Unreachable case: Given constraints ensure the number exists
	panic("Unreachable: The guess function should have found the number")
}
```

## **Go Best Practices Applied**
### ✅ **Avoid Redundant Constants**
Instead of defining `higher, lower, eq := -1, 1, 0`, we directly use `guess(mid)` results inside the `switch` statement.

### ✅ **Safe Integer Midpoint Calculation**
The correct way to compute the midpoint **without overflow**:
```go
mid := low + (high - low) / 2
```
This prevents integer overflow, unlike `mid = (low + high) / 2`, which can cause issues with very large numbers.

### ✅ **Reduce API Calls**
Instead of calling `guess(mid)` multiple times, we store the result in a variable (`res`) and use it in `switch` to avoid redundant calls.

### ✅ **Use `panic()` for Unreachable Cases**
Since the problem guarantees a solution, we add a `panic()` in case of an unexpected state instead of returning `0` arbitrarily.

### ✅ **Use Clear Variable Naming**
- `low` and `high` instead of `left` and `right` to match common binary search terminology.
- This enhances readability and maintainability.

## **Key Takeaways**
1. **Adjusting the search space with `mid - 1` and `mid + 1` ensures we do not recheck already discarded values.**
2. **Binary search operates in `O(log n)`, making it optimal for large inputs.**
3. **Go best practices help write more efficient, readable, and maintainable code.**

🚀 With this understanding, you can confidently apply binary search in Go while following best practices! Happy coding! 🎯


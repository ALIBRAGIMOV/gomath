# 📌 **Find First and Last Position of Element in Sorted Array**

## 📝 **Problem Statement**
Given a sorted array `nums` in non-decreasing order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

### **Constraints:**
- Must run in **O(log n)** time complexity.

---

## 🚀 **Optimized Approach Using Binary Search**
Since the array is sorted, we use **binary search** twice:
1. **Find first occurrence** of `target`.
2. **Find first occurrence of `target+1`**, then subtract `1` to get the last occurrence.

This ensures an **O(log n)** time complexity instead of a linear scan.

### 🔹 **Binary Search Helper Function (findFirst)**
- Finds the leftmost index where `target` appears.
- If `target` is missing, it returns the insertion index.

### 🔹 **Main Function (searchRange)**
- Calls `findFirst(nums, target)` to get **left boundary**.
- Calls `findFirst(nums, target+1) - 1` to get **right boundary**.
- Ensures correctness before returning the result.

---

## 💻 **Go Implementation**

```go
package main

import "fmt"

func searchRange(nums []int, target int) []int {
    left := findFirst(nums, target)
    right := findFirst(nums, target+1) - 1

    // Ensure the target exists in range
    if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
        return []int{left, right}
    }
    return []int{-1, -1}
}

func findFirst(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2 // Avoids overflow
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main() {
    nums := []int{5, 7, 7, 8, 8, 10}
    target := 8
    fmt.Println(searchRange(nums, target)) // Output: [3, 4]
}
```

---

## 📊 **Binary Search Execution Breakdown**
Example Input: `nums = [5,7,7,8,8,10], target = 8`

| Iteration | `left` | `right` | `mid` | `nums[mid]` | Action Taken |
|-----------|--------|---------|-------|------------|--------------|
| 1         | 0      | 6       | 3     | 8          | Move `right = mid` |
| 2         | 0      | 3       | 1     | 7          | Move `left = mid + 1` |
| 3         | 2      | 3       | 2     | 7          | Move `left = mid + 1` |
| 4         | 3      | 3       | -     | -          | Found first at index `3` |

Then, repeat for `target + 1 = 9` to find `right = 4`.

---

## ✅ **Final Output**
```plaintext
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3, 4]
```

---

## 🏆 **Key Takeaways**
- **Two binary searches** ensure an **O(log n)** solution.
- **Bounds checking is essential** after search to confirm correctness.
- **Edge cases:**
    - `nums` is empty → `[-1, -1]`
    - `target` not found → `[-1, -1]`
    - `nums` has only one occurrence of `target` → returns correct index range.

By following this structured binary search approach, we can efficiently find the range of `target` in a sorted array while ensuring **optimal time complexity.** 🚀


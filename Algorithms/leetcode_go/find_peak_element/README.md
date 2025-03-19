## 🚀 Find Peak Element with Binary Search in Go

### 📌 **Problem Statement**

Given an integer array `nums`, a **peak element** is an element strictly greater than its neighbors. Assume that `nums[-1] = nums[n] = -∞`. The task is to find the index of **any peak element** in the array. The solution must achieve a time complexity of **O(log n)**.

---

### 📌 **Understanding Binary Search for Peaks**

Binary search effectively tackles this problem because:

- A peak always exists.
- We can eliminate half the array at each step based on comparing mid and mid+1.

**Core insights:**
- If `nums[mid] > nums[mid + 1]`, a peak exists in the left half (inclusive).
- Otherwise, a peak exists in the right half.

This approach quickly reduces the search space and efficiently leads to a peak.

---

### ✅ **Go Implementation**

```go
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
```

**Complexity Analysis:**
- Time Complexity: **O(log n)**
- Space Complexity: **O(1)**

---

### 🎯 **Detailed Step-by-Step Execution Example**

**Let's simulate the algorithm step-by-step with two interesting examples:**

### **Example 1:**

Input: `nums = [1, 2, 1, 3, 5, 6, 4]`

| Step | Left (`l`) | Right (`r`) | Mid | nums[Mid] | nums[Mid+1] | Decision                   | Array Snapshot          |
|------|------------|-------------|-----|-----------|--------------|----------------------------|-------------------------|
| 1    | 0          | 6           | 3   | 3         | 5            | move right (`l = mid+1`)   | `[1,2,1,3,|5,6,4]`      |
| 2    | 4          | 6           | 5   | 6         | 4            | move left (`r = mid`)      | `[1,2,1,3,5,|6,4]`      |
| 3    | 4          | 5           | 4   | 5         | 6            | move right (`l = mid + 1`) | `[1,2,1,3,5,|6|,4]`     |
| ✅   | 5          | 5           | -   | -         | -            | **Peak found at index 5**  | `[1,2,1,3,5,(6),4]`     |

### **Example 2:** (Another interesting case!)

Input: `nums = [10, 20, 15, 2, 23, 90, 67]`

| Step | Left (`l`) | Right (`r`) | Mid | nums[Mid] | nums[Mid+1] | Decision                   | Array Snapshot            |
|------|------------|-------------|-----|-----------|--------------|----------------------------|---------------------------|
| 1    | 0          | 6           | 3   | 2         | 23           | move right (`l = mid+1`)   | `[10,20,15,2,|23,90,67]`  |
| 2    | 4          | 6           | 5   | 90        | 67           | move left (`r = mid`)      | `[10,20,15,2,23,|90,67]`  |
| 3    | 4          | 5           | 4   | 23        | 90           | move right (`l = mid + 1`) | `[10,20,15,2,23,|90|,67]` |
| ✅   | 5          | 5           | -   | -         | -            | **Peak found at index 5**  | `[10,20,15,2,23,(90),67]` |

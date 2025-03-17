# **Three Sum Problem**

## **Problem Statement**
Given an integer array `nums`, return **all unique triplets** `[nums[i], nums[j], nums[k]]` such that:

- `i != j`, `i != k`, and `j != k`
- `nums[i] + nums[j] + nums[k] == 0`
- The solution set **must not contain duplicate triplets**.

### **Example 1:**
```go
Input: nums = [-1, 0, 1, 2, -1, -4]
Output: [[-1, -1, 2], [-1, 0, 1]]
```

### **Example 2:**
```go
Input: nums = [0,1,1]
Output: []
```

### **Example 3:**
```go
Input: nums = [0,0,0]
Output: [[0,0,0]]
```

## **Optimal Approach: Two-Pointer Technique**
### **Time Complexity: O(n²), Space Complexity: O(1)**
This algorithm uses sorting and a two-pointer technique to efficiently find unique triplets in **O(n²)** time complexity. The approach is as follows:

1. **Sort the array** to make duplicate removal and two-pointer traversal efficient.
2. **Iterate over the array**, fixing one element `nums[i]`.
3. **Use two pointers (`l` and `r`)** to find pairs that sum to `-nums[i]`.
4. **Skip duplicate values** to ensure unique triplets.
5. **Adjust pointers (`l++`, `r--`)** to explore possible solutions.

## **Golang Solution**
```go
// threeSum returns all unique triplets in the array that sum to zero.
func threeSum(nums []int) [][]int {
    // Step 1: Sort the array to enable two-pointer search and duplicate handling.
    sort.Ints(nums)
    n := len(nums)
    var res [][]int
    
    // Step 2: Iterate over the array, fixing one element at a time.
    for i := 0; i < n-2; i++ {
        // Optimization: If the current number is greater than 0,
        // then no three numbers can sum to 0.
        if nums[i] > 0 {
            break
        }
        // Skip duplicate values for the fixed element.
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        // Step 3: Initialize two pointers for the two-sum problem.
        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                // Found a valid triplet.
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                // Skip duplicates for the left pointer.
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                // Skip duplicates for the right pointer.
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else if sum < 0 {
                // Increase sum by moving left pointer to the right.
                l++
            } else {
                // Decrease sum by moving right pointer to the left.
                r--
            }
        }
    }
    return res
}
```

---
## **Illustration of Algorithm Execution**
### **Example Input: `nums = [-4, -1, -1, 0, 1, 2]`**
After sorting: `[-4, -1, -1, 0, 1, 2]`

| Step | `i` (Fixed) | `l` (Left) | `r` (Right) | Sum (`nums[i] + nums[l] + nums[r]`) | Action |
|------|------------|------------|-------------|-----------------------------------|--------|
| 1    | `-4`       | `-1`       | `2`         | `-4 + (-1) + 2 = -3`             | Move `l++` |
| 2    | `-4`       | `-1`       | `1`         | `-4 + (-1) + 1 = -4`             | Move `l++` |
| 3    | `-4`       | `-1`       | `0`         | `-4 + (-1) + 0 = -5`             | Move `l++` |
| 4    | `-1`       | `-1`       | `2`         | `-1 + (-1) + 2 = 0`              | Add `[-1, -1, 2]`, Move `l++, r--` |
| 5    | `-1`       | `0`        | `1`         | `-1 + 0 + 1 = 0`                 | Add `[-1, 0, 1]`, Move `l++, r--` |
| 6    | `-1`       | `1`        | `0`         | **(Pointers Crossed, Exit Loop)** | |

### **Final Output: `[[-1, -1, 2], [-1, 0, 1]]`**

---
## **Key Takeaways for Coding Interviews**
- ✅ **Sorting Helps:** Enables two-pointer traversal and easier duplicate handling.
- ✅ **Two-Pointer Technique:** Efficiently reduces the search space from **O(n³) to O(n²)**.
- ✅ **Skipping Duplicates:** Ensures only unique triplets are included in the result.
- ✅ **Edge Cases Considered:** Arrays with all zeros, no valid triplets, or duplicates.

---
## **Complexity Analysis**
| Step | Complexity |
|------|------------|
| Sorting | **O(n log n)** |
| Outer loop (`i`) | **O(n)** |
| Two-pointer search (`l, r`) | **O(n)** |
| **Total Complexity** | **O(n²)** |
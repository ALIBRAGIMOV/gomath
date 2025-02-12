# **Big-O Complexity Analysis: Ultimate Study Guide**

## **1. Introduction to Complexity Analysis**
### **What is Big-O Notation?**
Big-O notation describes how an algorithm's runtime or space requirements grow relative to input size (\( n \)).

### **Why is it Important?**
- Helps evaluate algorithm efficiency.
- Determines scalability for large inputs.
- Essential for technical interviews and performance tuning.

---

## **2. Understanding Time Complexity**
### **Rules of Thumb**
1. **Ignore constants and low-order terms**: \( O(2n) \) simplifies to \( O(n) \).
2. **Focus on dominant operations**: Nested loops dominate over single loops.
3. **Common patterns**:
    - Iterating once: **O(n)**
    - Nested loops: **O(n²)**
    - Halving input: **O(log n)**
    - Recursion with branching: **O(2ⁿ)**

### **Common Time Complexities**
| Complexity | Name | Example |
|------------|-----------|------------|
| O(1) | Constant | Accessing an array element |
| O(log n) | Logarithmic | Binary search |
| O(n) | Linear | Iterating through an array |
| O(n log n) | Linearithmic | Merge Sort |
| O(n²) | Quadratic | Bubble Sort |
| O(2ⁿ) | Exponential | Fibonacci (naive recursion) |
| O(n!) | Factorial | Traveling Salesman (brute-force) |

---

## **3. Analyzing Loops & Recursion**
### **Iterative Complexity Analysis**
#### **Example 1: Single Loop (O(n))**
```go
func printItems(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(i)
    }
}
```
✅ Loop runs \( n \) times → **O(n)**.

#### **Example 2: Nested Loops (O(n²))**
```go
func printPairs(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            fmt.Println(arr[i], arr[j])
        }
    }
}
```
✅ Each loop runs \( n \) times → **O(n²)**.

### **Recursive Complexity Analysis**
#### **Example 3: Logarithmic Complexity (O(log n))**
```go
func binarySearch(arr []int, target int) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1
}
```
✅ Each step halves the array → **O(log n)**.

#### **Example 4: Exponential Complexity (O(2ⁿ))**
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
✅ Each call branches into two → **O(2ⁿ)**.

---

## **4. Understanding Space Complexity**
### **Rules of Thumb**
1. **Primitive variables** (int, float, bool) → **O(1)**.
2. **Arrays/Slices** of size \( n \) → **O(n)**.
3. **Recursive calls** → Depth of recursion determines space.

### **Examples**
#### **Example 1: Constant Space (O(1))**
```go
func sum(a, b int) int {
    return a + b
}
```
✅ Uses fixed memory → **O(1)**.

#### **Example 2: Linear Space (O(n))**
```go
func createArray(n int) []int {
    arr := make([]int, n)
    return arr
}
```
✅ Allocates an array of size \( n \) → **O(n)**.

---

## **5. Choosing the Best Algorithm for a Problem**
### **Step 1: Identify Constraints**
| Input Size (n) | Acceptable Complexity |
|--------------|---------------------|
| \( n \leq 100 \) | O(n²) is acceptable |
| \( n \leq 10^6 \) | O(n log n) is required |
| \( n > 10^6 \) | O(n) or better needed |

### **Step 2: Match the Problem Type**
| Problem | Best Approach | Complexity |
|-------------|--------------|------------|
| Searching  | Binary search | O(log n) |
| Sorting    | Quick/Merge sort | O(n log n) |
| Pathfinding | Dijkstra’s algorithm | O(V + E log V) |
| Recursion  | Dynamic programming | O(n) |

---

## **6. Interview Strategy: How to Evaluate Complexity in Problems**
### **1. Ask Clarifying Questions**
- "What is the maximum input size?"
- "Should we optimize for time or space?"

### **2. Identify the Core Operation**
- **Single loop?** → **O(n)**
- **Nested loops?** → **O(n²)**
- **Halving the input?** → **O(log n)**
- **Recursive calls?** → **Depends on recurrence relation**

### **3. Explain Your Thought Process**
- "Since we iterate through the array once, the complexity is O(n)."
- "Each level of recursion reduces input size by half, so it’s O(log n)."

---

## **7. Example Interview Problem: Finding Duplicates**
### **Question:** Given an array of integers, return all duplicates.

### **Brute Force Solution (O(n²))**
```go
func findDuplicatesBrute(arr []int) []int {
    n := len(arr)
    var result []int
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if arr[i] == arr[j] {
                result = append(result, arr[i])
            }
        }
    }
    return result
}
```

### **Optimized Solution Using Hash Map (O(n))**
```go
func findDuplicates(arr []int) []int {
    seen := make(map[int]bool)
    var result []int
    for _, num := range arr {
        if seen[num] {
            result = append(result, num)
        }
        seen[num] = true
    }
    return result
}
```

### **Comparison**
| Approach | Time Complexity | Space Complexity |
|----------|---------------|---------------|
| Brute force | O(n²) | O(1) |
| Hash map | O(n) | O(n) |

---

## **8. Summary**
✅ **Understand complexity patterns** (loops, recursion, branching).

✅ **Use the right approach based on constraints** (O(n) for large input sizes).

✅ **Explain your choices clearly in interviews.**

🚀 **TODO improve this page do it more informal and add more interesting cases for understating,
and cases for how start and chose correct interview Questions step by step solver**


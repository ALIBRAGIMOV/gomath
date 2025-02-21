# **Feasible Function — Best Explanation on Golang**

A **feasible function** is a function that is **efficient, scalable, and practical** in terms of execution time, memory usage, and maintainability. It adheres to **Golang best practices** and is optimized for **high-load applications**.

- And We will call the function feasible to signify whether the element at the current index is feasible (True) or not (False) to meet the problem constraints.
-  a feasible function is a critical concept that serves as the condition upon which the binary search decision-making is based. This function examines an element at a particular index to determine whether it satisfies the conditions of the problem at hand.
- Binary search operates on a sorted array (or more generally, a monotonic condition), and it is traditionally used to find a specific target value within that array. However, sometimes the goal is not to find a specific value but to determine the first element in the array that meets a certain condition. This condition is what the feasible function defines.
---

## **1. Key Characteristics of a Feasible Function**
A function is considered **feasible** if it meets the following criteria:

### ✅ **Computational Efficiency**
- Runs within an **optimal time complexity** for the given problem.
- **Prefer O(1), O(log n), O(n) when possible**; avoid **O(2ⁿ) (exponential growth)**.
- Uses **iterative solutions instead of deep recursion** to prevent stack overflow.

#### **Example: Feasible Binary Search (O(log n))**
```go
func BinarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := int(uint(left+right) >> 1) // Prevents integer overflow

        switch {
        case nums[mid] == target:
            return mid
        case nums[mid] < target:
            left = mid + 1
        default:
            right = mid - 1
        }
    }
    return -1
}
```
**Why Feasible?**
- **O(log n) complexity** ensures efficiency.
- **Avoids integer overflow** with `uint(left+right) >> 1`.
- **Clean and structured logic** using `switch`.

---

### ✅ **Minimal & Predictable Resource Usage**
- Avoids **excessive memory allocations and CPU usage**.
- **No redundant computations or unnecessary recursion**.
- Uses **stack-friendly** iterative approaches when possible.

#### **Example: Infeasible vs. Feasible Fibonacci Calculation**
🚨 **Bad: Exponential Recursion (O(2ⁿ))**
```go
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2) // 🚨 Slow and memory-intensive
}
```
✔ **Good: Iterative (O(n), O(1) space)**
```go
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```
**Why Feasible?**
- **Runs in O(n) time instead of O(2ⁿ).**
- **Uses constant O(1) memory.**
- **Prevents stack overflow** from deep recursion.

---

### ✅ **Edge Case Handling**
- A feasible function **never panics unexpectedly**.
- Always **handles invalid input gracefully**.
- Returns **errors** instead of allowing silent failures.

#### **Example: Handling Division by Zero**
```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero is not allowed")
    }
    return a / b, nil
}
```
**Why Feasible?**
- **Prevents division by zero errors.**
- **Returns an error instead of panicking.**
- **Makes error handling explicit.**

---

### ✅ **Readable, Maintainable & Minimalistic**
- Uses **clear variable names** and **consistent formatting**.
- Follows **Golang best practices** for simplicity.
- **Prefers explicit over clever but unreadable code**.

#### **Example: Improving Readability**
🚨 **Unfeasible: Hard to Read**
```go
func f(x int) int {
    if x > 0 {
        if x%2 == 0 {
            return x / 2
        } else {
            return x*3 + 1
        }
    }
    return 0
}
```
✔ **Feasible: Structured & Maintainable**
```go
func Transform(x int) int {
    switch {
    case x <= 0:
        return 0
    case x%2 == 0:
        return x / 2
    default:
        return x*3 + 1
    }
}
```
**Why Feasible?**
- **Uses `switch` for clarity.**
- **Reduces unnecessary nesting.**
- **Improves readability and maintainability.**

---

## **2. Summary: Mastering Feasible Functions**
| Factor                | Feasible Example            | Infeasible Example |
|-----------------------|----------------------------|---------------------|
| **Time Complexity**  | O(log n) Binary Search     | O(2ⁿ) Recursive Fibonacci |
| **Memory Usage**      | O(1) Iterative Fibonacci   | O(n) Recursion Stack |
| **Edge Cases**        | Handles division by zero  | Panics on zero division |
| **Readability**       | Uses `switch` for clarity | Nested `if-else` mess |
| **Concurrency**       | Uses Goroutines & channels | Unoptimized blocking calls |

---
🔹 **Mastering feasibility in function design will make you a top-tier backend developer.**  
Always **optimize for clarity, performance, and scalability!** 🚀


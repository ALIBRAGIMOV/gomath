# Understanding the Sliding Window Algorithm for Finding the Longest Substring Without Repeating Characters

## 🚀 Problem Statement
Given a string `s`, find the length of the longest substring without repeating characters.

---

## 🔍 Solution Approach: Sliding Window with a Hash Map

### **Key Variables & Their Roles**
- **`l` (Left Pointer)** – Marks the beginning of the current substring.
- **`r` (Right Pointer)** – Expands the substring by iterating through `s`.
- **`lastSeen` (Hash Map `map[string]int`)** – Stores the last seen index of each character.
- **`maxLen` (Maximum Length Found)** – Tracks the longest substring encountered so far.

### **Core Idea**
1. Move `r` (right pointer) through the string while adding characters to the window.
2. If a character is already in `lastSeen` **and** within the window (`lastSeen[char] >= l`), move `l` to skip past the duplicate.
3. Update `lastSeen[char]` with the new index of `char`.
4. Update `maxLen` with the current window size (`r - l + 1`).

---

## 📝 Code Implementation

```go
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    lastSeen := make(map[string]int) // Track the last index of each character as a string
    l, maxLen := 0, 0

    for r := 0; r < len(s); r++ {
        char := string(s[r]) // Convert only once per iteration

        // If character is found and is within the current window, move l forward
        if lastIdx, found := lastSeen[char]; found && lastIdx >= l {
            l = lastIdx + 1 // Shift left pointer past the duplicate
        }

        lastSeen[char] = r // Update last seen index
        maxLen = max(maxLen, r-l+1) // Update max length
    }

    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

## 🔬 **Deep Dive: `lastIdx` Explanation**
### What is `lastIdx`?
`lastIdx` represents the **last recorded index** of the character `char` in the `lastSeen` map. When `char` is found in `lastSeen`, it means that this character was previously encountered at index `lastIdx`.

### Why Do We Check `lastIdx >= l`?
If `lastIdx` is **before** `l`, it means that the previous occurrence of `char` is no longer part of the active window, so we **don't** need to move `l`.

If `lastIdx >= l`, the character is inside our window, meaning a duplicate exists. Thus, we **must** move `l` to `lastIdx + 1` to remove the duplicate from the window.

### Example Walkthrough: `s = "abcabcbb"`
#### **Initialization:**
```
s = "abcabcbb"
lastSeen = {}  // Empty map
l = 0, maxLen = 0
```

#### **Step-by-Step Execution:**
| `r` | `s[r]` | `lastSeen` (stored index) | `l` (left pointer) | `r - l + 1` (window size) | `maxLen` |
|----|----|----------------|----|----|----|
| 0  | 'a' | `{a: 0}` | 0  | 1  | 1  |
| 1  | 'b' | `{a: 0, b: 1}` | 0  | 2  | 2  |
| 2  | 'c' | `{a: 0, b: 1, c: 2}` | 0  | 3  | 3  |
| 3  | 'a' | `{a: 3, b: 1, c: 2}` | **1** | 3  | 3  |
| 4  | 'b' | `{a: 3, b: 4, c: 2}` | **2** | 3  | 3  |
| 5  | 'c' | `{a: 3, b: 4, c: 5}` | **3** | 3  | 3  |
| 6  | 'b' | `{a: 3, b: 6, c: 5}` | **5** | 2  | 3  |
| 7  | 'b' | `{a: 3, b: 7, c: 5}` | **7** | 1  | 3  |

✅ **Final Answer**: `maxLen = 3` (longest substring: "abc").

---

## 📈 Complexity Analysis
| Complexity | Calculation |
|------------|-------------|
| **Time Complexity** | **O(n)** - Each character is processed at most twice (once by `r`, once by `l`). |
| **Space Complexity** | **O(min(n, |Σ|))** - The map stores at most `n` entries or `|Σ|` (alphabet size). |

- **Best case:** `O(1)` space (small alphabet, e.g., all lowercase English letters).
- **Worst case:** `O(n)` space (all unique characters in a large string).

---

## ✅ Why This Approach Works Best
1. **Sliding Window Efficiency** – The left pointer only moves **forward**, ensuring `O(n)` time complexity.
2. **Hash Map for Quick Lookups** – `lastSeen[char]` provides `O(1)` time complexity for checking duplicates.
3. **No Extra Allocations** – No substring slicing, no extra arrays, only `map` storage.

---

## 🔥 Summary
- **Move `r` to expand the window**.
- **Use `map` to track last-seen indices**.
- **If `s[r]` repeats, move `l` past the last occurrence**.
- **Update `maxLen` dynamically**.

💡 **Sliding Window + Hash Map = Optimal O(n) solution!** 🚀


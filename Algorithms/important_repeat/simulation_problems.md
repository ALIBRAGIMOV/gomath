- https://algo.monster/problems/simulation_intro

# LeetCode Simulation & Problem-Solving Protocol

# Guided Practice

We've designed a **step-by-step guided practice system** that mirrors exactly how successful candidates solve problems in real technical interviews.

## **1. Identify the Pattern**
- Apply common patterns from **AlgoMonster** to quickly recognize the best problem-solving approach.
- Map the problem to well-known techniques: **sliding window, two pointers, DFS/BFS, dynamic programming, graph traversal, etc.**

## **2. Implement the Solution**
- Write **clean, bug-free code** following best practices.
- Use structured templates to reduce errors and improve readability.
- Leverage **AI assistance** for refining edge cases and optimizing performance.

## **3. Analyze Time & Space Complexity**
- Clearly demonstrate an understanding of **algorithm efficiency**.
- Explain the trade-offs between different approaches.
- Ensure the solution meets problem constraints without unnecessary overhead.


## **Step-by-Step Problem-Solving Process**

### **1. Mental Preparation (1-2 min)**
- Take deep breaths (4-7-8 method) to reduce stress.
- Visualize explaining and solving the problem confidently.
- Adopt a positive mindset: *"I will solve this systematically."*

---

### **2. Understand the Problem (5-8 min)**
1. **Read the problem carefully.** Read twice to ensure clarity.
2. **Restate the problem in simple terms.** Summarize the goal out loud or in writing.
3. **Identify key constraints:**
    - **Input**: Type, size, range.
    - **Output**: Expected format.
    - **Edge cases**: Empty input, duplicates, max/min limits.
4. **Visualize the process in your mind.** Imagine how the solution works step by step.
5. **Write down observations.** Highlight patterns or special conditions.

---

### **3. Plan the Solution (8-12 min)**
1. **Identify the best approach** (e.g., sliding window, two-pointer, DFS, DP, hash map, sorting).
2. **Choose data structures.** Pick the most efficient ones for the problem.
3. **Sketch the logic with comments.** Write out the main steps before coding.
4. **Write pseudocode** in a structured manner to guide implementation.
5. **Predict time & space complexity.** Ensure it meets constraints.

> **Detailed Example Plan for 'Longest Substring Without Repeating Characters':**
```go
// Step 1: Understand the problem
// - We need to find the length of the longest substring without repeating characters.
// - A substring is a continuous sequence of characters.
// - The order must be maintained.
//
// Step 2: Identify constraints and edge cases
// - Input: a string 's' (1 ≤ len(s) ≤ 10^5)
// - Output: an integer representing the max substring length.
// - Constraints: must be O(n) for efficiency.
// - Edge cases: "", "aaaaa", "abcabcbb", "pwwkew"
//
// Step 3: Select an approach
// - Using a sliding window technique is efficient because:
//   - We maintain a moving window of non-repeating characters.
//   - When a duplicate is found, we move the left pointer to remove previous instances.
//   - The right pointer expands the window while tracking max length.
//
// Step 4: Choose a data structure
// - A hash map (map[rune]int) is used to track the last index of characters in the substring.
// - Two pointers (start and end) track the boundaries of the window.
//
// Step 5: Outline algorithm and pseudocode
// - Initialize an empty hash map to store character indexes.
// - Initialize variables maxLen = 0 and start = 0.
// - Iterate through the string with an 'end' pointer:
//   - If s[end] exists in hash map and its index is >= start:
//     - Move start to last seen index + 1.
//   - Update the last seen index of s[end].
//   - Calculate current window size and update maxLen if needed.
// - Return maxLen.
```

---

### **4. Implement Efficiently (15-25 min)**
1. **Write clean, modular code.** Small functions, clear variable names.
2. **Explain each line as you write it.** This reinforces understanding.
3. **Follow Go best practices:**
    - Use `map` for fast lookups.
    - Avoid deep nesting; prefer early returns.
    - Handle edge cases upfront.

```go
func longestSubstring(s string) int {
    lastSeen := make(map[rune]int) // Store the last position of each character
    maxLen, start := 0, 0         // Initialize max length and start of window

    for end, char := range s {    // Iterate through characters
        if idx, found := lastSeen[char]; found && idx >= start {
            start = idx + 1        // Move the start to avoid duplicate characters
        }
        lastSeen[char] = end       // Update last seen position of current character
        maxLen = max(maxLen, end-start+1) // Update max length of substring
    }
    return maxLen // Return the result
}
```

---

### **5. Test and Debug (8-12 min)**
1. **Run basic test cases.** Check normal cases first.
2. **Test edge cases.** Empty input, single char, long string.
3. **Use print statements if needed.** Debug by tracking key variables.
4. **Step through execution in your mind.** Imagine how variables change.

```go
tests := []struct {
    input string
    want  int
}{
    {"", 0}, {"aaaaa", 1}, {"abcabcbb", 3}, {"pwwkew", 3},
}
for _, tc := range tests {
    got := longestSubstring(tc.input)
    if got != tc.want {
        fmt.Printf("FAIL %q: got %d, want %d\n", tc.input, got, tc.want)
    }
}
```

---

### **6. Optimize and Reflect (5-10 min)**
1. **Review complexity:**
    - Time: O(n) – each char is processed once.
    - Space: O(min(n, charset_size)) – limited by unique chars.
2. **Refactor for readability.** Remove redundant code.
3. **Write key takeaways:**
    - What worked well?
    - What slowed me down?
    - What will I do differently next time?

---

## **Daily Neuroscience-Based Practice**
1. **Before Coding:** Visualization + deep breathing.
2. **During Coding:** Pomodoro method (25 min focus / 5 min break).
3. **After Coding:** Write reflections, revisit past problems weekly.

---

## **Why This Method Works**
✅ **Eliminates Random Guesswork** – Replaces trial-and-error with structured, predictable problem-solving.  
✅ **Builds Strong Problem-Solving Reflexes** – Reinforces **muscle memory** needed for fast and accurate coding.  
✅ **Develops Recruiter-Preferred Skills** – Enhances **pattern recognition, clean implementation, and complexity analysis**.

---

Each practice problem is carefully selected to **reinforce specific patterns** from the curriculum, ensuring a **seamless learning journey** from concept to application.
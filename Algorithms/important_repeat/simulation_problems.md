- https://algo.monster/problems/simulation_intro

# LeetCode Simulation & Problem-Solving Protocol

## 🚀 Guided Practice (FAANG+ Interview Methodology)

- This structured approach mirrors how successful candidates consistently perform in technical interviews at FAANG+ companies. 
- Incorporating neuroscience-based learning, active recall techniques, and mindset growth principles from *Mindset: The New Psychology of Success*, it prepares you to efficiently solve medium, hard, and super hard algorithmic problems.

---

## ✅ Step-by-Step Problem-Solving Framework

### **1. Mental Preparation (1-2 min)**
- Practice deep breathing (HRV resonance 4-6 method) to reduce anxiety.
- Visualize yourself confidently explaining and solving the problem.
- Adopt a growth mindset: *"I will learn and master this systematically."*

---

### **2. Understand the Problem (5-8 min)**
- **Read thoroughly twice** to avoid misunderstandings.
- Restate the problem clearly in simple terms.
- Clarify constraints explicitly:
   - Input: Types, sizes, and range
   - Output format
   - Edge cases: Empty inputs, duplicates, max/min limits
- Visualize the entire solution step-by-step in your mind.
- Document initial insights and observations clearly.

---

### **3. Plan the Solution (8-12 min)**
- Match the problem to familiar patterns:
   - Sliding window, two pointers, binary search
   - DFS/BFS, dynamic programming, graph algorithms
- Choose optimal data structures based on constraints.
- Clearly outline your logic in structured pseudocode or high-level comments.
- Explicitly predict time and space complexity.

**Example (Sliding Window Problem):**
```go
// Identify the longest substring without repeating characters:
// 1. Sliding window tracks unique characters.
// 2. Move left pointer forward on duplicates.
// 3. Hash map records last-seen indices.
// Complexity: Time O(n), Space O(min(n, alphabet size))
```

---

### **4. Implement Clearly and Efficiently (10-25 min)**
- Write clean, modular, and readable code.
- Explain each step as you code (think aloud).
- Follow Go best practices:
   - Clear variable naming, minimal nesting, early returns.
   - Immediate handling of edge cases.
- Ensure code is easily maintainable and efficient.

```go
func longestSubstring(s string) int {
    lastSeen := make(map[rune]int)
    maxLen, start := 0, 0

    for end, char := range s {
        if idx, found := lastSeen[char]; found && idx >= start {
            start = idx + 1
        }
        lastSeen[char] = end
        maxLen = max(maxLen, end-start+1)
    }
    return maxLen
}
```

---

### **4. Test & Debug Effectively (8-12 min)**
- Run targeted test cases covering standard scenarios first.
- Test rigorously for edge cases (empty inputs, duplicates, limits).
- Use structured debugging techniques:
   - Print statements to monitor variable states if necessary.
   - Mentally simulate execution step-by-step.

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

### **5. Reflect & Optimize (5-10 min)**
- Review time & space complexities explicitly.
- Refactor your code for readability and efficiency.
- Document key insights:
   - What strategies worked well?
   - What slowed your progress?
   - Specific steps for improvement next time.

---

## 🔑 Neuroscience-Based Daily Practice
- **Before Problem Solving:** Visualization, breathing techniques, positive affirmations.
- **During Practice:** Pomodoro technique (25 min focused coding, 5 min rest).
- **Post-Session Reflection:** Actively recall solutions, weekly revision using Anki cards.

---

## ⚠️ If Stuck After 45 Minutes

### **For Medium Problems**
- **Pause & Reflect (5 min)**:
   - Clearly identify what’s confusing you.
   - Restate the problem aloud and check for missed constraints.

- **Structured Mini-Planning (10 min)**:
   - Break down problem into smaller, simpler steps.
   - Restate clearly, using pseudocode.

- **Hint Strategy (5-10 min)**:
   - Seek minimal assistance (hint only).
   - Immediately try to solve based on the hint.

- **Active Solution Study (15 min)**:
   - Deeply analyze the provided solution.
   - Re-implement solution from memory, comparing differences.

- **Immediate Reinforcement**:
   - Solve a similar or slightly modified problem to reinforce the concept.

---

## 📌 Weekly Reflection & Active Recall
- Maintain a log (Anki cards or spreadsheet) of recurring issues or difficult concepts.
- Each week, revisit and actively solve previous challenging problems.

---

## 📚 Growth Mindset Principles (Carol Dweck’s "Mindset")
- View every challenge as an opportunity for growth, not as failure.
- Embrace struggles as valuable feedback, guiding your learning process.
- Understand that problem-solving skills improve through consistent effort and deliberate practice, rather than innate ability alone.

---

## ✅ Why This Protocol Succeeds
- Replaces random guesswork with a structured, proven method.
- Develops essential skills: pattern recognition, methodical planning, clear communication.
- Creates a resilient mindset, preparing you not only for interviews but also for ongoing professional growth.

---

## 🌟 Final Mindset Reminder
Adopt Carol Dweck’s growth mindset philosophy:
- Challenges are opportunities, not setbacks.
- Mistakes are critical learning moments.
- Your goal isn't perfection—it's continuous, systematic growth toward mastery.
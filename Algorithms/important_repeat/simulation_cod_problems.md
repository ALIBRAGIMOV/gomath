- https://algo.monster/problems/simulation_intro

## What are Simulation Problems?

Simulation problems in coding interviews require you to create a program that imitates the behavior of a system or a process. These problems involve stepping through a scenario described in the problem statement, maintaining state, and making decisions based on given rules. The goal is to accurately model a real-world process while adhering to problem constraints.

Simulation problems are generally easy to understand but demand meticulous implementation to account for all edge cases. Unlike algorithm-heavy problems, they focus more on accurately modeling the process and managing state transitions.

### Key Characteristics of Simulation Problems:

1. **State Management** – Tracking and updating the state of different elements in the system.
2. **Step-by-Step Execution** – Iterating through each step and modifying the state accordingly.
3. **Conditional Logic** – Making decisions based on the current state and defined rules.
4. **Edge Case Handling** – Managing unusual scenarios to ensure correctness.

---

## Example: Design a Parking System

### **Problem Statement**

Design a parking system for a parking lot with three types of parking spaces: **big, medium,** and **small**, each with a fixed number of slots. Implement a class `ParkingSystem` with the following methods:

- **`ParkingSystem(int big, int medium, int small)`**: Initializes the object with the number of slots for each parking type.
- **`bool addCar(int carType)`**: Checks if there is a parking space available for a car of the given type (`1` for big, `2` for medium, and `3` for small). If available, it parks the car and returns `true`. Otherwise, it returns `false`.

---

## Thought Process for Solving Simulation Problems

### **Step 1: Understand the Problem**

- We need to manage parking slots for three car types: **big, medium, and small**.
- Each `addCar` operation modifies the system state by reducing available slots.

### **Step 2: Identify the Simulation Aspects**

- Maintain and update the count of available slots for each car type.
- Check the availability before parking a car.
- Ensure the parking system handles edge cases (e.g., zero slots available).

### **Step 3: Plan the Simulation**

1. **Initialize the Parking System**
    - Store available slot counts for each car type.
2. **Handle `addCar` Operations**
    - Check the `carType`.
    - If a slot is available for the given type, **reduce the count** and return `true`.
    - If no slots are available, return `false`.
3. **Edge Case Handling**
    - If parking slots are initialized with `0`, ensure `addCar` always returns `false` for that type.

### **Step 4: Implement the Solution in Golang**

```go
package main

import "fmt"

// ParkingSystem struct holds available slots for each type of car
type ParkingSystem struct {
    big, medium, small int
}

// Constructor initializes the parking system
func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big, medium, small}
}

// addCar method checks availability and updates the state
func (p *ParkingSystem) addCar(carType int) bool {
    switch carType {
    case 1:
        if p.big > 0 {
            p.big--
            return true
        }
    case 2:
        if p.medium > 0 {
            p.medium--
            return true
        }
    case 3:
        if p.small > 0 {
            p.small--
            return true
        }
    }
    return false
}

func main() {
    ps := Constructor(1, 1, 0) // 1 big, 1 medium, 0 small slots
    fmt.Println(ps.addCar(1)) // true (big slot available)
    fmt.Println(ps.addCar(2)) // true (medium slot available)
    fmt.Println(ps.addCar(3)) // false (no small slots)
    fmt.Println(ps.addCar(1)) // false (no more big slots left)
}
```

---

## **Key Takeaways**

- **Simulation problems test your ability to model dynamic systems.**
- **Clearly define states** and **transition rules**.
- **Edge cases matter** – consider boundary conditions before implementing.
- **Use structured problem-solving techniques** to break down the problem step by step.

Simulation problems are a crucial aspect of coding interviews, assessing how well you can model real-world processes and handle state transitions efficiently. Mastering these will significantly improve your problem-solving skills for technical interviews. 🚀


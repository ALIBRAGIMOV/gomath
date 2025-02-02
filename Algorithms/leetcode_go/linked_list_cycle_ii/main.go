package main

import (
	"fmt"
)

// ListNode represents a single node in a singly linked list.
type ListNode struct {
	Val  int       // Value stored in the node
	Next *ListNode // Pointer to the next node in the list
}

// hasCycle detects whether a linked list has a cycle using the Floyd’s Cycle Detection Algorithm.
// This is also known as the "two pointers" or "fast and slow" technique.
//
// - The `slow` pointer moves one step at a time.
// - The `fast` pointer moves two steps at a time.
// - If there is a cycle, the fast pointer will eventually meet the slow pointer.
// - If there is no cycle, the fast pointer will reach the end of the list (nil).
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false // Empty list has no cycle
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer by one step
		fast = fast.Next.Next // Move fast pointer by two steps

		if slow == fast {
			return true // If they meet, there is a cycle
		}
	}

	return false // If fast reaches the end, there is no cycle
}

// implement - https://leetcode.com/problems/linked-list-cycle/
// Time Complexity: O(N), where N is the number of nodes.
// Space Complexity: O(1), since only two pointers (slow and fast) are used.

func main() {
	cycleIndex := -1
	// Example 1: No cycle
	list1 := []int{1, 2, 3, 4, 5}
	lN1 := createLinkedList(list1, cycleIndex)
	fmt.Printf("list1 - %+v - cycle_index %d - %t \n", list1, cycleIndex, hasCycle(lN1)) //	list1 - [1 2 3 4 5] - cycle_index -1 - false

	// Example 2: Cycle at index 2 (3rd node)
	list2 := []int{10, 20, 30, 40, 50}

	cycleIndex = 2
	lN2 := createLinkedList(list2, cycleIndex)
	fmt.Printf("list2 - %+v - cycle_index %d - %t \n", list2, cycleIndex, hasCycle(lN2)) //	list2 - [10 20 30 40 50] - cycle_index 2 - true
}

// createLinkedList constructs a linked list from a given slice of integers.
// If cycleIndex is valid (>= 0), it creates a cycle at the given index.
func createLinkedList(values []int, cycleIndex int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	var cycleNode *ListNode

	for i := 1; i < len(values); i++ {
		node := &ListNode{Val: values[i]}
		current.Next = node
		current = node

		if i == cycleIndex {
			cycleNode = node
		}
	}

	// If cycleIndex is valid, connect the last node to cycleNode
	if cycleIndex >= 0 {
		current.Next = cycleNode
	}

	return head
}

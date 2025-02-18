package main

import (
	"errors"
	"fmt"
)

// Stack represents a generic stack data structure.
type Stack[T any] struct {
	elements []T
}

// NewStack creates a new Stack instance.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack.
// It returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T // Return zero value of T
		return zero, errors.New("stack is empty")
	}

	// Get the last element
	n := len(s.elements) - 1
	elem := s.elements[n]

	// Resize slice to remove the last element
	s.elements = s.elements[:n]
	return elem, nil
}

// Peek returns the top element without removing it.
// It returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func main() {
	// Create a stack for integers
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)

	value, _ := intStack.Pop()
	fmt.Println(value) // Output: 20

	// Create a stack for strings
	stringStack := NewStack[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")

	top, _ := stringStack.Peek()
	fmt.Println(top) // Output: World
}

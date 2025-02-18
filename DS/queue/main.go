package main

import "errors"

// Queue represents a generic queue data structure.
type Queue[T any] struct {
	elements []T
}

// NewQueue creates a new Queue instance.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{elements: make([]T, 0)}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element from the queue.
// It returns an error if the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	elem := q.elements[0]
	q.elements = q.elements[1:]
	return elem, nil
}

// Peek returns the front element without removing it.
// It returns an error if the queue is empty.
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

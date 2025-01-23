package data_structures

import "errors"

type Queue[T any] struct {
	elements []T
	size     int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0),
		size:     0,
	}
}

func (q *Queue[T]) EnqueueElements(values []T) {
	q.elements = append(q.elements, values...)
	q.size += len(values)
}

func (q *Queue[T]) Enqueue(val T) {
	q.elements = append(q.elements, val)
	q.size += 1
}

func (q *Queue[T]) Dequeue() (T, error) {
	var element T
	if q.IsEmpty() {
		// if queue is already empty, element will be zero value of type T
		return element, errors.New("queue is empty")
	}
	element = q.elements[0]
	q.elements = q.elements[1:]
	q.size -= 1

	return element, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

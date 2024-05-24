package store

import (
	"errors"
)

var GlobalState GlobalStateType

func init() {
	GlobalState = GlobalStateType{
		State: make([]*StateType, 11), // store the last 10 states
		size:  11,
		front: 0,
		rear:  0,
	}
}

func NewState() *StateType {
	return &StateType{
		Status:     false,
		Height:     0,
		BlockSign:  false,
		OracleSign: false,
		DoubleSign: false,
	}
}

// IsEmpty checks if the queue is empty
func (q *GlobalStateType) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull checks if the queue is full
func (q *GlobalStateType) IsFull() bool {
	return (q.rear+1)%q.size == q.front
}

// Enqueue adds an element to the queue
func (q *GlobalStateType) Enqueue(value *StateType) error {
	// if q.IsFull() {
	// 		return errors.New("queue is full")
	// 	}
	q.State[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	return nil
}

// Dequeue removes and returns an element from the queue
func (q *GlobalStateType) Dequeue() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	value := q.State[q.front]
	q.front = (q.front + 1) % q.size
	return value, nil
}

// Front returns the front element of the queue without removing it
func (q *GlobalStateType) Front() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.State[q.front], nil
}

// Rear returns the rear element of the queue without removing it
func (q *GlobalStateType) Rear() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	// rear-1 is the last enqueued element
	return q.State[(q.rear-1+q.size)%q.size], nil
}

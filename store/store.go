package store

import (
	"errors"
	"fmt"
)

var GlobalState GlobalStateType

var Window int64 = 30
var Threshold uint64 = 10

func init() {
	GlobalState = GlobalStateType{
		Status: &StatusType{
			Status:        true,
			OracleMissed:  fmt.Sprintf("0 / %d", Window),
			Uptime:        "100%",
			WindowSize:    uint64(Window),
			OracleMissCnt: 0,
		},
		States: make([]*StateType, 11), // store the last 10 states
		size:   11,
		front:  0,
		rear:   0,
	}
}

func NewState() *StateType {
	return &StateType{
		Height:           0,
		BlockSign:        false,
		OracleSign:       false,
		OracleDoubleSign: false,
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
	q.States[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	return nil
}

// Dequeue removes and returns an element from the queue
func (q *GlobalStateType) Dequeue() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	value := q.States[q.front]
	q.front = (q.front + 1) % q.size
	return value, nil
}

// Front returns the front element of the queue without removing it
func (q *GlobalStateType) Front() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.States[q.front], nil
}

// Rear returns the rear element of the queue without removing it
func (q *GlobalStateType) Rear() (*StateType, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	// rear-1 is the last enqueued element
	return q.States[(q.rear-1+q.size)%q.size], nil
}

package main

import "fmt"

// CircularQueue defines a circular queue
type CircularQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

// NewCircularQueue creates a queue
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// IsEmpty returns true if queue is empty
func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// IsFull returns true if queue is full
func (q *CircularQueue) IsFull() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

// Enqueue puts a element in the tail of queue
func (q *CircularQueue) enqueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// Dequeue fetches a element from queue
func (q *CircularQueue) dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

// String prints the queue
func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}

func main() {
	cq := NewCircularQueue(5)
	cq.enqueue(1)
	cq.enqueue(2)
	cq.enqueue(3)
	cq.enqueue(4)
	cq.enqueue(5)
	cq.enqueue(6)

	for i := 0; i < 10; i++ {
		x := cq.dequeue()
		fmt.Println(x)
	}
}

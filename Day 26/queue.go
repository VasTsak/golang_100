package main

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func main() {
	var queue Queue

	queue.enqueue("element 1")
	queue.enqueue("element 2")
	queue.enqueue("element 3")

	for len(queue) > 0 {
		x, y := queue.dequeue()
		if y == true {
			fmt.Println(x)
		}
	}
}

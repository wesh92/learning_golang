package main

import (
	"fmt"
)

type Queue struct {
	vals []uint8
}

func (q *Queue) Enqueue(val uint8) {
	q.vals = append(q.vals, val)
}

func (q *Queue) Dequeue() uint8 {
	returnItem := q.vals[0]
	q.vals = q.vals[1:]
	return returnItem
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(15)
	q.Enqueue(67)
	for len(q.vals) != 0 {
		fmt.Println(q.Dequeue())
	}
}

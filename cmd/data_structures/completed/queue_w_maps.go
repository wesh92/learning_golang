package main

import "fmt"

type MapQueue struct {
	vals []map[string]int
}

func (q *MapQueue) Enqueue(inVals map[string]int) {
	q.vals = append(q.vals, inVals)
}

func (q *MapQueue) Dequeue() map[string]int {
	returnMap := q.vals[0]
	q.vals = q.vals[1:]
	return returnMap
}

func main() {
	q := MapQueue{}
	q.Enqueue(map[string]int{
		"A": 1,
		"B": 2,
	})
	q.Enqueue(map[string]int{
		"C": 3,
		"B": 4,
	})
	for len(q.vals) != 0 {
		fmt.Println(q.Dequeue())
	}
}

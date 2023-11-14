package main

import "fmt"

type Queue struct {
	vals []uint8
}

func (q *Queue) EnqueueSlice(sliceVals []uint8) {
	q.vals = sliceVals
}

func (q *Queue) DequeueSlice() uint8 {
	returnItem := q.vals[0]
	q.vals = q.vals[1:]
	return returnItem
}

func main() {
	q := &Queue{}
	q.EnqueueSlice([]uint8{1, 6, 8, 90, 80})
	for len(q.vals) != 0 {
		fmt.Println(q.DequeueSlice())
	}
}

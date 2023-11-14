package main

import "fmt"

type Node struct {
	val  uint8
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(inValue uint8) {
	node := &Node{val: inValue}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) PriorityEnqueue(inValue uint8) {
	node := &Node{val: inValue}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		prev_head := q.head
		q.head = node
		q.head.next = prev_head
	}
	q.size++
}

func (q *Queue) EarlyDequeue(dqValue uint8) bool {
	if q.size == 0 {
		return false
	}

	if q.head.val == dqValue {
		q.head = q.head.next
		q.size--
		return true
	}

	prev := q.head
	curr := q.head.next
	for curr != nil && curr.val != dqValue {
		prev = curr
		curr = curr.next
	}

	if curr != nil {
		prev.next = curr.next
		q.size--
		if q.tail == curr {
			q.tail = prev
		}
		return true
	}

	return false
}

func (q *Queue) Dequeue() uint8 {
	returnValue := q.head.val
	q.head = q.head.next
	q.size--
	return returnValue
}

func main() {
	q := &Queue{}
	q.Enqueue(9)
	q.Enqueue(10)
	q.Enqueue(36)
	q.Enqueue(5)
	if !q.EarlyDequeue(5) {
		fmt.Println("Value not found")
	}
	q.PriorityEnqueue(37)
	q.PriorityEnqueue(254)
	for q.size != 0 {
		fmt.Println(q.Dequeue())
	}
}

package main

import "fmt"

//Реализация на зацикленном массиве

type Queue struct {
	Data []int
	Head int
	Tail int
}

func newQueue(size int) Queue {
	data := make([]int, size+1)
	return Queue{Data: data, Head: 0, Tail: 0}
}

func (q *Queue) isEmpty() bool {
	if q.Head == q.Tail {
		return true
	}
	return false
}

func (q *Queue) increaseSize() {
	newQ := newQueue(len(q.Data) * 2)
	for i := 0; !q.isEmpty(); i++ {
		newQ.push_back(q.pop_front())
	}
	*q = newQ
}

func (q *Queue) push_back(x int) {
	if (q.Tail+1 == q.Head) || (q.Tail == len(q.Data)-1 && q.Head == 0) {
		q.increaseSize()
	}
	q.Data[q.Tail] = x
	q.Tail++
	if q.Tail == len(q.Data) {
		q.Tail = 0
	}
}

func (q *Queue) pop_front() int {
	if q.isEmpty() {
		fmt.Println("{Empty}")
		return 0
	}
	num := q.Data[q.Head]
	q.Data[q.Head] = 0
	q.Head++
	if q.Head == len(q.Data) {
		q.Head = 0
	}
	return num
}

func (q *Queue) peek() int {
	var last_index int
	if q.Tail == 0 {
		last_index = len(q.Data) - 1
	} else {
		last_index = q.Tail - 1
	}
	return q.Data[last_index]
}

func main() {
	myQueue := newQueue(3)
	myQueue.push_back(1)
	myQueue.push_back(2)
	myQueue.push_back(3)
	myQueue.push_back(4)
	fmt.Println(myQueue)
}

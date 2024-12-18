package main

import (
	"Go-Queue-DoubleStack/stack"
	"fmt"
)

type Queue struct {
	Size   int
	Input  stack.Stack
	Output stack.Stack
}

func NewQueue() Queue {
	input := stack.NewStack(make([]int, 2))
	output := stack.NewStack(make([]int, 2))
	return Queue{Size: 2, Input: input, Output: output}
}

func (q *Queue) IsEmpty() bool {
	return q.Input.IsEmpty() && q.Output.IsEmpty()
}

// Этот метод в реализации не нуждается, т.к. он реализован в стеках
// func (q *Queue) IncreaseSize() {}

func (q *Queue) Push_back(x int) {
	q.Input.Push_back(x)
}

func (q *Queue) Pop_front() int {
	if q.IsEmpty() {
		fmt.Println("{Empty}")
		return 0
	}
	if q.Output.IsEmpty() {
		for !q.Input.IsEmpty() {
			q.Output.Push_back(q.Input.Pop_back())
		}
	}
	num := q.Output.Pop_back()
	return num
}

func (q *Queue) Peek() int {
	return q.Input.Top()
}

func main() {
	myQueue := NewQueue()
	myQueue.Push_back(1)
	myQueue.Push_back(2)
	myQueue.Push_back(3)
	myQueue.Push_back(4)
	fmt.Println(myQueue.Peek())
}

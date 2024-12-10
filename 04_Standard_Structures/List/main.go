package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Begin *Node
}

func NewList() *List {
	return &List{Begin: nil}
}

func (l *List) push_front(x int) {
	newNode := &Node{Val: x, Next: l.Begin}
	l.Begin = newNode
}

func (l *List) push_back(x int) {
	newNode := &Node{Val: x, Next: nil}
	current := l.Begin
	if current == nil {
		l.Begin = newNode
	} else {
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *List) print() {
	current := l.Begin
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func (l *List) copyEverySecond() List {
	modList := NewList()
	current := l.Begin
	for current != nil {
		modList.push_back(current.Val)
		current = current.Next.Next
	}
	return *modList
}

func (l *List) getSize() (counter int) {
	current := l.Begin
	for current != nil {
		counter++
		current = current.Next
	}
	return
}

func (l *List) toArray() []int {
	arr := make([]int, l.getSize())
	current := l.Begin
	for i := 0; current != nil; i++ {
		arr[i] = current.Val
		current = current.Next
	}
	return arr
}

func (l *List) removeAfter(x *Node) {
	if x.Next != nil && x != nil {
		x.Next = x.Next.Next
	}
}

func (l *List) removeFirst() {
	l.Begin = l.Begin.Next
}

func (l *List) insertAfter(x *Node, val int) {
	if x != nil {
		newNode := Node{Val: val, Next: x.Next}
		x.Next = &newNode
	}
}

func (l *List) filterDivisible(x int) {
	l.push_front(0)
	current := l.Begin
	for current.Next != nil {
		if current.Next.Val%x == 0 {
			l.removeAfter(current)
		} else {
			current = current.Next
		}
	}
	l.removeFirst()
}

func (l *List) getAt(index int) Node {
	current := l.Begin
	for i := 0; i != index; i++ {
		current = current.Next
	}
	return *current
}

func fromArray(arr []int) (l List) {
	for i := len(arr) - 1; i >= 0; i-- {
		l.push_front(arr[i])
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	myList := fromArray(arr)
	myList.print()
}

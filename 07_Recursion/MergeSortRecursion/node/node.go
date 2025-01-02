package node

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

func (l *List) Push_front(x int) {
	newNode := &Node{Val: x, Next: l.Begin}
	l.Begin = newNode
}

func (l *List) Push_back(x int) {
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

func (l *List) Print() {
	current := l.Begin
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func (l *List) CopyEverySecond() List {
	modList := NewList()
	current := l.Begin
	for current != nil {
		modList.Push_back(current.Val)
		current = current.Next.Next
	}
	return *modList
}

func (l *List) GetSize() (counter int) {
	current := l.Begin
	for current != nil {
		counter++
		current = current.Next
	}
	return
}

func (l *List) ToArray() []int {
	arr := make([]int, l.GetSize())
	current := l.Begin
	for i := 0; current != nil; i++ {
		arr[i] = current.Val
		current = current.Next
	}
	return arr
}

func (l *List) RemoveAfter(x *Node) {
	if x.Next != nil && x != nil {
		x.Next = x.Next.Next
	}
}

func (l *List) RemoveFirst() {
	l.Begin = l.Begin.Next
}

func (l *List) InsertAfter(x *Node, val int) {
	if x != nil {
		newNode := Node{Val: val, Next: x.Next}
		x.Next = &newNode
	}
}

func (l *List) FilterDivisible(x int) {
	l.Push_front(0)
	current := l.Begin
	for current.Next != nil {
		if current.Next.Val%x == 0 {
			l.RemoveAfter(current)
		} else {
			current = current.Next
		}
	}
	l.RemoveFirst()
}

func (l *List) GetAt(index int) Node {
	current := l.Begin
	for i := 0; i != index; i++ {
		current = current.Next
	}
	return *current
}

func FromArray(arr []int) (l List) {
	for i := len(arr) - 1; i >= 0; i-- {
		l.Push_front(arr[i])
	}
	return
}

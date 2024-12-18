package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func newNode(x int) Node {
	return Node{value: x, prev: nil, next: nil}
}

type DLList struct {
	begin *Node
	end   *Node
}

func newDLList() *DLList {
	return &DLList{begin: nil, end: nil}
}

func (dll *DLList) push_front(x int) {
	node := newNode(x)
	if dll.begin == nil {
		dll.begin = &node
		dll.end = &node
		return
	}
	dll.begin.prev = &node
	node.next = dll.begin
	dll.begin = &node
}

func (dll *DLList) push_back(x int) {
	node := newNode(x)
	if dll.end == nil {
		dll.begin = &node
		dll.end = &node
		return
	}
	dll.end.next = &node
	node.prev = dll.end
	dll.end = &node
}

func (dll *DLList) print() {
	if dll.begin == nil {
		fmt.Println("{Empty}")
		return
	}
	current := dll.begin
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func (dll *DLList) getSize() (size int) {
	current := dll.begin
	for current != nil {
		size++
		current = current.next
	}
	return
}

func (dll *DLList) toArray() []int {
	dllArray := make([]int, dll.getSize())
	current := dll.begin
	for i := 0; current != nil; i++ {
		dllArray[i] = current.value
		current = current.next
	}
	return dllArray
}

func (dll *DLList) remove(x *Node) {
	if x.prev == nil {
		dll.pop_front()
		return
	}
	if x.next == nil {
		dll.pop_back()
		return
	}
	x.next.prev = x.prev
	x.prev.next = x.next
}

func (dll *DLList) pop_front() {
	if dll.begin == nil {
		return
	}
	if dll.begin == dll.end {
		dll.begin = nil
		dll.end = nil
		return
	}
	dll.begin = dll.begin.next
	dll.begin.prev = nil
}

func (dll *DLList) pop_back() {
	if dll.end == nil {
		return
	}
	if dll.begin == dll.end {
		dll.begin = nil
		dll.end = nil
		return
	}
	dll.end = dll.end.prev
	dll.end.next = nil
}

func (dll *DLList) insertAfter(x *Node, val int) {
	node := newNode(val)
	if x.next == nil {
		dll.push_back(val)
		return
	}
	node.prev = x
	node.next = x.next
	x.next = &node
	node.next.prev = &node
}

func (dll *DLList) getAt(index int) Node {
	if index < 0 {
		return Node{}
	}
	current := dll.begin
	for i := 0; i < index; i++ {
		if current == nil {
			return Node{}
		}
		current = current.next
	}
	if current == nil {
		return Node{}
	}
	return *current
}

func main() {
	myDLList := newDLList()
	myDLList.push_back(2)
	myDLList.push_front(1)
	myDLList.print()
	myDLList.pop_back()
	myDLList.pop_front()
	myDLList.pop_front()
	myDLList.print()
}

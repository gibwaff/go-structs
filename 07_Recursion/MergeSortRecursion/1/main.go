package main

import (
	"local/node"
)

// func MergeTwoListsByCycle(a, b *node.Node) *node.List {
// 	res := node.NewList()
// 	for a != nil || b != nil {
// 		if a == nil {
// 			res.Push_back(b.Val)
// 			b = b.Next
// 		} else if b == nil {
// 			res.Push_back(a.Val)
// 			a = a.Next
// 		} else if a.Val > b.Val {
// 			res.Push_back(b.Val)
// 			b = b.Next
// 		} else {
// 			res.Push_back(a.Val)
// 			a = a.Next
// 		}
// 	}
// 	return res
// }

func MergeTwoLists(a, b *node.Node) *node.List {
	res := node.NewList()
	res = MergeTwoListsRecursion(a, b, res)
	return res
}

func MergeTwoListsRecursion(a, b *node.Node, res *node.List) *node.List {
	if a != nil || b != nil {
		if a == nil {
			res.Push_back(b.Val)
			b = b.Next
		} else if b == nil {
			res.Push_back(a.Val)
			a = a.Next
		} else if a.Val > b.Val {
			res.Push_back(b.Val)
			b = b.Next
		} else {
			res.Push_back(a.Val)
			a = a.Next
		}
		res = MergeTwoListsRecursion(a, b, res)
	}
	return res
}

func main() {
	a1 := node.FromArray([]int{1, 3, 5})
	a2 := node.FromArray([]int{2, 2, 8, 10})
	MergeTwoLists(a1.Begin, a2.Begin).Print()
}

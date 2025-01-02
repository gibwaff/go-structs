package main

import (
	"local/node"
)

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

func MergeThreeLists(a, b, c *node.Node) *node.List {
	res := MergeTwoLists(a, b)
	res = MergeTwoLists(c, res.Begin)
	return res
}

func main() {
	a1 := node.FromArray([]int{1, 3, 5, 10})
	a2 := node.FromArray([]int{2, 2, 8})
	a3 := node.FromArray([]int{1, 4, 7, 7})
	MergeThreeLists(a1.Begin, a2.Begin, a3.Begin).Print()
}

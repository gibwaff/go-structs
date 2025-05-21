package main

import (
	"fmt"
)

type Contact struct {
	Name   string
	Number string
}

type BNode struct {
	Children map[rune]*BNode
	IsEnd    bool
	Contact  Contact
}

func NewNode() *BNode {
	return &BNode{Children: make(map[rune]*BNode)}
}

type BTree struct {
	Node *BNode
}

func NewBTree() *BTree {
	return &BTree{Node: NewNode()}
}

func (bt *BTree) Add(s Contact) {
	node := bt.Node
	for _, ch := range s.Name {
		if _, exists := node.Children[ch]; !exists {
			node.Children[ch] = NewNode()
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
	node.Contact = s
}

func (bt *BTree) Contains(name string) bool {
	node := bt.Node
	for _, ch := range name {
		if _, exists := node.Children[ch]; !exists {
			return false
		}
		node = node.Children[ch]
	}
	return node.IsEnd
}

func (bt *BTree) CountStartsWith(pref string) int {
	node := bt.Node
	for _, ch := range pref {
		if _, exists := node.Children[ch]; !exists {
			return 0
		}
		node = node.Children[ch]
	}
	return countContacts(node)
}

func countContacts(node *BNode) int {
	count := 0
	if node.IsEnd {
		count++
	}
	for _, child := range node.Children {
		count += countContacts(child)
	}
	return count
}

func (bt *BTree) StartsWith(pref string) []Contact {
	node := bt.Node
	for _, ch := range pref {
		if _, exists := node.Children[ch]; !exists {
			return nil
		}
		node = node.Children[ch]
	}
	var result []Contact
	collectContacts(node, &result)
	return result
}

func collectContacts(node *BNode, result *[]Contact) {
	if node.IsEnd {
		*result = append(*result, node.Contact)
	}
	for _, child := range node.Children {
		collectContacts(child, result)
	}
}

func main() {
	bt := NewBTree()

	bt.Add(Contact{"Alice", "123"})
	bt.Add(Contact{"Alicia", "456"})
	bt.Add(Contact{"Bob", "789"})

	fmt.Println(bt.Contains("Alice"))      // true
	fmt.Println(bt.Contains("Alic"))       // false
	fmt.Println(bt.CountStartsWith("Ali")) // 2
	fmt.Println(bt.StartsWith("Ali"))      // [{Alice 123} {Alicia 456}]
	fmt.Println(bt.StartsWith("B"))        // [{Bob 789}]
}

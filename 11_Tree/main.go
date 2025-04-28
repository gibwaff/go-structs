package main

import "fmt"

type Virus struct{}

type Node struct {
	virus Virus
	v     []*Node
}

func NewNode() *Node {
	return &Node{
		virus: Virus{},
		v:     make([]*Node, 0),
	}
}

func NewNodeWithVirus(virus Virus) *Node {
	return &Node{
		virus: virus,
		v:     make([]*Node, 0),
	}
}

func (n *Node) GetVirus() Virus {
	return n.virus
}

//Задание 1
func FromList(root *Node) int {
	if root == nil {
		return 0
	}

	maxDepth := 0

	for _, child := range root.v {
		depth := FromList(child)
		if depth > maxDepth {
			maxDepth = depth
		}
	}

	return maxDepth + 1
}

//Задание 2
func FromListIndexes(index []int, elements []Virus) *Node {
	if len(index) != len(elements) {
		panic("index and elements must be of same length")
	}

	nodes := make([]*Node, len(elements))
	for i, virus := range elements {
		nodes[i] = &Node{virus: virus}
	}

	var root *Node
	for i := 0; i < len(index); i++ {
		if index[i] == -1 {
			root = nodes[i]
		} else {
			if index[i] < 0 || index[i] >= len(nodes) {
				panic(fmt.Sprintf("invalid parent index %d for node %d", index[i], i))
			}
			nodes[index[i]].v = append(nodes[index[i]].v, nodes[i])
		}
	}

	if root == nil {
		panic("no root found (no node with index -1)")
	}

	return root
}

// Задание 3
func AllOnCurrDepth(root *Node, generation int) []*Node {
	var result []*Node

	var dfs func(node *Node, depth int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}

		if depth == generation {
			result = append(result, node)
		}

		if depth < generation {
			for _, child := range node.v {
				dfs(child, depth+1)
			}
		}
	}

	dfs(root, 0)

	return result
}

// Задание 4
func LCA(root, first, second *Node) *Node {
	path1 := findPath(root, first)
	path2 := findPath(root, second)

	var parent *Node
	minLength := len(path1)
	if len(path2) < minLength {
		minLength = len(path2)
	}

	for i := 0; i < minLength; i++ {
		if path1[i] == path2[i] {
			parent = path1[i]
		} else {
			break
		}
	}

	return parent
}

func findPath(root, target *Node) []*Node {
	var path []*Node
	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == nil {
			return false
		}

		path = append(path, node)

		if node == target {
			return true
		}

		for _, child := range node.v {
			if dfs(child) {
				return true
			}
		}

		path = path[:len(path)-1]
		return false
	}

	dfs(root)
	return path
}

func main() {
	virus1 := Virus{}
	virus2 := Virus{}
	virus3 := Virus{}
	virus4 := Virus{}

	elements := []Virus{virus1, virus2, virus3, virus4}
	index := []int{-1, 0, 0, 1}

	root := FromListIndexes(index, elements)

	fmt.Println(AllOnCurrDepth(root, 1))
}

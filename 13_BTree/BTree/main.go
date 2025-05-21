package main

import "fmt"

type BNode struct {
	Children [10]*BNode
	IsEnd    bool
}

func NewBNode() *BNode {
	return &BNode{IsEnd: false}
}

type BTree struct {
	Node *BNode
}

func NewBTree() *BTree {
	return &BTree{Node: NewBNode()}
}

func (bt *BTree) Add(n int) {
	node := bt.Node
	digits := ToDigits(n)
	for _, digit := range digits {
		if node.Children[digit] == nil {
			node.Children[digit] = NewBNode()
		}
		node = node.Children[digit]
	}
	node.IsEnd = true
}

func (bt *BTree) Contains(n int) bool {
	node := bt.Node
	digits := ToDigits(n)
	for _, digit := range digits {
		if node.Children[digit] == nil {
			return false
		}
		node = node.Children[digit]
	}
	return node.IsEnd
}

func ToDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits

}

//выводит +1 уровень с учётом корня
func (bt *BTree) GetMaxHeight() int {
	return getHeight(bt.Node)
}

func getHeight(node *BNode) int {
	if node == nil {
		return 0
	}
	maxDepth := 0
	for _, child := range node.Children {
		if child != nil {
			depth := getHeight(child)
			maxDepth = max(maxDepth, depth)
		}
	}
	return maxDepth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (bt *BTree) GetSorted() []int {
	var result []int
	var dfs func(node *BNode, path []int)

	dfs = func(node *BNode, path []int) {
		if node.IsEnd {
			// Преобразуем путь из цифр в число и запоминаем
			num := 0
			for _, digit := range path {
				num = num*10 + digit
			}
			result = append(result, num)
		}

		// Проходим по детям от 0 до 9 это и будет сортнутый массив
		for digit, child := range node.Children {
			if child != nil {
				dfs(child, append(path, digit)) // Рекурсивный вызов с добавленной цифрой
			}
		}
	}

	dfs(bt.Node, []int{})
	return result
}

func main() {
	MyBTree := NewBTree()
	MyBTree.Add(100)
	MyBTree.Add(221)
	MyBTree.Add(139)
	MyBTree.Add(317)
	fmt.Println(MyBTree.GetSorted())
}

package gobinarytree

import "fmt"

type Node struct {
	X      int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewBinaryTree(x int) *Node {
	return &Node{X: x}
}

func (n *Node) CreateNode(x int) {
	if n == nil {
		return
	}

	if x <= n.X {
		if n.Left == nil {
			n.Left = &Node{X: x, Parent: n}
		} else {
			n.Left.CreateNode(x)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{X: x, Parent: n}
		} else {
			n.Right.CreateNode(x)
		}
	}

	if !n.IsBalanced() {
		n = n.BalanceTree()
	}
}

func (n *Node) PrintTree() {
	if n == nil {
		return
	}
	n.Left.PrintTree()
	fmt.Print(n.X, " ")
	n.Right.PrintTree()
}

func (n *Node) TreeInArray() []int {
	return n.treeArrayRecursion()
}

func (n *Node) treeArrayRecursion() []int {
	if n == nil {
		return []int{}
	}
	arr := n.Left.treeArrayRecursion()
	arr = append(arr, n.X)
	arr = append(arr, n.Right.treeArrayRecursion()...)

	return arr
}

func SortedArrayInTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2
	NewTree := NewBinaryTree(arr[mid])

	NewTree.Left = SortedArrayInTree(arr[:mid])
	if NewTree.Left != nil {
		NewTree.Left.Parent = NewTree
	}

	NewTree.Right = SortedArrayInTree(arr[mid+1:])
	if NewTree.Right != nil {
		NewTree.Right.Parent = NewTree
	}

	return NewTree
}

func ArrayInTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	arr_Tree := NewBinaryTree(arr[0])
	for i := 1; i < len(arr); i++ {
		arr_Tree.CreateNode(arr[i])
	}
	sorted_arr := arr_Tree.TreeInArray()

	NewTree := SortedArrayInTree(sorted_arr)

	return NewTree
}

func (n *Node) Find(x int) *Node {
	if n == nil {
		return nil
	}

	if x == n.X {
		return n
	} else if x > n.X {
		return n.Right.Find(x)
	} else {
		return n.Left.Find(x)
	}
}

func (n *Node) Min() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.Min()
}

func (n *Node) Max() *Node {
	if n.Right == nil {
		return n
	}
	return n.Right.Max()
}

func (n *Node) BalanceTree() *Node {
	buf_arr := n.TreeInArray()
	NewTree := SortedArrayInTree(buf_arr)
	return NewTree
}

func (n *Node) Delete(x int) {
	NodeToDelete := n.Find(x)
	if NodeToDelete == nil {
		fmt.Println("Узел не найден")
		return
	}

	// Если удаляемый узел корневой
	if NodeToDelete.Parent == nil {
		var MinRight *Node
		// Ищем минимальный в правом поддереве если там есть потомки
		if NodeToDelete.Right != nil {
			MinRight = NodeToDelete.Right.Min()
			//Если МинРайт правый потомок корня
			if MinRight == NodeToDelete.Right {
				NodeToDelete.X = MinRight.X
				NodeToDelete.Right = MinRight.Right
				if NodeToDelete.Right != nil {
					NodeToDelete.Right.Parent = NodeToDelete
				}
			} else {
				NodeToDelete.X = MinRight.X
				MinRight.Parent.Left = MinRight.Right
				if MinRight.Right != nil {
					MinRight.Right.Parent = MinRight.Parent
				}
			}

		} else if NodeToDelete.Left != nil {
			NodeToDelete.X = NodeToDelete.Left.X
			NodeToDelete.Right = NodeToDelete.Left.Right
			if NodeToDelete.Right != nil {
				NodeToDelete.Right.Parent = NodeToDelete
			}
			NodeToDelete.Left = NodeToDelete.Left.Left
			if NodeToDelete.Left != nil {
				NodeToDelete.Left.Parent = NodeToDelete
			}

		} else {
			// Если дерево состоит только из одного узла
			*n = Node{}
			return
		}

		//Узел с двумя детьми
	} else if NodeToDelete.Left != nil && NodeToDelete.Right != nil {
		MinRight := NodeToDelete.Right.Min()
		//Если Минрайт потомок УдУз (удаляемого узла)
		if MinRight == NodeToDelete.Right {
			NodeToDelete.X = MinRight.X
			NodeToDelete.Right = MinRight.Right
			if NodeToDelete.Right != nil {
				NodeToDelete.Right.Parent = NodeToDelete
			}
		} else {
			NodeToDelete.X = MinRight.X
			MinRight.Parent.Left = MinRight.Right
			if MinRight.Right != nil {
				MinRight.Right.Parent = MinRight.Parent
			}
		}

	} else {
		if NodeToDelete.Left != nil { //Если есть левый потомок
			if NodeToDelete.Parent.Left == NodeToDelete {
				NodeToDelete.Parent.Left = NodeToDelete.Left
			} else {
				NodeToDelete.Parent.Right = NodeToDelete.Left
			}
			NodeToDelete.Left.Parent = NodeToDelete.Parent

		} else if NodeToDelete.Right != nil { //Если есть правый потомок
			if NodeToDelete.Parent.Left == NodeToDelete {
				NodeToDelete.Parent.Left = NodeToDelete.Right
			} else {
				NodeToDelete.Parent.Right = NodeToDelete.Right
			}
			NodeToDelete.Right.Parent = NodeToDelete.Parent

		} else { //Если нет детей
			if NodeToDelete.Parent.Left == NodeToDelete {
				NodeToDelete.Parent.Left = nil
			} else {
				NodeToDelete.Parent.Right = nil
			}
		}
	}

	//балансировка
	if !n.IsBalanced() {
		n = n.BalanceTree()
	}
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}
	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (n *Node) CountNodes() int {
	if n == nil {
		return 0
	}
	return n.Left.CountNodes() + n.Right.CountNodes() + 1
}

func (n *Node) IsBalanced() bool {
	if n == nil {
		return true
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return n.Left.IsBalanced() && n.Right.IsBalanced()
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (n *Node) Clear() {
	*n = Node{}
}

func main() {
	MyBinaryTree := NewBinaryTree(1)

	MyBinaryTree.CreateNode(2)
	MyBinaryTree.CreateNode(3)
	MyBinaryTree.CreateNode(4)
	MyBinaryTree.CreateNode(5)
	MyBinaryTree.CreateNode(5)
	MyBinaryTree.CreateNode(5)

	MyBinaryTree.PrintTree()
}

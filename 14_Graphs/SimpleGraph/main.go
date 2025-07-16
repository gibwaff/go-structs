package main

import "fmt"

//Задание 1
type Employee struct {
	id           int
	importance   int
	subordinates []int
}

func NewEmployee(id, importance int, subordinates []int) Employee {
	return Employee{id: id, importance: importance, subordinates: subordinates}
}

type Collective struct {
	employers []Employee
}

func NewCollective() Collective {
	return Collective{employers: []Employee{}}
}

func (c *Collective) AddEmployee(id, importance int, subordinates []int) {
	worker := NewEmployee(id, importance, subordinates)
	c.employers = append(c.employers, worker)
}

func (c *Collective) GetByID(id int) Employee {
	for i := range c.employers {
		if c.employers[i].id == id {
			return c.employers[i]
		}
	}
	return Employee{}
}

func (c *Collective) PrintEmployees() {
	queue := NewQueue()
	seem := []Employee{}
	queue.AddEmployee(c.employers[0])
	seem = append(seem, c.employers[0])
	for !queue.IsEmpty() {
		worker := queue.GetEmployee()
		for _, id := range worker.subordinates {
			selected_worker := c.GetByID(id)
			queue.AddEmployee(selected_worker)
			if !IsHave(seem, selected_worker) {
				seem = append(seem, selected_worker)
			}
		}
	}
	for i := range seem {
		fmt.Println(seem[i])
	}
}

func IsHave(arr []Employee, n Employee) bool {
	for i := range arr {
		if arr[i].id == n.id {
			return true
		}
	}
	return false
}

type Queue struct {
	workers []Employee
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) AddEmployee(worker Employee) {
	q.workers = append(q.workers, worker)
}

func (q *Queue) IsEmpty() bool {
	return len(q.workers) == 0
}

func (q *Queue) GetEmployee() Employee {
	neworkers := make([]Employee, len(q.workers)-1)
	for i := 1; i < len(q.workers); i++ {
		neworkers[i-1] = q.workers[i]
	}
	worker := q.workers[0]
	q.workers = neworkers
	return worker
}

func (q *Queue) IsInQueue(n Employee) bool {
	for _, worker := range q.workers {
		if n.id == worker.id {
			return true
		}
	}
	return false
}

func (c *Collective) GetImportance(ID int) (Importance int) {
	selected_worker := c.GetByID(ID)
	queue := NewQueue()
	seem := []Employee{}
	queue.AddEmployee(selected_worker)
	seem = append(seem, selected_worker)
	for !queue.IsEmpty() {
		worker := queue.GetEmployee()
		for _, id := range worker.subordinates {
			selected_worker := c.GetByID(id)
			queue.AddEmployee(selected_worker)
			if !IsHave(seem, selected_worker) {
				seem = append(seem, selected_worker)
			}
		}
	}
	for i := range seem {
		Importance += seem[i].importance
	}
	return
}

//Задание 2
type Node struct {
	val      string
	neigbors []Node
}

func NewNode() Node {
	return Node{}
}

func cloneGraph(node Node) Node {
	res_graph := NewNode()
	res_graph.val = node.val
	res_graph.neigbors = node.neigbors
	seem := []Node{}
	queue := NewNQueue()
	for _, el := range node.neigbors {
		queue.AddNode(el)
	}
	for !queue.IsEmpty() {
		cur_node := queue.GetNode()
		if !IsInArr(seem, cur_node) {
			seem = append(seem, cur_node)
		}
		for _, el := range cur_node.neigbors {
			queue.AddNode(el)
		}
	}
	return res_graph
}

type NQueue struct {
	container []Node
}

func NewNQueue() NQueue {
	return NQueue{}
}

func (nq *NQueue) AddNode(n Node) {
	nq.container = append(nq.container, n)
}

func (nq *NQueue) GetNode() Node {
	node := nq.container[len(nq.container)-1]
	new_cont := []Node{}
	for i := 1; i < len(nq.container); i++ {
		new_cont = append(new_cont, nq.container[i])
	}
	nq.container = new_cont
	return node
}

func (nq *NQueue) IsEmpty() bool {
	return len(nq.container) == 0
}

func (nq *NQueue) IsInNQueue(node Node) bool {
	for _, el := range nq.container {
		if el.val == node.val {
			return true
		}
	}
	return false
}

func IsInArr(arr []Node, node Node) bool {
	for _, el := range arr {
		if node.val == el.val {
			return true
		}
	}
	return false
}

func main() {
	//Задание 1
	MyWorkers := NewCollective()
	MyWorkers.AddEmployee(1, 15, []int{2, 3, 6})
	MyWorkers.AddEmployee(2, 10, []int{3, 4, 5})
	MyWorkers.AddEmployee(3, 5, []int{4, 5})
	MyWorkers.AddEmployee(4, 4, []int{5})
	MyWorkers.AddEmployee(5, 3, []int{6})
	MyWorkers.AddEmployee(6, 1, []int{})

	fmt.Println(MyWorkers.GetImportance(2))
	// 23, т.к. суммарно в подчинении у id2 : [3, 4, 5, 6]

	//Задание 2
	MyNodes2 := []Node{{val: "second"}, {val: "third"}}
	MyNodes := []Node{{val: "sas", neigbors: MyNodes2}, {val: "sss", neigbors: MyNodes2}, {val: "key"}}
	MyNode := cloneGraph(Node{val: "swag", neigbors: MyNodes})
	fmt.Println(MyNode)
}

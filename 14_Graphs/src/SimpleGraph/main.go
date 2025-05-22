package main

import "fmt"

type SimpleGraph struct {
}

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

type Node struct {
	val      string
	neigbors []Node
}

func main() {
	MyWorkers := NewCollective()
	MyWorkers.AddEmployee(1, 15, []int{2, 3})
	MyWorkers.AddEmployee(2, 10, []int{3})
	MyWorkers.AddEmployee(3, 5, []int{})

	MyWorkers.PrintEmployees()
}

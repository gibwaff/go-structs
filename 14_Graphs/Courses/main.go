package main

import "fmt"

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Строим граф зависимостей
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		a, b := pair[0], pair[1]
		graph[a] = append(graph[a], b)
	}

	// Статусы курсов: 0 — не посещён, 1 — в процессе (в стеке), 2 — завершён
	visited := make([]int, numCourses)

	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if visited[course] == 1 {
			return true
		}
		if visited[course] == 2 {
			return false
		}

		visited[course] = 1
		for _, prereq := range graph[course] {
			if hasCycle(prereq) {
				return true
			}
		}
		visited[course] = 2
		return false
	}

	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanFinish(2, [][]int{{1, 0}}))         // true
	fmt.Println(CanFinish(2, [][]int{{1, 0}, {0, 1}})) // false (цикл)

}

package main

import (
	"fmt"
	"math"
)

// Задание 1

func isProjectSuccess(project map[string][]string) bool {
	if len(project) <= 1 {
		return false
	}

	edgeCount := 0
	degrees := make(map[string]int)

	for node, neighbors := range project {
		degrees[node] += len(neighbors)
		edgeCount += len(neighbors)
	}

	// Делим пополам, потому что мосты в обе стороны
	if edgeCount/2 <= 1 {
		return false
	}

	// Проверка связности графа
	visited := make(map[string]bool)
	var dfs func(string)
	dfs = func(node string) {
		visited[node] = true
		for _, neighbor := range project[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Запускаем DFS с любого узла
	for node := range project {
		dfs(node)
		break
	}

	// Проверка, что все узлы достижимы
	for node := range project {
		if len(project[node]) > 0 && !visited[node] {
			return false
		}
	}

	// Подсчёт количества вершин с нечетной степенью
	odd := 0
	for _, deg := range degrees {
		if deg%2 != 0 {
			odd++
		}
	}
	return odd == 0 || odd == 2
}

//Задание 2

type Point struct {
	x, y int
}

func shortestPathDuration(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	rows := len(grid)
	cols := len(grid[0])

	// Веса для каждого узла
	dist := make([][]int, rows)
	visited := make([][]bool, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		visited[i] = make([]bool, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[0][0] = 0

	// Смещения по направлениям: вверх, вниз, влево, вправо
	dir := [][]int{
		{0, 1}, {1, 0}, {-1, 0}, {0, -1},
	}

	for {
		minDist := math.MaxInt32
		var current Point
		found := false

		// Находим наименее затратную еще не посещенную вершину
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if !visited[i][j] && dist[i][j] < minDist {
					minDist = dist[i][j]
					current = Point{i, j}
					found = true
				}
			}
		}

		if !found {
			break
		}

		x, y := current.x, current.y
		visited[x][y] = true

		for _, d := range dir {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols {
				newDist := dist[x][y] + grid[nx][ny]
				if newDist < dist[nx][ny] {
					dist[nx][ny] = newDist
				}
			}
		}
	}

	return dist[rows-1][cols-1]
}

func main() {
	//Задание 1
	project1 := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "C"},
		"C": {"A", "B"},
	}
	fmt.Println(isProjectSuccess(project1)) // true

	//Задание 2
	grid := [][]int{
		{0, 1, 5, 5, 1, 1, 2},
		{1, 1, 5, 5, 1, 1, 1},
		{1, 1, 8, 2, 2, 2, 1},
		{8, 8, 8, 2, 2, 1, 1},
		{8, 2, 2, 8, 1, 1, 0},
	}
	fmt.Println(shortestPathDuration(grid)) // 17
}

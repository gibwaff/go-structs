package main

import "fmt"

//Задание 1
func paint(image [][]int, row, col, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	startColor := image[row][col]
	if startColor == newColor {
		return image
	}

	rows := len(image)
	cols := len(image[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		if image[r][c] != startColor {
			return
		}

		image[r][c] = newColor

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	dfs(row, col)
	return image
}

//Задание 2

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}

		grid[r][c] = '0'

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}

func main() {
	//Задание 1
	image := [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
	}

	result := paint(image, 0, 1, 1)

	for _, row := range result {
		fmt.Println(row)
	}

	//Задание 2
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid)) // Вывод: 3
}

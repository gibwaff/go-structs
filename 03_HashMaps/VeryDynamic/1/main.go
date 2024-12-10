package main

import "fmt"

func resize(values []int, newSize int) ([]int, int) {
	newValues := make([]int, newSize)
	length := min(len(values), len(newValues))
	for i := 0; i < length; i++ {
		newValues[i] = values[i]
	}
	return newValues, newSize
}

func main() {
	values := make([]int, 2)
	size := 2
	currentIndex := 0

	add := func(value int) {
		values[currentIndex] = value
		currentIndex++
		if currentIndex == size {
			values, size = resize(values, size*2)
		}
		if currentIndex == size/4 {
			values, size = resize(values, size/2)
		}
	}

	deleteElement := func(index int) {
		for i := index; values[i] != 0; i++ {
			values[i] = values[i+1]
		}
		currentIndex--
		if currentIndex == size/4 {
			values, size = resize(values, size/2)
		}
	}

	add(1)
	add(2)
	add(3)
	add(4)
	fmt.Println(values)
	deleteElement(2)
	deleteElement(2)
	fmt.Println(values)
	deleteElement(1)
	fmt.Println(values)
}

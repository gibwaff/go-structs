package main

import "fmt"

type SuperArray interface {
	getSize() int
	get(position int) int
	set(position, value int)
}

type MyArray struct {
	Array []int
}

func (arr *MyArray) getSize() int {
	return len(arr.Array)
}

func (arr *MyArray) get(position int) int {
	return arr.Array[position]
}

func (arr *MyArray) set(position, value int) {
	arr.Array[position] = value
}

func (arr *MyArray) SelectionSort() {
	for i := 0; i < arr.getSize()-1; i++ {
		pt := i
		for j := i + 1; j < arr.getSize(); j++ {
			if arr.get(j) <= arr.get(pt) {
				pt = j
			}
		}
		c := arr.get(i)
		arr.set(i, arr.get(pt))
		arr.set(pt, c)
	}
}

func main() {
	myArray := MyArray{Array: []int{10, 2, 7, 7, 2, 5, 1, 8}}
	myArray.SelectionSort()
	fmt.Println(myArray)
}

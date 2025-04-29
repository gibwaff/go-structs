package main

import "fmt"

// 1 задание

func CreateHeap(a int) []int {
	return []int{a}
}
func AddToHeap(heap []int, a int) []int {
	heap = append(heap, a)
	for i := len(heap) - 1; a > heap[(i-1)/2]; i = (i - 1) / 2 {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
	}
	return heap
}

// конечная функция - buildHeapFromArray
func ArrayToHeap(arr []int) []int {
	heap := CreateHeap(arr[0])
	for i := 1; i < len(arr); i++ {
		heap = AddToHeap(heap, arr[i])
	}
	return heap
}

func main() {
	MyArray := []int{10, 2, 4, 5, 15, 3}
	MyHeap := ArrayToHeap(MyArray)
	fmt.Println(MyHeap)
}

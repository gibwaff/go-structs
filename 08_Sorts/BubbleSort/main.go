package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	MyArray := []int{2, 14, 15, 8, 3, 3, 0}
	fmt.Println(BubbleSort(MyArray))
}

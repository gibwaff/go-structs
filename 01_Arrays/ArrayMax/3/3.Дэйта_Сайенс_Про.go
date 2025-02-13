package main

import (
	"fmt"
)

func main() {
	var time = []int{9, 4, 1, 8, 7, 9, 4, 1, 8, 7, 8, 7, 18, 3, 13, 6, 5}
	var sum, averageTime float32
	for _, el := range time {
		sum += float32(el)
	}
	averageTime = sum / float32(len(time))
	fmt.Printf("%.2f", averageTime)
}

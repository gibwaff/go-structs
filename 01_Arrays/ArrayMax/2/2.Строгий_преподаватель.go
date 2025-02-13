package main

import (
	"fmt"
	"math"
)

func main() {
	var studentsMistakes = []int{9, 4, 1, 8, 7, 13, 6, 5}
	var mini = math.MaxInt32
	for _, el := range studentsMistakes {
		if el < mini {
			mini = el
		}
	}
	fmt.Println(mini)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{-1392, -1920, -7, -453, -91234}
	var maxi = math.MinInt32
	for _, el := range arr {
		if el > maxi {
			maxi = el
		}
	}
	fmt.Println(maxi)
}

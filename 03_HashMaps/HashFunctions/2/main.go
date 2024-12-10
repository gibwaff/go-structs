package main

import "fmt"

func hashInt(input int) (index int) {
	index = 1
	var numbers []int

	del := 1
	for input >= del {
		del *= 10
	}
	del /= 10
	for del != 0 {
		numbers = append(numbers, (input/del)%10)
		del /= 10
	}

	for _, el := range numbers {
		if el != 0 {
			index *= el
		}
	}

	return
}

func main() {
	var valueArray = []int{234, 204, 1000, 10002, 2300}
	for _, value := range valueArray {
		fmt.Println(hashInt(value))
	}
}

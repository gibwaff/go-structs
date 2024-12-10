package main

import "fmt"

func findPhoneNumber(sortedPhoneNumbers []int, phone int) (index int) {
	index = -1
	left, right := 0, len(sortedPhoneNumbers)-1
	var middle int
	if sortedPhoneNumbers[left] < sortedPhoneNumbers[right] {
		for left <= right {
			middle = (left + right) / 2
			if sortedPhoneNumbers[middle] < phone {
				left = middle + 1
			} else if sortedPhoneNumbers[middle] > phone {
				right = middle - 1
			} else {
				index = middle
				break
			}
		}
	} else {
		for left <= right {
			middle = (left + right) / 2
			if sortedPhoneNumbers[middle] > phone {
				left = middle + 1
			} else if sortedPhoneNumbers[middle] < phone {
				right = middle - 1
			} else {
				index = middle
				break
			}
		}
	}
	return
}

func main() {
	var sortedPhoneNumbers = []int{
		+79630000000,
		+79630000001,
		+79630000002,
		+79630000003,
		+79630000004,
	}
	phone := +79630000003
	fmt.Println(findPhoneNumber(sortedPhoneNumbers, phone))
}

package main

import "fmt"

type Student struct {
	age  int
	name string
}

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

func hashString(input string) (index int) {
	for _, el := range input {
		index += int(el)
	}
	return
}

func hashStudent(input Student) (index int) {
	return hashInt(hashString(input.name)) + input.age
}

func main() {
	var students = []Student{
		{20, "Max"},
		{21, "Andrey"},
		{21, "Ivan"},
		{20, "Andrey"},
	}

	for _, student := range students {
		fmt.Println(hashStudent(student))
	}

}

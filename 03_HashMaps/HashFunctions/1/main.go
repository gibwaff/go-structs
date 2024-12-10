package main

import "fmt"

func hashString(input string) (index int) {
	for _, el := range input {
		index += int(el)
	}
	return
}

func main() {
	var valueArray = []string{"Andrey", "Max", "Igor"}
	for _, value := range valueArray {
		fmt.Println(hashString(value))
	}
}

package main

import "fmt"

func doIKnowThisLanguage(languagesList []string, language string) int {
	left := 0
	right := len(languagesList) - 1
	var middle int

	for left <= right {
		middle = (left + right) / 2
		if languagesList[middle] < language {
			left = middle + 1
		} else if languagesList[middle] > language {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	var languageList = []string{
		"ADA",
		"ALGOL",
		"B",
		"BASIC",
		"C",
		"C++",
	}
	fmt.Println(doIKnowThisLanguage(languageList, "BASIC"))
}

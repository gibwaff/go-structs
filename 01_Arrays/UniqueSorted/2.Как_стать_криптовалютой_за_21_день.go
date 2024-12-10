package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("task.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	date := make(map[string][]int)
	for scanner.Scan() {
		var s []string = strings.Split(scanner.Text(), ":")
		money, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
		}
		_, exists := date[s[0]]
		if exists {
			date[s[0]][0] += money
			date[s[0]][1] += 1
		} else {
			date[s[0]] = []int{money, 1}
		}
	}

	for key, value := range date {
		fmt.Printf("%s: %d\n", key, value[0]/value[1])
	}
}

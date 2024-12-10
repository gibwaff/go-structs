package main

import "fmt"

type KeyValuePair struct {
	key   string
	value string
}

func hashFunction(key string) int {
	runeKey := []rune(key)
	var hash = 1
	for _, el := range runeKey {
		hash *= int(el)
	}
	return hash
}

func resize(values []KeyValuePair, newSize int) ([]KeyValuePair, int) {
	newValues := make([]KeyValuePair, newSize)
	length := min(len(values), len(newValues))
	for i := 0; i < length; i++ {
		newValues[i] = values[i]
	}
	return newValues, newSize
}

func findGoodIndex(entries []KeyValuePair, key string, size int) int {
	hash := hashFunction(key)
	index := hash % size

	for i := 0; i < size; i++ {
		probingIndex := (index + i) % size
		entry := entries[probingIndex]
		if entry.key == "" || entry.key == key {
			return probingIndex
		}
	}
	return -1
}

func main() {
	var entries = make([]KeyValuePair, 8)

	size := 8
	numberOfElements := 0

	add := func(key string, value string) {
		index := findGoodIndex(entries, key, size)
		entries[index] = KeyValuePair{key, value}
		numberOfElements++
		if numberOfElements == size {
			entries, size = resize(entries, size*2)
		}
		if numberOfElements == size/4 && size > 8 {
			entries, size = resize(entries, size/2)
		}
	}
	get := func(key string) string {
		index := findGoodIndex(entries, key, size)
		if index == -1 {
			fmt.Println("Find error")
		}
		entry := entries[index]
		if entry.value == "" {
			fmt.Println("Value is null")
		}
		return entry.value
	}
	deleteKey := func(key string) {
		index := findGoodIndex(entries, key, size)
		entries[index] = KeyValuePair{}
		numberOfElements--
		if numberOfElements == size/4 && size > 8 {
			entries, size = resize(entries, size/2)
		}
	}
	getAllKeys := func() (keys []string) {
		for i := 0; i < size; i++ {
			if entries[i].key != "" {
				keys = append(keys, entries[i].key)
			}
		}
		return
	}
	getAllValues := func() (values []string) {
		for i := 0; i < size; i++ {
			if entries[i].value != "" {
				values = append(values, entries[i].value)
			}
		}
		return
	}

	add("Skill", "Box")
	add("Bob", "Builder")
	add("Bird", "Crow")
	fmt.Println(get("Skill"))
	deleteKey("Skill")
	fmt.Println(getAllKeys())
	fmt.Println(getAllValues())
}

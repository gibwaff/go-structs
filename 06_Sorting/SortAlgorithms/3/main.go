package main

import "fmt"

type AnimalInZoo struct {
	Spice  string
	Weight int
}

type Animal interface {
	getWeight() int
}

func (a *AnimalInZoo) getWeight() int {
	return a.Weight
}

func RaiseArray(arr []AnimalInZoo) []AnimalInZoo {
	new_arr := make([]AnimalInZoo, len(arr)+1)
	for i := range arr {
		new_arr[i] = arr[i]
	}
	return new_arr
}

func Insert(arr []AnimalInZoo, pos int, animal AnimalInZoo) []AnimalInZoo {
	arr = RaiseArray(arr)
	for i := len(arr) - 1; i > pos; i-- {
		arr[i] = arr[i-1]
	}
	arr[pos] = animal
	return arr
}

func ReduceArray(arr []AnimalInZoo) []AnimalInZoo {
	new_arr := make([]AnimalInZoo, len(arr)-1)
	for i := range new_arr {
		new_arr[i] = arr[i]
	}
	return new_arr
}

func RemoveElem(arr []AnimalInZoo, pos int) []AnimalInZoo {
	for i := pos; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr = ReduceArray(arr)
	return arr
}

func InsertionSort(arr []AnimalInZoo) []AnimalInZoo {
	for i := 0; i < len(arr); i++ {
		c := arr[i]
		for j := i; j > -1; j-- {
			if j == 0 {
				arr = RemoveElem(arr, i)
				arr = Insert(arr, j, c)
				break
			} else if arr[j-1].getWeight() <= c.getWeight() {
				arr = RemoveElem(arr, i)
				arr = Insert(arr, j, c)
				break
			}
		}
	}
	return arr
}

func main() {
	MyZoo := []AnimalInZoo{
		{Spice: "Bobr", Weight: 18},
		{Spice: "Manul", Weight: 15},
		{Spice: "Rat", Weight: 3},
		{Spice: "Mice", Weight: 3},
		{Spice: "Tiger", Weight: 84},
		{Spice: "Bear", Weight: 173},
		{Spice: "Butterfly", Weight: 1},
	}

	MyZoo = InsertionSort(MyZoo)
	fmt.Println(MyZoo)

}

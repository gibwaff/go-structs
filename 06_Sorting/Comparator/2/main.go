package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float32
}

func (b *Book) getTitle() string {
	return b.Title
}

func (b *Book) getAuthor() string {
	return b.Author
}

func (b *Book) getYear() int {
	return b.Year
}

func (b *Book) getPrice() float32 {
	return b.Price
}

type Comparable interface {
	compareTo(other Book) int
}

func (b *Book) compareTo(other Book) int {
	if b.getYear() != other.getYear() {
		return -1
	}
	return 0
}

func compareInt(first, second int) int {
	switch {
	case first < second:
		return -1
	case first == second:
		return 0
	case first > second:
		return 1
	}
	return 10
}

func compareString(first, second string) int {
	first_arr := []byte(first)
	second_arr := []byte(second)
	for i := range min(len(first_arr), len(second_arr)) {
		switch {
		case first_arr[i] < second_arr[i]:
			return -1
		case first_arr[i] > second_arr[i]:
			return 1
		}
	}
	switch {
	case len(first_arr) < len(second_arr):
		return -1
	case len(first_arr) == len(second_arr):
		return 0
	case len(first_arr) > len(second_arr):
		return 1
	}
	return 10
}

func compareBooksByAuthorAndTitleAndYear(first, second Book) int {
	if compareString(first.getAuthor(), second.getAuthor()) != 0 {
		return compareString(first.getAuthor(), second.getAuthor())
	} else if compareString(first.getTitle(), second.getTitle()) != 0 {
		return compareString(first.getTitle(), second.getTitle())
	} else {
		return compareInt(first.getYear(), second.getYear())
	}
}

func main() {
	book1 := Book{Title: "Грокаем алгоритмы", Author: "Адитья Бхаргава", Year: 2017, Price: 20.0}
	book2 := Book{Title: "Грокаем алгоритмы", Author: "Адитья Бхаргава", Year: 2021, Price: 25.0}

	fmt.Println(compareBooksByAuthorAndTitleAndYear(book1, book2))
}

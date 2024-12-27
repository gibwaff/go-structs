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

func compareFloat(first, second float32) int {
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

func (b *Book) compareTo(other Book) (t, a, y, p int) {
	t = compareString(b.getAuthor(), other.getAuthor())
	a = compareString(b.getAuthor(), b.getTitle())
	y = compareInt(b.getYear(), other.getYear())
	p = compareFloat(b.getPrice(), other.getPrice())
	return
}

func main() {
	book1 := Book{Title: "Book A", Author: "Author A", Year: 2000, Price: 20.0}
	book2 := Book{Title: "Book B", Author: "Author B", Year: 1995, Price: 25.0}

	fmt.Println(book1.compareTo(book2))
}

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

func compareBooksByPrice(first, second Book) int {
	switch {
	case first.getPrice() < second.getPrice():
		return -1
	case first.getPrice() == second.getPrice():
		return 0
	case first.getPrice() > second.getPrice():
		return 1
	}
	return 10
}

func main() {
	book1 := Book{Title: "Book A", Author: "Author A", Year: 2000, Price: 20.0}
	book2 := Book{Title: "Book B", Author: "Author B", Year: 1995, Price: 25.0}

	fmt.Println(compareBooksByPrice(book1, book2))
}

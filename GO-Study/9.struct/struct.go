package main

import "fmt"

// myint == int
type myint int

// struct
type Book struct {
	title string
	auth  string
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("a type = %T\n", a)

	var book1 Book
	book1.title = "Go"
	book1.auth = "King"

	fmt.Printf("%v\n", book1) // Go King
	changeBook(book1)
	fmt.Printf("%v\n", book1) // Go King

	changebook2(&book1)
	fmt.Printf("%v\n", book1) // Golang King

}

func changeBook(book Book) {
	book.title = "Golang"
}

func changebook2(book *Book) {
	book.title = "Golang"
}

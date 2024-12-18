package main

import "fmt"

type Reader interface {
	Reader()
}

type Writer interface {
	Writer()
}

type Book struct{}

func (book *Book) Reader() {
	fmt.Println("Read a book")
}

func (book *Book) Writer() {
	fmt.Println("Writer a book")
}

func main() {
	// pair <type :Book, value : book{}地址 >
	b := &Book{}

	var r Reader
	// pair <type :Book, value : book{}地址 >
	r = b
	r.Reader()

	var w Writer
	// pair <type :Book, value : book{}地址 >
	w = r.(Writer)
	w.Writer()

	t := w.(Reader)
	t.Reader()
}

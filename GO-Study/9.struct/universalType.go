package main

import "fmt"

// interface{}   universal data type
func f(arg interface{}) {
	fmt.Println("f is called ...")
	fmt.Println(arg)

	// assertion
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type , value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

type Book struct {
	auth string
}

func main() {
	book := Book{"king"}
	f(book)
	f(100)
	f(1.1)
	f("123")
	f(true)
}

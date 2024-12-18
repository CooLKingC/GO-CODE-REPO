package main

import "fmt"

func main() {
	defer fmt.Println("main end1") // FILO
	defer fmt.Println("main end2")

	fmt.Println("go 1 ")
	fmt.Println("go 2 ")

	// return and defer
	returnAndDefer() // return >  defer
}

func returnAndDefer() int {
	defer deferfunc()
	return returnfunc()
}

func returnfunc() int {
	fmt.Println("return")
	return 0
}

func deferfunc() int {
	fmt.Println("defer")
	return 0
}

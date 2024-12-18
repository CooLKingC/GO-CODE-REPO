package main

import "fmt"

func main() {
	Demo(20)
	fmt.Println("continue...")

}

func Demo(i int) {
	var arr [10]int
	// anonymous function
	defer func() {
		err := recover() // recover is used to catch any potential panics that may occur
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 10
}

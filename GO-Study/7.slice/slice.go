package main

import "fmt"

func main() {
	// dynamic array
	arr := []int{1, 2, 3, 4}
	fmt.Printf("arr type is %T\n", arr)
	fmt.Println(arr)

	fmt.Println("====")

	printArr(arr)

	fmt.Println("====")

	for _, value := range arr {
		fmt.Println("value = ", value)
	}
}

func printArr(arr []int) { // pass by reference
	// _ anonymous variable
	for _, value := range arr {
		fmt.Println("value = ", value)
	}
	arr[0] = 100
}

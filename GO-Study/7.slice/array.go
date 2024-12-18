package main

import "fmt"

func main() {
	// fixed length
	var arr1 [10]int

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr2 := [10]int{1, 2, 3, 4}
	printarr(arr2)

	fmt.Println("====")

	for index, value := range arr2 {
		fmt.Println("index = ", index, "value = ", value) // 1,2,3,4,0,0,0,0,0,0
	}

	fmt.Printf("arr1 types = %T\n", arr1)
	fmt.Printf("arr2 types = %T\n", arr2)
}

func printarr(arr [10]int) {
	arr[0] = 123
	for index, value := range arr {
		fmt.Println("index = ", index, "value = ", value) // 123,2,3,4,0,0,0,0,0,0
	}
	arr[1] = 100
}

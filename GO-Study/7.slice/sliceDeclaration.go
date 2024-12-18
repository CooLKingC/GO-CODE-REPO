package main

import "fmt"

func main() {
	// 1   len = 3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	// 2   len = 0  space not allocated yet
	var slice2 []int
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2) // 0 []
	fmt.Println("====")
	// allocate memory space
	slice2 = make([]int, 3)
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2) // 3 [0,0,0]

	// 3   len = 3
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice3 = %v\n", len(slice3), slice3) // 3 [0,0,0]

	// 4  len = 3
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4) // 3 [0,0,0]

	// other
	var slice5 []int
	if slice5 == nil {
		fmt.Println("empty slice")
	} else {
		fmt.Println("non-empty slice")
	}

}

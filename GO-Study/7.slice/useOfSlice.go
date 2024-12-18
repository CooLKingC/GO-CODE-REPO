package main

import "fmt"

/*
  when the amount of append is greater than the cap,
  the cap value will be expanded to twice the current cap
*/

func main() {
	// len = 3 , cap = 5 , [0,0,0]
	numbers := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(numbers), cap(numbers), numbers)

	// append  , len = 4 , [0,0,0,1] , cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(numbers), cap(numbers), numbers)
	// append , len = 5 , [0,0,0,1,2] , cap = 5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(numbers), cap(numbers), numbers)
	// append , len = 6 , [0,0,0,1,2,3] , cap = 10
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("====")
	// len , cap = 3 , [0,0,0]
	ns := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(ns), cap(ns), ns)
	ns = append(ns, 1)
	// len = 4 , cap = 6 [0,0,0,1]
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(ns), cap(ns), ns)

	fmt.Println("====slicing=====")
	s := []int{1, 2, 3} // len , cap = 3
	// [0,2)
	s1 := s[0:2] // [1,2]    pointing to the same memory address
	fmt.Println("s1 = ", s1)
	s1[0] = 100
	fmt.Println(s)  // 100 , 2 ,3
	fmt.Println(s1) // 100, 2

	fmt.Println("-----")

	nn := []int{1, 2, 3, 4, 5}
	fmt.Println("nn[:3] = ", nn[:3]) // 0:3  [0,3)  1,2,3
	fmt.Println("nn[3:] = ", nn[3:]) // 3:len(nn) [3,5)  4,5

	fmt.Println("====copy====")
	// copy
	s2 := make([]int, 3) // 0,0,0
	copy(s2, s)          // copy the content without pointing to the same address
	fmt.Println(s2)
	s2[0] = 123
	fmt.Println(s2) // 123 , 2 ,3
	fmt.Println(s)  // 100, 2, 3

}

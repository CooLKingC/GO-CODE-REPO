package main

import "fmt"

func main() {
	c := fun1("123", 1)
	fmt.Println("c = ", c)

	i1, i2 := fun2()
	fmt.Println("i1 = ", i1)
	fmt.Println("i2 = ", i2)

	r1, r2 := fun3()
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r3, r4 := fun4()
	fmt.Println("r3 = ", r3)
	fmt.Println("r4 = ", r4)

	func() {
		fmt.Println("this is an anonymous function ")
	}()

}

func fun1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 1
	return c
}

// multiple anonymous return values
func fun2() (int, int) {
	return 1, 2
}

// multiple return values
func fun3() (r1 int, r2 int) { // r1 r2 default to 0  ; Scope is within the function
	r1 = 123
	r2 = 321
	return r1, r2
}

// if the return types are all the same
func fun4() (r1, r2 int) {
	r1 = 1234
	r2 = 4321
	return r1, r2
}

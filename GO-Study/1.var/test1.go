package main

/*
  Four ways to declare variables
*/

import (
	"fmt"
)

// global variable
var Ga int
var Gb int = 100
var Gc = 200

func main() {
	// method 1 : declare a variable , default to 0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// method 2 : declare a variable and initialize a value
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// method 3 :  Automatically match data types through values
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// method 4 : Omit 'var' and automatically metch data types   (recommendation)
	//  Cannot be used for global variable
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	e := "abcd"
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := 3.14
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	// Print global variables
	fmt.Println("Ga = ", Ga, "Gb = ", Gb, "Gc = ", Gc)

	// define multiple variable
	var Ma, Mb int = 300, 400
	fmt.Println("Ma = ", Ma, "Mb = ", Mb)
	var Mc, Md = "king", 3.14
	fmt.Println("Mc = ", Mc, "Md = ", Md)

	var (
		lo = 520
		ve = true
	)
	fmt.Println("lo = ", lo, "ve = ", ve)

}

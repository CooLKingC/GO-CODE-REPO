package main

import (
	"fmt"
)

func main() {

	// Implicit type definition
	const an = 100
	fmt.Println("an = ", an)

	// Explicit type definition
	const bn string = "bn"
	fmt.Println("bn = ", bn)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Println("area = ", area)

	// Constants cna also be used for enumeration
	const (
		Unknow = 0 // unkow
		Female = 1 // man
		Male   = 2 // woman
	)

	// keywords iota
	const (
		//  iota += 1
		SHANGHAI = iota // the first line defaults to 0
		BEIJING         // iota += 1   , so here , iota = 1
		SHENZHEN        // iota = 2
	)

	fmt.Println("SHANGHAI = ", SHANGHAI) // 0
	fmt.Println("BEIJIGN = ", BEIJING)   // 1
	fmt.Println("SHENZHEN = ", SHENZHEN) //2

	// modify add value
	const (
		one = iota * 10
		two
		three
	)
	fmt.Println("one = ", one)     // 0
	fmt.Println("two = ", two)     // 10
	fmt.Println("three = ", three) // 20

	// practice
	const (
		a, b = iota + 1, iota + 2 // iota = 0 , a = 1 , b = 2
		c, d                      // iota = 1 , c = iota + 1 = 2 , d = iota + 2 =  3
		e, f                      // iota = 2 , e = iota + 1 = 3 , f = iota + 2 = 4
		g, h = iota * 2, iota * 3 // iota = 3 , g = iota * 2 = 6 , h = iota * 3 = 9
		i, k                      // iota = 4 , i = iota * 2 =8 , k = iota * 3 = 12
	)
	fmt.Println("a = ", a)
	fmt.Println(b, c, d, e, f, g, h, i, k)
}

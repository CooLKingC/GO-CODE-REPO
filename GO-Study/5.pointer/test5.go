package main

import "fmt"

func main() {
	a := 1
	changeValue(&a)        // pass in the memory address     p = &a
	fmt.Println("a = ", a) // 10

	// practice
	b := 10
	c := 20
	swap(&b, &c)
	fmt.Println("b =", b)
	fmt.Println("c = ", c)

	var i *int
	i = &a
	fmt.Println(&a)
	fmt.Println(i)
	var ii **int // second rank pointer
	ii = &i
	fmt.Println(&i)
	fmt.Println(ii)

}

func changeValue(p *int) { // 0x00000000
	*p = 10 // assian a changeValue
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

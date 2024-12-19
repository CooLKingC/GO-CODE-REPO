package main

import "fmt"

func main() {
	// buffer channel
	c := make(chan int, 3)

	fmt.Println("len(c) :", len(c), " cap(c) : ", cap(c))

	go func() {
		defer fmt.Println("goroutine end")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("goroutine out : ", i, "len(c) : ", len(c), "cap(c):", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main end")
}

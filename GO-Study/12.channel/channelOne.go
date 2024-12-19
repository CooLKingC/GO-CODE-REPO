package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine run...")

		c <- 666
	}()

	num := <-c

	fmt.Println("num = ", num)
	defer fmt.Println("main goroutine end")

	fmt.Println("忧伤的音符飘上了天空")
}

package main

import "fmt"

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fun(c, quit)
}

func fun(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = x+y, x
		case <-quit:
			fmt.Println("break")
			return
		}
	}
}

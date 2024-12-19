package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
		func(a int, b int) bool {
			fmt.Println("a = ", a, "b = ", b)
			return true
		}(10, 20)
	*/
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

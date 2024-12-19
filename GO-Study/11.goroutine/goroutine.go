package main

import (
	"fmt"
	"time"
)

// child goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : %d\n", i)
		time.Sleep(1 * time.Second)

	}
}

// main goroutine
func main() {
	// new go thread
	go newTask()

	i := 0

	for i < 6 {
		i++
		fmt.Printf("main goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("main goroutine exit")
}

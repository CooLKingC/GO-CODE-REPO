package main

import "fmt"

func main() {
	var a string
	// pair <statictype:string, value :"Golang"
	a = "Golang"

	// pair <type:string , value : "Golang"
	var allType interface{}
	allType = a
	value, ok := allType.(string)
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println(value)
	}
}

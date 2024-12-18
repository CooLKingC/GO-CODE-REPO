package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("type : ", reflect.ValueOf(arg))
}

func main() {
	i := 3.14
	reflectNum(i)
}

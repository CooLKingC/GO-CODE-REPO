package main

import "fmt"

func main() {
	// map , key : int , value : string
	// 1.
	var m map[string]string
	if m == nil {
		fmt.Println("nil")
	}
	// allocate space
	m = make(map[string]string, 10)

	m["one"] = "Go"
	m["two"] = "Java"
	m["three"] = "C++"
	m["zxcv"] = "C"
	m["adf"] = "Python"
	m["wer"] = "Lua"
	m["oijj"] = "JavaScript"
	m["zcv"] = "HTML"
	m["oiawjeofi"] = "CSS"

	fmt.Println(m) // sort by key

	// 2.
	m2 := make(map[int]string)
	m2[1] = "Go"
	m2[2] = "java"
	m2[3] = "Python"
	fmt.Println(m2)

	// 3.
	m3 := map[int]float64{
		1: 1.01,
		2: 2.01,
		3: 3.01,
	}
	fmt.Println(m3)
}

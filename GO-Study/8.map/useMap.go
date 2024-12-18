package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "Go"
	m[2] = "Java"
	m[3] = "python"
	f(m)

	fmt.Println("===delete===")
	delete(m, 2)
	f(m)

	fmt.Println("===modify===")
	m[3] = "Python"
	changeMap(m)
	f(m)

}

func f(m map[int]string) {
	// pass by reference
	for key, value := range m {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func changeMap(m map[int]string) {
	m[4] = "C/C++"
}

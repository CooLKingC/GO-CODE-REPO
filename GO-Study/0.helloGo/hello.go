package main // Package Name

/*
import "fmt"
import "time"
*/

import ( // import package
	"fmt"
	"time"
)

// main
func main() { // this '{' must be on the same lines as the function name
	fmt.Println("hello Go! ") // can be added ';'  not recommended added
	time.Sleep(1 * time.Second)
}

// go run hello.go
// go build hello.go && ./hello

package main

import (
	"./t1"       // initialization will load init()
	_ "./t2"     // anonymous import , even if not used,there will be no errors
	name1 "./t3" // alias
	. "./t4"     //  call methods directly without specifiying the package name
	// ^  be careful of package name conflicts
)

func main() {
	t1.Api()
	name1.Api()
	Apit4()
}

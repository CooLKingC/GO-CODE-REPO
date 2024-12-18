package main

import "fmt"

// Class names and properties in uppercase represent public visibility
// and so do methods

/*
	type Hero struct {
		name  string
		ad    int
		Level int
	}

	func (this Hero) Show() {
		fmt.Println("Hero = ", this)
	}

	func (this Hero) GetName() string {
		return this.name
	}

	func (this Hero) SetName(name string) {
		// this is a copy
		this.name = name
	}
*/

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("Hero = ", this.Name, this.Ad, this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (hero *Hero) SetName(Name string) {
	hero.Name = Name // 'this' is not a keyword and can be named arbitrarily
}

func main() {
	hero := Hero{Name: "king", Ad: 1, Level: 100}
	hero.Show()
	fmt.Printf("hero.name =  %v\n", hero.GetName())
	hero.SetName("King")
	fmt.Printf("hero.name =  %v\n", hero.GetName())
}

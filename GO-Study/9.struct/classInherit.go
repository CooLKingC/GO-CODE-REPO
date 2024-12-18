package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (human *Human) Eat() {
	fmt.Println("human eat...")
}
func (human *Human) Walk() {
	fmt.Println("human walk...")
}

type SuperMan struct {
	Human // inherit human
	Level int
}

// override a method from the superclass
func (superman *SuperMan) Eat() {
	fmt.Println("superman eat...")
}

func (superman *SuperMan) Fly() {
	fmt.Println("superman fly...")
}

func (superman *SuperMan) Print() {
	fmt.Println("name = ", superman.Name)
	fmt.Println("sex = ", superman.Sex)
	fmt.Println("level = ", superman.Level)
}

func main() {
	// h := Human{Name: "king", Sex: "man"}
	h := Human{"king", "man"}
	h.Eat()
	h.Walk()

	// s := SuperMan{Human{"King", "man"}, 100}
	var s SuperMan
	s.Name = "King"
	s.Sex = "man"
	s.Level = 100

	s.Fly()
	s.Eat()
	s.Walk()

	s.Print()

}

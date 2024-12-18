package main

import "fmt"

// is essentially a pointer
type AniamlIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

/*
Only by overriding all the methods of the interface
is it considered to have implemented the current interface
*/
type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("cat is slepp...")
}
func (cat *Cat) GetColor() string {
	return cat.color
}
func (cat *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("dog is slepp...")
}
func (dog *Dog) GetColor() string {
	return dog.color
}
func (dog *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AniamlIF) {
	animal.Sleep()
	fmt.Println("animal color : ", animal.GetColor())
	fmt.Println("animal type : ", animal.GetType())
}

func main() {
	var animal AniamlIF
	animal = &Cat{"Green"}
	animal.Sleep() // cat sleep

	animal = &Dog{"Yellow"}
	animal.Sleep() // dog sleep

	showAnimal(&Cat{"cat"})
	showAnimal(animal)
}

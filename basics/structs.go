package main

import "fmt"

// Struct declaration
type Person struct {
	name   string
	age    int
	gender string
}

// Struct method
func (p Person) introduce() {
	fmt.Printf("Hello, my name is %s. I'm %d years old and I am %s.\n", p.name, p.age, p.gender)
}

func main() {
	// Creating an instance of the Person struct
	person := Person{name: "Aj.codes", age: 24, gender: "male"}

	// Accessing struct fields
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("Gender:", person.gender)

	// Modifying struct fields
	person.name = "dev-Arjun"
	person.age = 26
	person.gender = "male"

	// Calling struct method
	person.introduce()
}

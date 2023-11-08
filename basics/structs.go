package main

import (
	"unsafe"
	"fmt"
)

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
type EmptyStruct struct{}

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




	// empty structs
	/*main advantage of an empty struct is its minimal memory footprint, making it useful in situations where you want to signal the absence of data without using any additional memory */
	var es EmptyStruct

	// Printing the size of an empty struct
	fmt.Println("Size of an empty struct:", unsafe.Sizeof(es))

	/* 
	in a concurrent program, you could use an empty struct to communicate that a task has been completed
	*/
	
	// Create a channel to communicate task completion
	done := make(chan struct{})

	// Start the worker in a separate goroutine
	go worker(done)

	// Block until receiving a signal that the task is complete
	<-done
	fmt.Println("Task is complete!")
}

// Worker function that performs some task
func worker(done chan struct{}) {
	fmt.Println("Working...")
	// Signal that the task is complete
	done <- struct{}{}
}
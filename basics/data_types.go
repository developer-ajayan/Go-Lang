package main

import "fmt"

func main() {
	// Numeric Types
	var age int = 25
	var temperature float64 = 28.5
	var isAdult bool = true

	fmt.Println("Age:", age)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Is Adult:", isAdult)

	// String Type
	var name string = "John Doe"
	fmt.Println("Name:", name)

	// Character Types
	var ch byte = 'A'
	var unicodeChar rune = '♥'
	fmt.Println("Byte Character:", ch)
	fmt.Println("Unicode Character:", unicodeChar)

	// Collection Types
	var numbers [3]int = [3]int{1, 2, 3}
	var fruits []string = []string{"Apple", "Banana", "Orange"}
	var person = struct {
		name string
		age  int
	}{"Alice", 30}
	var scores = map[string]int{"Math": 90, "English": 85}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)
	fmt.Println("Person:", person)
	fmt.Println("Scores:", scores)

	// Pointer Types
	var numPtr *int
	num := 10
	numPtr = &num
	fmt.Println("Number:", num)
	fmt.Println("Number Pointer:", numPtr)
	fmt.Println("Dereferenced Number Pointer:", *numPtr)

	// Function Types
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Println("Addition Result:", result)

	// Type Aliases
	type Celsius float64
	var temperatureC Celsius = 23.8
	fmt.Println("Temperature in Celsius:", temperatureC)
}

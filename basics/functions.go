package main

import "fmt"

// Function with input parameter and return value
func add(x, y int) int {
	return x + y
}

// Function with variadic parameter
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Factorial of number
func factorial(num int) int {
	if num > 0 {
		return num * factorial(num-1)
	} else {
		return 1
	}

}

// Function returning multiple values
func divide(x, y int) (int, bool) {
	if y == 0 {
		return 0, false
	}
	return x / y, true
}

// Function returning a function
func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// Calling function with input parameter and return value
	result := add(10, 5)
	fmt.Println("Addition:", result)

	// Calling function with variadic parameter
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sumResult)

	// factorial of a number
	facResult := factorial(5)
	fmt.Println("Sum:", facResult)

	// Calling function with multiple return values
	quotient, success := divide(10, 2)
	fmt.Println("Division Result:", quotient, "Success:", success)

	// Calling function returning a function
	multiplyByTwo := getMultiplier(2)
	multiplyByThree := getMultiplier(3)

	fmt.Println("Multiplication by 2:", multiplyByTwo(5))
	fmt.Println("Multiplication by 3:", multiplyByThree(5))

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}()
}

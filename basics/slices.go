package main

import (
	"fmt"
)

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println("slice0 ....", slice)
	slice1 := []int{1, 2, 3, 10, 12, 13}
	fmt.Println("slice1 ....", slice1)
	slice2 := [2]int{1, 3} // exceeding the size will cause panic
	fmt.Println("slice2 ....", slice2)
	// number of element in the slice
	fmt.Println("length of slice...", len(slice2))
	//capacity is number of elements in the underlying array
	fmt.Println("cappacity of slice...", cap(slice2))

	// slice modification
	slice3 := []int{1, 2, 3, 10, 12, 13}
	sliced := slice3[2:5] // it doesn't copy slice data It creates a new slice value that points to the original array.
	sliced[0] = 111
	fmt.Println("sliced ....", slice3)

}

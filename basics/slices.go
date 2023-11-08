package main

import (
	"fmt"
	"slices"
)

//https://rezakhademi.medium.com/slices-in-golang-common-mistakes-and-best-practices-76c30857d4e4
// slices are a dynamically sized flexible way of view array elements
// slices are pass by reference by default
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


	// copy slices
	s1 := []int{22, 44, 66, 88}

	s2 := append([]int{}, s1[:2]...)
	s2 = append(s2, s1[3:]...)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// slices equal
	original := []int{22, 44, 66, 88}
	copyslice := make([]int, len(original))
	copy(copyslice, original)
	original[0] = 44
	original[1] = 22
	result:=slices.Equal(copyslice, original)
	fmt.Println(original, copyslice, result)


}

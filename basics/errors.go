package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// error is like a default datatype no need to import
// errors is a package to create new error
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return f * f, nil
	}

}
func main() {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success", f)
	}

	// As function
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("As function populated error")
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	// Is Function
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Is function populated error")
			fmt.Println("Failed at path:non-existing not found ")
		} else {
			fmt.Println(err)
		}
	}

	// Join function
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err = errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}

}

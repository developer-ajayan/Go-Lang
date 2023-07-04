package main

import (
	"fmt"
)

func main(){
	var a,b int=10,10
	// storing the address of a into p
	var p=&a
	// increse the value of a
	a++
	// increse the value of address stored in p
	(*p)++
	// the answer will be 22
	fmt.Println("value of a+b on last",a+b)
}
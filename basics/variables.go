package main
import(
	"fmt"
)
func main(){
	// default 
	var numberone int
	numberone=10
	fmt.Println("number one",numberone)
	
	// declaring and initializing a variable in single line
	var numbertwo int = 30
	fmt.Println("number second",numbertwo)

	/*
	 You can declare multiple variables at once.
	*/
	var stringone, stringtwo string = "test1","test2"
    fmt.Println("multiple variables at once: ",stringone, stringtwo)

	/* 
		The := syntax is shorthand for 
	    declaring and initializing a variable 
		This syntax is only available inside functions.
	*/
	numberthree:=20
	fmt.Println("number three",numberthree)	

	// constants
	const pi = 3.1415

}
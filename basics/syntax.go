// package declaration
package main  

// importing packages
import "fmt"
//performance : https://blog.stackademic.com/best-practices-in-go-golang-writing-clean-efficient-and-maintainable-code-dccf61542b57
// function declaration
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
}

// variable declaration
var variableName type
var variableName = value

// constants declaration
const constantName type = value

// control statements
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}

// for loop
for initialization; condition; postIteration {
    // code to execute in each iteration
}

// infinity loop (while loop)
for condition {
    // code to execute while the condition is true
}

//switch statement
switch expression {
case value1:
    // code to execute if expression matches value1
case value2:
    // code to execute if expression matches value2
default:
    // code to execute if expression matches none of the values
}

// slices
var sliceName []type

// structs
type StructName struct {
    field1 type1
    field2 type2
    // ...
}

// pointers
var variableName *type  // declaration
variableName = &value   // assigning the address of value to the pointer
*variableName           // dereferencing the pointer to get the value

// Error handling
if err != nil {
    // handle the error
}

// function call
functionName(argument1, argument2)

/*  Keypoints:
nil represent the zero/uninitialized value for certain types(pointers, interfaces, maps, slices, channels, functions)
for example the type of a nil pointer  has *int or *string. a nil slice has type like []int or []string
nil is a predeclared identifier, and used to indicate the absence of a valid value.


*/

package main

import "fmt"

func main() {
	performOperation()
}

func performOperation() {
	defer func() { // defer is used to execute a function just before the current function returns
		if r := recover(); r != nil { // recover is used to handle a panic within performOperation or its nested functions
			// if panic occurs defer is invoked, and the recovered value is assigned to r
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Performing operation...")

	// Simulating an error condition
	panic("Something went wrong!") // this cause the program to panic and invoke deferred functions
}

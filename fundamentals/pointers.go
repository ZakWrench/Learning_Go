/*
	Key points

default or zero-value pointers are nil. uninitialized pointer will always have a nil value
*/
package main

import "fmt"

func main() {
	//storing hex values in vars
	c := 0xFFF
	fmt.Printf("type of the variable c is %T\n", c)
	fmt.Printf("in decimal: c = %d \nin binary: c = %b \nin hex: c = %#x\n", c, c, c)

	//declaring a pointer
	var pointername *int // contain the moemory address of an integer variable
	fmt.Printf("type of the variable pointername is %T\n", pointername)

	// Pointer initialization
	x := 44
	var st *int // st is a pointer to an int variable
	st = &x     // st now points to the memory address of x
	fmt.Printf("x value is %d \nx memory address is %v \nst value is %v \nst memory address is %v\n", x, &x, st, &st)

	// Type inference for pointers
	typeless := 8
	var typelessPointer = &typeless
	fmt.Printf("type of the variable typelessPointer is %T\n", typelessPointer)
}

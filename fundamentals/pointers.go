/*
	Key points

default or zero-value pointers are nil. uninitialized pointer will always have a nil value
When Comparing pointers, we check whether they point to the same memory address, not the values they point to
*/
package main

import "fmt"


type employee struct {
	name string
	empid int
	salary int
}

func ptf(b *int){
	*b = 8
}

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

	// Dereferencing a pointer
	// x value is 44 and it's memory address is 0xc0000a4020
	// st value is 0xc0000a4020 and it's memory address is 0xc00009c020

	var a = &x 
	fmt.Printf("defereferencing `a` pointer: %v\n", *a) 
	*a = 100
	fmt.Printf("defereferencing `a` pointer: %v\n", *a)
	fmt.Println("Value stored in x (*a) after Changing =", x)

	// struct instantiation using the new keyword

	emp1 := new(employee)
	emp1.name = "ABC"
	emp1.empid = 123
	emp1.salary = 10000
	fmt.Println(emp1)

	// Pointers to a function
	y := 200
	fmt.Printf("Value of y before function call is : %d\n",y)
	pb := &y
	ptf(pb)
	fmt.Printf("Value of y after function call is : %d\n",y)

	// Passing Address of the variable to a function call
	z := 300
	fmt.Printf("Value of z before function call is : %d\n",z)
	ptf(&z)
	fmt.Printf("Value of z after function call is : %d\n",z)

	// Pointer to a Struct

	emp := employee{ "XYZ", 456, 20000}
	pts := &emp
	fmt.Println(pts)
	fmt.Printf("Name of the employee is %s\n", pts.name)
	fmt.Println((*pts).name)

	// Declaring Pointer to a Pointer

	var V int = 999
	var pt1 *int = &V
	var pt2 **int = &pt1
	fmt.Println("Value of Variable V is =", V)
    fmt.Println("The Address of variable V is =", &V)
    fmt.Println("Value of pt1 is =", pt1)
    fmt.Println("The Address of pt1 is =", &pt1)
    fmt.Println("Value of pt2 is =", pt2)
    fmt.Println("The Value at the address of pt2 is or *pt2 =", *pt2)
    fmt.Println("*(The Value at the address of pt2 is) or **pt2 =", **pt2)

	// Pointers Comparision

	comp1 := 88
	comp2 := 88

	comp1_ptr := &comp1
	comp2_ptr := &comp2
	
	// ==
	if comp1_ptr == comp2_ptr {
		fmt.Println("comp1_ptr and comp2_ptr are pointing to the same memory address")
	} else {
		fmt.Println("comp1_ptr and comp2_ptr are not pointing to the same memory address")
	}
	// !=
	if comp1_ptr != comp2_ptr {
		fmt.Println("comp1_ptr and comp2_ptr are not pointing to the same memory address")
	} else {
		fmt.Println("comp1_ptr and comp2_ptr are pointing to the same memory address")
	}
	
}

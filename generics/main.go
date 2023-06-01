package main

import "fmt"

//SumInts add together the values of m.

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

//SumFloats adds together the values of m.

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// in the above code, we declared 2 functions to add together the values of a map and return the sum
// SumFloats takes a map of string to float64 vlaues
// SumInts takes a map of string to int64 values.

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64 as types for map values
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

/*
  - In the above code, we declared SumIntsOrFloats function with two type parameters(inside the square brackets),
    K and V, and one argument that uses the type parameters, m of type map[K]V. The function returns a value of type V.

  - Specify for the K type parameter the type constrain `comparable`, which is intended specifically for cases like these, the `comparable` constraint is predeclared in Go.
    it allows any types whose values may be used as an operand of the comparision operators == and !=. Go requires that map keys be comparable. So declaring K as comparable is necessary so you use K as the key inthe map variable.
    It also ensures that calling code uses an allowable type for map keys

  - Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using `|` speciffies a union of the two types, meaning that this constraint allows
    either type. Either type will be permitted by the compilter as an argument in the calling code.

  - Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters. Note that we know map[K]V is a valid map type because K is a comparable type.
    If we hadn't declared K comparable, the compiler would reject the reference to map[K]V
*/
func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generics Sums: %v and %v\n",
		SumInts(ints), SumFloats(floats))

	// In the above code, we initialize a map of float64 values and a map of int64 values, each with two entries.
	// Call the two functions we declared earlier to find the sum of each map's values

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// in the above code we call the generic function we just declared passing each of the maps we created.
	// Specified type arguments, the type names in square brackets, to be clear about the types thatshould replace type parameters in the function you're calling.
	// We can often omit the type arguments in the function call. Go can often infer them from the code

	/*
		When running the code, in each call the compiler replaced the type parameters with the concrete types specified in that call.
		In calling the generic function, we specified type arguments that told the compiler what types to use in place of the function's type parameters.
	*/

	// REMOVING THE TYPE ARGUMENTS WHEN CALLING THE GENERIC FUNCTION
}

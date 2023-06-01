/*
	Key points
	Structs:
		A struct is a composite data type that allows you to group together zero or more values with different types under a single entity.
		It provides a way to create a user-defined types with their own set of fields.
		Structs are often used to represent objects or records.
		Structs can also be embedded within each oter, allowing you to create more complex data structures.
		You can define methods associated with a struct to provide behavior specific to that struct.
*/

/*
Moar on structs:
	- Struct fields:
	Each value within a struct is called a `field`, to access the fields of a struct, we can use the dot notation. for ex: struct name is `Person` with a field `Name`, to access it we use `person.Name` where `person` is an instance of `Person`

	- Pointers to a Struct:
	When we have a pointer to a struct, we can modify the struct's fields directly, even if we pass it to a function. Which allows us to avoid making unnecessary copies of the struct.

	- Anonymous structures:
	Anonymous structs are useful when we need to create temporary data structure or when the structure is only used in one place

	*/


package main

import "fmt"

type Person struct {
	Name string // exported field
	Age int
}
type Employee struct {
	Person // anonymous field
	EmployeeID int 
	Salary float64
}

type Anon_Fields struct{
	string
	int
}



func main(){

	me := Person{
		Name : "Zak",
		Age : 29,
	}
	fmt.Println(me.Name)
	me_composition := Employee{
		Person : Person{
			Name : "Zak",
			Age : 29,
		},
		EmployeeID : 123,
		Salary : 10000,
	}
	fmt.Println(me_composition.Name, me_composition.Age, me_composition.EmployeeID, me_composition.Salary)

	me_pointers := &Person{
		Name : "Zak",
		Age : 29,
	}
	me_pointers.Age = 30
	fmt.Println(me_pointers.Name, me_pointers.Age)

	// Anonymous structs
	anon := struct {
		Name string
		id int
	}{
		Name : "Achraf",
		id : 4,
	}
	fmt.Println(anon.Name, anon.id)

	// Anonymous fields
	anonField := Anon_Fields{
		"Moad",
		28,
	}
	fmt.Println(anonField.string, anonField.int)
}
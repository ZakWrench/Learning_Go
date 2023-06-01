package main

import "fmt"

// Creating the structure
type Students struct {
	Name     string
	BranchNo string
	Year     int
}

// Creating the nested structure
type Teachers struct {
	Name    string
	Subject string
	Experience int
	Details Students
}

func main() {
	// Initializing fields of the structure
	results := Teachers{
		Name:    "Sunita",
		Subject: "PHP",
		Experience: 2,
		Details: Students{
			Name:     "Rahil",
			BranchNo: "CDE",
			Year:     4,
		},
	}

	// Display the values
	fmt.Println("Details of the Teachers")
	fmt.Println("Teacher's name: ", results.Name)
	fmt.Println("Subject: ", results.Subject)
	fmt.Println("Experience: ", results.Experience)

	fmt.Println("\nDetails of Students")
	fmt.Println("Student's name: ", results.Details.Name)
	fmt.Println("Student's branch name: ", results.Details.BranchNo)
	fmt.Println("Year: ", results.Details.Year)
}

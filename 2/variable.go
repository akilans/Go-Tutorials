// Create variable to store values
// Create function to retrieve values

package main

import "fmt"

func main() {
	fmt.Println("My name is -> ", getName())
	fmt.Println("I am working in -> ", getCompany())
	fmt.Println("My role is -> ", getRole())
}

func getName() string {

	// var name string
	// name = "Akilan"

	name := "Akilan" // Auto assign to string
	return name
}

func getCompany() string {
	company := "Infosys"
	return company
}

func getRole() string {
	role := "Devops Engineer"
	return role
}

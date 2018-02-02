package main

import (
	"fmt"
	"strconv"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	variableDeclaration() // Basic variable declaration
	structMethod1()       // Struct method 1
	structMethod2()       // Struct method 2
	structMethod3()       // Struct method 3
}

func structMethod1() {
	akilan := person{"Akilan", "Subramanian", 29}
	fmt.Println("My Name is " + akilan.firstName + " - My father name is " + akilan.lastName + " - My age is " + strconv.Itoa(akilan.age))
}

func structMethod2() {
	akilan := person{firstName: "Akilan", lastName: "Subramanian", age: 29}
	fmt.Println("My Name is " + akilan.firstName + " - My father name is " + akilan.lastName + " - My age is " + strconv.Itoa(akilan.age))
}

func structMethod3() {
	var akilan person
	fmt.Println("Empty person details", akilan)
	fmt.Printf("%+v", akilan)

	fmt.Println("My Name is " + akilan.firstName + " - My father name is " + akilan.lastName + " - My age is " + strconv.Itoa(akilan.age))

	// Set value
	akilan.firstName = "Akilan"
	akilan.lastName = "Subramanian"
	akilan.age = 29

	fmt.Println("My Name is " + akilan.firstName + " - My father name is " + akilan.lastName + " - My age is " + strconv.Itoa(akilan.age))
}

func variableDeclaration() {
	var name string // Mention the type of variable explicitly
	name = "Akilan" // Assign the value
	fmt.Println("My name is ", name)

	company := "Infosys" // Auto Assign to string
	fmt.Println("I am working in ", company)

	age := 29 // Auto Assign to Integer
	fmt.Println("My age is ", age)
}

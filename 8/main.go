// pointer explained
// go run main.go struct.go
// pass by value vs pass by reference clearly explained
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
	variableDeclaration()                            // Basic variable declaration
	structMethod1()                                  // Struct method 1
	structMethod2()                                  // Struct method 2
	structMethod3()                                  // Struct method 3
	emp := newStructType()                           // Call struct inside struct example in struct.go file
	emp.updateEmailByValue("karuna@infosys.com")     // update employee email by passing value
	fmt.Println(emp)                                 // email not update if we pass by value
	updateEmailByPointer(&emp, "karuna@infosys.com") // pass by reference
	fmt.Println(emp)
	newEmp := &emp                                  // Get the memory address from value
	newEmp.updateEmailByPointer("karuna@yahoo.com") // update by pointer
	fmt.Println(emp)

	emp.updateEmailByPointer("struct@gmail.com") // Receiver Function knows how to deal
	fmt.Println(emp)

	// Integeer also same
	age := 29
	changeValue(age)
	fmt.Println(age)

	// slice ??? By default it is pass by reference not by value
	// Slice use two address to store 1 slice 1) data 2) header pointing to that data
	// When update happens it copies the header pointing to that data
	// so it updates the correct data !!!!
	// Maps, Channels, pointers, functions behaves same way
	friends := []string{"Akilan", "Alex", "Jegan"}
	changeSliceValue(friends)
	fmt.Println(friends)

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

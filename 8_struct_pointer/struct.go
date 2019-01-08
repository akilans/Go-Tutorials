package main

import "fmt"

type contactInfo struct {
	location string
	email    string
}

type employee struct {
	name    string
	role    string
	contact contactInfo
}

func newStructType() employee {
	fmt.Println("Hello Akilan")
	karuna := employee{
		name: "Karuna",
		role: "Technology Analayst",
		contact: contactInfo{
			location: "Hyderabad",
			email:    "karuna@gmail.com",
		},
	}
	return karuna
}

func (e employee) updateEmailByValue(email string) {
	e.contact.email = email
}

func updateEmailByPointer(e *employee, email string) {
	e.contact.email = email
}

func (e *employee) updateEmailByPointer(email string) {
	(*e).contact.email = email // * gives the value of pointer pointing
}

func changeValue(age int) {
	age = 30
}

func changeSliceValue(friends []string) {
	friends[0] = "Aki"
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type welcome struct {
	Name     string
	Language string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Akilan, Welcome to Go Programming!!!</h1>")
}

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go Web Server is Running!!!<h1>")
}
*/

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	welcomeData := welcome{Name: "Akilan", Language: "Go"}
	t, _ := template.ParseFiles("hello.html")
	t.Execute(w, welcomeData)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/welcome", htmlHandler)
	http.ListenAndServe(":8000", nil)
}

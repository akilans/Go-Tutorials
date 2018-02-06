// Create a  Web server
// Place HTML tags

package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Web Server is Running!!!")
}

func akilanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Akilan, Welcome to Go Programming!!!")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Print Multiple line</h1>
		<p>Here we Go!!!</p>
		<p>Go Rocks!!!</p>
		`)
	fmt.Fprintf(w, "Hi %s, Welcome to %s Programming!!!", "Akilan", "Go")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/akilan", akilanHandler)
	http.HandleFunc("/html", htmlHandler)
	http.ListenAndServe(":8000", nil)
}

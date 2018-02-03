package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error is", err)
		os.Exit(1)
	}
	//fmt.Println(resp) // Body content will not print

	// byteSlice := make([]byte, 99999) // Creates bytes of slices with 99999 empty
	// so that it get updated while readind body from response
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	io.Copy(os.Stdout, resp.Body) // Simple way to print on screen

}

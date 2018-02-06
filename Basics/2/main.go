// Get data from URL
// It is not working in Infy network
// XML parsing needs to check "encoding/xml" package

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com") // _ ignore error value

	if err != nil {
		fmt.Print("Error - ", err)
		os.Exit(1)
	}

	bodyInBytes, _ := ioutil.ReadAll(resp.Body) // _ ignore error value

	content := string(bodyInBytes)

	fmt.Println(content)
}

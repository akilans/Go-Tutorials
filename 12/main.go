package main

import (
	"fmt"
	"net/http"
)

func main() {

	websites := []string{
		"http://facebook.com",
		"http://google.com",
		"http://tamilrockessrss.com",
		"http://twitter.com",
	}

	for _, link := range websites {
		testLink(link)
	}

}

func testLink(l string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " is down!!!")
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println(l, " is up!!!")
}

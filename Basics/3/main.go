package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type siteMap struct {
	Locations []string `xml:"sitemap>loc"`
}

type news struct {
	Title   []string `xml:"url>news>title"`
	Keyword []string `xml:"url>news>keyword"`
}

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to index page<h1>")
	fmt.Fprintf(w, `
		<h4>Multi line print Example</h4>
		<p>First Line</p>
		<p>Second Line</p>
		`)
}
*/

func main() {
	var urls siteMap
	var news news
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bodyInBytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyInBytes))
	xml.Unmarshal(bodyInBytes, &urls)
	//fmt.Println(urls)

	for _, url := range urls.Locations {
		//fmt.Println(i+1, " ", url)

		resp, _ := http.Get(url)
		bodyInBytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bodyInBytes, &news)

		for i, title := range news.Title {
			fmt.Println(i+1, " - ", title)
		}

	}
	//http.HandleFunc("/", indexHandler)
	//http.ListenAndServe(":8000", nil)
}

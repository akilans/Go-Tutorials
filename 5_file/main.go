// Convert slice of strings to single string
// Write strings of data in a file & store it in hard disk

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	skills := []string{"PHP", "PYTHON", "RUBY", "GO"}

	skillString := toString(skills) // Convert slice of strings into single string

	saveTofile("skills.txt", skillString) // Write file function call

	skills = readFile("skills.txt") // Read file function call
	printSliceofStrings(skills)
}

func toString(s []string) string {
	fmt.Println("Converting slice of strings into single string....")
	return strings.Join(s, ",")
}

// Print slice of strings using for loop

func printSliceofStrings(ss []string) {
	fmt.Println("Printing Slice of strings....")
	for i, s := range ss {
		fmt.Println(i+1, s)
	}
}

// write file function

func saveTofile(filename string, skills string) error {
	fmt.Println("Saving file....")
	return ioutil.WriteFile(filename, []byte(skills), 0666) // Write to disk by bytes of data
}

// Read file function

func readFile(filename string) []string {
	fmt.Println("Reading file....")
	byteSlice, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error - ", error)
		os.Exit(1)
	}
	fmt.Println(string(byteSlice))
	s := strings.Split(string(byteSlice), ",")
	return []string(s)
}

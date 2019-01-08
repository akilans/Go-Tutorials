// Map to store data in key & value pair
// All keys must be same type
// All values must be same type
// values extracted by keys in double quotes
// For loop not possible in struct

package main

import (
	"fmt"
)

func main() {
	friends := make(map[string]string)

	friends["Alex"] = "Chennai"
	friends["Jegan"] = "kovai"

	fmt.Println(friends)
	fmt.Println(friends["Alex"])

	delete(friends, "Alex") // Delete from map
	fmt.Println(friends)

	employees := map[string]string{
		"name":     "Akilan",
		"role":     "Devops Engineer",
		"location": "Bangalore",
	}
	fmt.Println(employees["name"])
	fmt.Println(employees)

	employees["name"] = "Aki"
	fmt.Println(employees)

	for key, value := range employees {
		fmt.Println(key, " is ", value)
	}
}

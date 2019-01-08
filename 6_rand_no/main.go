// Generate Random number
// Source will not change after one time, So need to generate source randomly everytime

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	skills := []string{"PHP", "PYTHON", "RUBY"}
	fmt.Println("Before shuffling - ", skills)
	shuffle(skills)
}

func shuffle(s []string) {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range s {
		newPosition := r.Intn(len(s) - 1)
		fmt.Println("Old position", i)
		fmt.Println("New Position", newPosition)
		s[i], s[newPosition] = s[newPosition], s[i]
	}
	fmt.Println("After shuffling...")
	fmt.Println(s)
}

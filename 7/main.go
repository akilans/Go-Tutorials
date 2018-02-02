package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSkills())
}

func getSkills() []string {
	skills := []string{"PHP", "PYTHON", "RUBY"}
	return skills
}

//Create slice to Store String as an array
//Print Slice data using for loop
package main

import "fmt"

func main() {

	skills := getAllSkills()
	fmt.Println(skills)

	for i, skill := range skills { // For loop example
		fmt.Println(i+1, skill) // i start from 0 [ index ]
	}
}

func getAllSkills() []string {
	languageSkills := []string{"PHP", "PYTHON", "RUBY"}
	languageSkills = append(languageSkills, "Go") // Append variable [ := used for first time declaration]
	return languageSkills
}

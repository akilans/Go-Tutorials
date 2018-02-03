// go run main.go skill.go
package main

import "fmt"

func main() {
	skills := skill{"PHP", "PYTHON", "RUBY"}
	fmt.Println(skills)

	skills = skills.addSkill("GO") // Call receiver Function
	skills.print()                 // Call receiver Function

	firstSkill, lastSkill := skills.firstLastSkill() // Call receiver Function
	firstSkill.print()
	lastSkill.print()

	fmt.Println(skills.toString())

}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	skills := getSkills()
	fmt.Println(skills)
	writeSkillFile("skills.txt", skills)
}

func getSkills() []string {
	skills := []string{"PHP", "PYTHON", "RUBY"}
	return skills
}

func writeSkillFile(filename string, skills []string) {
	fmt.Println("Writting in a file")
	skillsString := strings.Join(skills, ",")
	error := ioutil.WriteFile(filename, []byte(skillsString), 0666)

	if error != nil {
		fmt.Println("Error - ", error)
		os.Exit(1)
	}
}

func readSkillFile(filename string) []string {

	byteSlice, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error - ", error)
		os.Exit(1)
	}

	return []string(strings.Split(string(byteSlice), ","))

}

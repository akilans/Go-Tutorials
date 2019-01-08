package main

import (
	"os"
	"testing"
)

// Tesing number of skills

func TestGetSkills(t *testing.T) {

	skills := getSkills()

	if len(skills) != 3 {
		t.Errorf("Expected number of skills 3 but you have %v", len(skills))
	}

}

//// Tesing number of skills after writting

func TestWriteSkillFile(t *testing.T) {

	filename := "_skilltesting.txt"

	skills := getSkills()

	os.Remove(filename)

	writeSkillFile(filename, skills)

	skills = readSkillFile(filename)

	if len(skills) <= 2 {
		t.Errorf("Expected length is %v but got %v", 2, len(skills))
	}

	os.Remove(filename)

}

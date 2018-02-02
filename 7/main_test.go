package main

import (
	"testing"
)

func TestSkill(t *testing.T) {

	skills := getSkills()

	if len(skills) != 4 {
		t.Errorf("Expected number of skills 4 but you have %v", len(skills))
	}

}

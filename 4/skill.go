package main

import "fmt"

type skill []string

func (s skill) print() { // Only Instance of type skill can call this function [ Receiver ]
	for i, skill := range s {
		fmt.Println(i+1, skill)
	}
}

func (s skill) addSkill(skillName string) skill {
	s = append(s, skillName)
	return s
}

func (s skill) firstLastSkill() (firstSkill skill, lastSkill skill) {
	return s[:1], s[3:]
}

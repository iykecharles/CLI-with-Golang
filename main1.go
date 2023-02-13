package main

import (
	"fmt"
	"time"
)

func main1() {
	fmt.Println("welcome to the command line interface in Golang")
	time.Sleep(5 * time.Second)
	theQuestions()

}



func theQuestions() {
	fmt.Println("The program would help you calculate your CGPA. Enjoy!")
	time.Sleep(5 * time.Second)

	score := 0.00
	courses := []string{"Human factors", "Ecommerce system", "ecommerce design or layleh 1st semester course", "Ethics and Reserch methodology", "OOP", "Database analysis and design", "System Design", "Computer Systems and Networks", "Msc Project"}
	gradesMap := make(map[string]string)

	for _, course := range courses {
		fmt.Println("what grade did you get for the course ...!", course)
		var answer string
		fmt.Scan(&answer)
		gradesMap[course] = answer

		switch answer {
		case "A":
			if course == "Msc Project" {
				score += 5 * 60
			} else {
				score += 5 * 15
			}
		case "B":
			if course == "Msc Project" {
				score += 4.0 * 60
			} else {
				score += 4.0 * 15
			}
		case "C":
			if course == "Msc Project" {
				score += 3.0 * 60
			} else {
				score += 3.0 * 15
			}
		}
	}
	fmt.Println(gradesMap)

	fmt.Println((score / 180))
}



/*
func theQuestions() {
	fmt.Println("The program would help you calculate your CGPA. Enjoy!")
	time.Sleep(5 * time.Second)

	score := 0.00
	courses := []string{"Human factors", "Ecommerce system", "ecommerce design or layleh 1st semester course", "Ethics and Reserch methodology", "OOP", "Database analysis and design", "System Design", "Computer Systems and Networks", "Msc Project"}
	grades := []string{}

	for _, course := range courses {
		fmt.Println("what grade did you get for the course ...!", course)
		var answer string
		fmt.Scan(&answer)
		grades = append(grades, answer)

		switch answer {
		case "A":
			if course == "Msc Project" {
				score += 5 * 60
			} else {
				score += 5 * 15
			}
		case "B":
			if course == "Msc Project" {
				score += 4.5 * 60
			} else {
				score += 4.5 * 15
			}
		}
	}
	fmt.Println(courses,grades)

	fmt.Println((score / 180))
}
*/
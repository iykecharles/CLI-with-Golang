package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hi, you can use this cli to check cgpa by entering your course grades. Enjoy! ")
	time.Sleep(5 * time.Second)
	cgpaCalc()
}

func cgpaCalc() {
	courses := []string{"Human factors", "Ecommerce system", "ecommerce design or layleh 1st semester course", "Ethics and Reserch methodology", "OOP", "Database analysis and design", "System Design", "Computer Systems and Networks", "Msc Project"}
	grademap := make(map[string]string) //map
	gradeslice := []string{}            //slice

	score := 0.00

	var name string
	fmt.Println("what is your name ...!")
	fmt.Scan(&name)

	for _, course := range courses {
		var answer string
		fmt.Println("what grade did you get in the course ...!", course)
		fmt.Scan(&answer)

		answer = strings.ToUpper(answer)
		grademap[course] = answer
		gradeslice = append(gradeslice, answer)

		creditMsc := 60.0
		credit := 15.0
		switch answer {
		case "A":
			if course == "Msc Project" {
				score += 5 * creditMsc
			} else {
				score += 5 * credit
			}
		case "B":
			if course == "Msc Project" {
				score += 4 * creditMsc
			} else {
				score += 4 * credit
			}
		case "C":
			if course == "Msc Project" {
				score += 3 * creditMsc
			} else {
				score += 3 * credit
			}
		case "D":
			if course == "Msc Project" {
				score += 2 * creditMsc
			} else {
				score += 2 * credit
			}
		default:
			if course == "Msc Project" {
				score += 0 * creditMsc
			} else {
				score += 0 * credit
			}
		}

	}
	Totalcgpa := score / 180
	fmt.Printf("Your cummulative cgpa is %v \n", Totalcgpa)
	fmt.Println("")

	fmt.Println("Breakdown of your result is as follows : ", grademap)
	fmt.Println("")
	fmt.Println(gradeslice)
	fmt.Println("")

	if Totalcgpa >= 4.5 {
		fmt.Println("Distinction! Well done!", name)
	} else if 3.5 <= Totalcgpa && Totalcgpa >= 4.444449 {
		fmt.Println("Merit! Wonderful!", name)

	} else if 3.0 <= Totalcgpa && Totalcgpa >= 3.49999 {
		fmt.Println("You cant kill yourself Jare! You tried,", name)

	} else {
		fmt.Println("Nice!", name)

	}

	//email
	fmt.Println("")
	time.Sleep(3 * time.Second)

	/*
		m:= fmt.Sprintf("%v", grademap)


		var myemail string
		fmt.Println("give your email address so that you get this information to your inbox")
		fmt.Scan(&myemail)
		auth:= smtp.PlainAuth("", "iykecharles316@gmail.com", "password", "mail.google.com")
		to:= []string{myemail}
		hostname := "10.28.139.5"
		host := "587"
		msg := m

		err := smtp.SendMail(hostname+":"+host, auth, "iykecharles316@yahoo.com", to, []byte(msg))
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Printf("Email finally sent successfully to you email.")

	*/

}

package main

import (
	"Go-CLI/helper"
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var noTicketsRemaining = remainingTickets == 0
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketCount {

			bookTicket(userTickets, lastName, firstName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names of the the bookings are: %v \n", firstNames)

			if noTicketsRemaining {
				fmt.Println("The conference is sold out, please check back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid, please try again")
			}
			if !isValidTicketCount {
				fmt.Println("Number of tickets is invalid, please try again. ")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the lectures!")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, fullName := range bookings {
		var splitNames = strings.Fields(fullName)
		firstNames = append(firstNames, splitNames[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many ticktets would you like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, lastName string, firstName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

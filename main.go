package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings []string

func main()  {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {		
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("The email address you entered does not contain @ sign")
			}

			if !isValidEmail {
				fmt.Println("The number of tickets you entered is invalid")
			}
		}	
	}	
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)  
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(fName string, lName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, fName string, lName string, email string)  {
	remainingTickets -= userTickets

	bookings = append(bookings, fName + " " + lName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fName, lName, userTickets, email)

	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
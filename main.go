package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const totalTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

var wg sync.WaitGroup

func main() {

	welcomeUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isUserTicketsValid := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isUserTicketsValid {
		remainingTickets = remainingTickets - userTickets

		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)
		wg.Add(1)
		sendTickets(userTickets, firstName, lastName, email)

		fmt.Printf("Thanks %v %v for booking %v tickets. You will recieve a confirmation mail on %v .\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining\n\n", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("Sorry, all our seats are booked. Come back next year.")

		}
	} else {
		if !isValidName {
			fmt.Printf("Your input name is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Your email address does not contain @ sign\n")
		}
		if !isUserTicketsValid {
			fmt.Printf("The number of tickets you entered is invalid\n")
		}
	}
	wg.Wait()
}

func welcomeUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still left\n", totalTickets, remainingTickets)
	fmt.Printf("Here you can book your tickets\n")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First Name:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scanln(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	fmt.Printf("%v tickets sent to %v %v's email address %v\n\n\n", userTickets, firstName, lastName, email)
	wg.Done()
}

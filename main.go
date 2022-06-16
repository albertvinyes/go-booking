package main

import (
	"fmt"
	"sync"
	"time"
)

var conferanceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	userName, userEmail, userTickets := getUserInput()
	isValidEmail, isValidName, isValidTicketNumber := validateUserInput(userName, userEmail, userTickets, remainingTickets)
	if isValidEmail && isValidName && isValidTicketNumber {

		bookTickets(userTickets, userName, userEmail)

		wg.Add(1)
		go sendTicket(userTickets, userName, userEmail)

		firstNames := getFirstnames()
		fmt.Printf("The names of our bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets <= 0
		if noTicketsRemaining {
			fmt.Printf("Our conference is booked out. Come next year")
		}
	} else {
		fmt.Println("Your input is invalid. Try again.")
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to booking %v application\n", conferanceName)
	fmt.Printf("We have a tofal of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var userTickets uint
	var userName string
	var email string

	fmt.Println("Enter your full name:")
	fmt.Scan(&userName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want")
	fmt.Scan(&userTickets)

	return userName, email, userTickets
}

func bookTickets(userTickets uint, userName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Lisft of bookings is %v", bookings)
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v \n", userName, userTickets, email)
	fmt.Printf("%v tickets reamining for %v. \n", remainingTickets, conferanceName)
}

func sendTicket(userTickets uint, userName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

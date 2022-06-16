package main

import "strings"

var MyGlobalVar = "something...."

func validateUserInput(userName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

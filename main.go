package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to the %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("E-mail address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to buy? ")
	fmt.Scan(&userTickets)

	userTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v!\n", firstName, lastName, userTickets, email)
}
package main

import (
	"fmt"
)

// Main funtion is entrypoint, 1 main per project, -go mod init <app name > to create module, -go run <filename> to run
func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go COnference"
	//short hand to assign value, cant use for constants and declared tpys
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//uint wont accept negative values

	fmt.Printf("conferenceName is %T , remainingTickets is %T\n", conferenceName, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("Get your tickets here!", remainingTickets, "tickets available! out of ", conferenceTickets)
	fmt.Printf("Get your tickets here! %v tickets available! out of %v  \n", remainingTickets, conferenceTickets) //printf to print using placeholders

	var firstName string
	var lastName string
	var email string
	var usertickets int
	//ask user for name using pointer & to wait ofr input
	fmt.Println("Enter your First name :")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name :")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address name :")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets required:")
	fmt.Scan(&usertickets)
	fmt.Printf("Thank you %v %v for booking  %v tickets.\n You will recive confirmation email at %v\n", firstName, lastName, usertickets, email)

}

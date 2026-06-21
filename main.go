package main

import (
	"fmt"
)

// Main funtion is entrypoint, 1 main per project, -go mod init <app name > to create module, -go run <filename> to run
func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go COnference" //short hand to assign value, cant use for constants and declared tpys
	const conferenceTickets = 50
	var remainingTickets uint = 50 //uint wont accept negative values

	fmt.Printf("conferenceName is %T , remainingTickets is %T\n", conferenceName, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("Get your tickets here!", remainingTickets, "tickets available! out of ", conferenceTickets)
	fmt.Printf("Get your tickets here! %v tickets available! out of %v  \n", remainingTickets, conferenceTickets) //printf to print using placeholders

	var userName string
	var usertickets int
	//ask user for name
	userName = "Tom"
	usertickets = 2
	fmt.Printf("User %v booked %v tickets to the %v ", userName, usertickets, conferenceName)
}

package main

import (
	"fmt"
)

// Main funtion is entrypoint, 1 main per project, -go mod init <app name > to create module, -go run <filename> to run
func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("Get your tickets here!", remainingTickets, "tickets available! out of ", conferenceTickets)
	fmt.Printf("Get your tickets here! %v tickets available! out of %v  \n", remainingTickets, conferenceTickets) //printf to print using placeholders


    var userName
	//ask user for name 
	userName = "Tom"
}

package main

import (
	"fmt"
)

func main(){
	var eventsName = "GoLang Developer Bootcamp"
	fmt.Printf("Welcome to %v Events-Booking Command-Line Application\n", eventsName)
	var eventTickets uint= 50
	var remainingTickets uint = 50
	var bookings[]string

	fmt.Printf("We have total of %v tickets and %v are still remaining\n", eventTickets, remainingTickets)
	
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	//User Inputs - Name, email, tickets
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Email")
	fmt.Scan(&email)
	fmt.Println("How many tickets?")
	fmt.Scan(&userTickets)

	//Bookings, RemainingTickets
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets

	fmt.Println(firstName, lastName, email, userTickets)
	fmt.Println(bookings)
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %T\n", len(bookings))
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventsName)
}
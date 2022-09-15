package main

import (
	"events_golang/helper"
	"fmt"
	"strconv"
)

var eventName = "GoLang Conference"
var eventTickets uint= 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //initialize a list/array of map/hash

func main(){

	greetUser()
	
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//Book ticket in system only if the userinputs are validated
		if(isValidName && isValidEmail && isValidTicketNumber) {
			
			//Bookings, RemainingTickets
			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The firstname of bookings are: %v\n", firstNames)

			//exit the application if no tickets are left
			if remainingTickets == 0 {
				fmt.Println("Our event is fully booked out! Sorry try next time")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}
	}

	//switch statement example
	city:= "London"
	switch city {
		case "New York":
			fmt.Println("Eastern city")
		case "Singapore", "Hong Kong":
			fmt.Println("Asian city")
		case "London", "Berlin":
			fmt.Println("European city")
		default:
			fmt.Println("No valid city selected")
	}
}

func greetUser(){
	fmt.Printf("Welcome to %v Events-Booking Command-Line Application\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", eventTickets, remainingTickets)
}

func getFirstNames() [] string{
	firstNames :=[]string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking["firstName"]) //direct access to firstName due to object based storing of booking
	}
	return firstNames
}

func getUserInput() (string, string, string, uint){
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
	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint){
	
	remainingTickets = remainingTickets - userTickets

	//create a map for a suer
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) //now added the whole map object (previously it was just strings of firstName + lastName)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you for %v %v for booking %v tickets. Check your email at %v for confirmation\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, eventName)
			
}
// fmt.Printf("Slice type: %T\n", bookings)
// fmt.Printf("Slice length: %T\n", len(bookings))

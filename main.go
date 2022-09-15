package main

import (
	"fmt"
	"strings"
)

func main(){
	eventsName := "GoLang Developer Bootcamp"
	fmt.Printf("Welcome to %v Events-Booking Command-Line Application\n", eventsName)
	var eventTickets uint= 50
	var remainingTickets uint = 50
	bookings :=[]string{}

	fmt.Printf("We have total of %v tickets and %v are still remaining\n", eventTickets, remainingTickets)
	
	for {
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

		//validate user input
		isValidName := len(firstName) >=3 && len(lastName) >=3 
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//Book ticket in system only if the userinputs are validated
		if(isValidName && isValidEmail && isValidTicketNumber) {
			//Bookings, RemainingTickets
			bookings = append(bookings, firstName + " " + lastName)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you for %v %v for booking %v tickets. Check your email at %v for confirmation\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, eventsName)

			firstNames :=[]string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
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

	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %T\n", len(bookings))

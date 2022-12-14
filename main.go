package main

import (
	"events_golang/helper"
	"fmt"
	"sync"
)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

var eventName = "GoLang Conference"
var eventTickets uint= 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //initialize a list/array of map/hash

func main(){

	greetUser()
	
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	//Book ticket in system only if the userinputs are validated
	if(isValidName && isValidEmail && isValidTicketNumber) {		
		//Bookings, RemainingTickets
		bookTicket(firstName, lastName, email, userTickets)
			
		wg.Add(1) //just before the new thread
	
		//send confirmation by email - spawn on another thread - by just adding go in the front
		go sendTicket(userTickets, firstName, lastName, email) //just added "go" in front to make it concurrent and golang handles it
		firstNames := getFirstNames()
		fmt.Printf("The firstname of bookings are: %v\n", firstNames)

		//exit the application if no tickets are left
		if remainingTickets == 0 {
			fmt.Println("Our event is fully booked out! Sorry try next time")
			//break  //main for loop was removed to demo concurrency go routine with sync.WaitGroup
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
		//continue //main for loop was removed to demo concurrency go routine with sync.WaitGroup
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
	wg.Wait() //Blocks until the WaitGroup counter is 0 - so that all concurrent threads are cleared up before exiting
}

func greetUser(){
	fmt.Printf("Welcome to %v Events-Booking Command-Line Application\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", eventTickets, remainingTickets)
}

func getFirstNames() [] string{
	firstNames :=[]string{}
	for _, booking := range bookings{
		
		//using map
		//firstNames = append(firstNames, booking["firstName"]) //direct access to firstName due to object based storing of booking
		
		//using struct
		firstNames = append(firstNames, booking.firstName) //direct access to struct firstName
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
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) //now added the whole map object (previously it was just strings of firstName + lastName)
	
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you for %v %v for booking %v tickets. Check your email at %v for confirmation\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, eventName)
			
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	// time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}
// fmt.Printf("Slice type: %T\n", bookings)
// fmt.Printf("Slice length: %T\n", len(bookings))

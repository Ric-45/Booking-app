package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferencename = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidNamer, isValidTicketNumber, isValidEmail := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidNamer && isValidTicketNumber && isValidEmail {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first name of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is book out.Come back next year.")

			//break
		}

	} else {
		if !isValidNamer {
			fmt.Println("Your first name or lasgt name is too short")
		}
		if !isValidEmail {
			fmt.Println("Your email doesn't contain @ sigh")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you enter is invalid")
		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to  %v booking application\n", conferencename)
	fmt.Printf("We have total of %v Tickets and %v are still available\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking.You will recieve a confirmation mail at %v.\n", firstName, lastName, email)
	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferencename)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v  %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending tickect: \n %v \n to email address %v", ticket, email)
	fmt.Println("######################")
	wg.Done()
}

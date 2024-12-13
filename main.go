package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "AWS Re:Invent 2025"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		userFirstName, userLastName, userEmail, userTickets := getUserData()

		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, userFirstName, userLastName, userEmail)

			//Concurrent execution
			wg.Add(1)
			go sendTicket(userTickets, userFirstName, userLastName, userEmail)

			printFirstNames()

			if remainingTickets == 0 {
				fmt.Printf("The %v conference tickets are all booked. Please come back next year.\n", conferenceName)
				break
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Hello, welcome to %s booking application. Get your tickets here to attend...\n", conferenceName)
	fmt.Printf("There are a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("Thes first names of bookings are: %v\n", firstNames)
}

func getUserData() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets to book:")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets uint, userFirstName string, userLastName string, userEmail string) {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:    userFirstName,
		lastName:     userLastName,
		email:        userEmail,
		numOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}

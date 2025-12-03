package main

import (
	"fmt"
	"time"

	"Booking-app/helper"
)

const conferenceTickets = 50

var ConferenceName = "Go conference"

// ConferenceName := "Go conference"// package level variables can not be declared using :=
var remainingTickets uint = 50

// var bookings [50]string
// var bookings = make([]map[string]string, 0) // list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	//%T is used to know the type

	// fmt.Println("Welcome to", ConferenceName, "application")
	// fmt.Println("We have total", conferenceTickets, "tickets and remaining tickets are", remainingTickets)

	for {

		firstname, lastname, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValidTicket && isValidEmail && isValidName {
			bookTicket(userTickets, firstname, lastname, email)
			go sendTicket(userTickets, firstname, lastname, email)

			firstNames := printFirstNames()

			fmt.Printf("The firstNames of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("There are no tickets left")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
			continue
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have total %v tickets and remaining tickets are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string {
	firstnames := []string{}

	for _, booking := range bookings {

		firstnames = append(firstnames, booking.firstName)
	}

	return firstnames
}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	// ask user for the name
	fmt.Println("Enter your first name: ")

	fmt.Scan(&firstname)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTicket(userTickets uint, firstname string, lastname string, email string) {
	remainingTickets -= userTickets
	// array
	// bookings[0] = firstname + " " + lastname

	// var userData = make(map[string]string)
	// userData["FirstName"] = firstname
	// userData["LastName"] = lastname
	// userData["Email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstname,
		lastName:        lastname,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets . You will receive a confirmation email at %v \n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets are remaining \n", remainingTickets)

}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstname, lastname)
	println(">>>>>>>>>>>>>")
	fmt.Printf("Sending ticket:\n %v \n to email address:\n %v\n", ticket, email)
	println(">>>>>>>>>>>>>")
}

package main

import (
	"bookingApp/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("hello world")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTicket uint = 50

	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("Total no. of ticket %v and Total no. of ticket available %v \n", conferenceTickets, remainingTicket)

	time.Sleep(2 * time.Second)
	var isLoggedin bool = false
	var bookings []string

	for {

		DoLogin(&isLoggedin)

		if isLoggedin {
			for {
				var firstName string
				var lastName string
				var email string
				var userTicket uint
				fmt.Println("Welcome to the Go Conference App")
				fmt.Println("################################")
				fmt.Println("################################")

				fmt.Println("Enter your First Name ?")
				fmt.Scan(&firstName)
				fmt.Println("Enter your Last Name ?")
				fmt.Scan(&lastName)
				fmt.Println("Enter your email ?")
				fmt.Scan(&email)
				fmt.Println("Enter number of Tickets you would like to book ?")
				fmt.Scan(&userTicket)
				if userTicket > remainingTicket {
					fmt.Println("Number of Tickets can't be more than available, Available Ticket: ", remainingTicket)
					fmt.Println("Enter number of Tickets you would like to book ?")
					fmt.Scan(&userTicket)
				}

				remainingTicket = remainingTicket - userTicket
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Println("Entered Name is : ", bookings[0], " Email Entered :", email, " Booked no. of Tickets is : ", userTicket)
				fmt.Printf("Total no. of ticket %v and Total no. of ticket available %v \n", conferenceTickets, remainingTicket)

				// Map key value pair with strict data type
				var userData = make(map[string]string)
				userData["firstName"] = firstName
				userData["lastName"] = lastName
				userData["email"] = email
				userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)

				// Struct for key value pair with loose data type
				userDataStruct := utils.DataStruct{
					FirstName:  firstName,
					LastName:   lastName,
					Email:      email,
					UserTicket: userTicket,
				}

				wg.Add(1)
				go utils.SendNotification(userData, &wg, userDataStruct)

				if remainingTicket < 1 {
					break
				}
			}

			fmt.Println("All Tickets have been sold out.")
			for index, booking := range bookings {
				// var fname = strings.Split(booking, " ")[0]
				var fname = strings.Fields(booking)[0]
				fmt.Println("No. ", index+1, " ", fname)
			}

			break
		}
	}
	wg.Wait()
}

func DoLogin(isLoggedin *bool) {
	var loginUser string
	var loginPass string

	fmt.Println("Enter your login name :")
	fmt.Scan(&loginUser)
	fmt.Println("Enter your password :")
	fmt.Scan(&loginPass)

	var isusercorrect = Login(loginUser, loginPass) // ✅ directly accessible
	// fmt.Printf("login response %v\n", isusercorrect)
	*isLoggedin = isusercorrect
}

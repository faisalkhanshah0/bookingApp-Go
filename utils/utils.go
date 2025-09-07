package utils

import (
	"fmt"
	"sync"
	"time"
)

type DataStruct struct {
	FirstName  string
	LastName   string
	Email      string
	UserTicket uint
}

func SendNotification(userData map[string]string, wg *sync.WaitGroup, userDataStruct DataStruct) string {
	time.Sleep(5 * time.Second)
	fmt.Println("sending notification to user ", userData["firstName"], " on ", userData["email"], " no of tickets booked : ", userData["userTicket"])
	time.Sleep(2 * time.Second)
	fmt.Println("Message Sent to user ", userDataStruct.FirstName, " on ", userDataStruct.Email, " no of tickets booked : ", userDataStruct.UserTicket)
	defer wg.Done()
	return "completed"
}

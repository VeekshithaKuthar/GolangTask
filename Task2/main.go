package main

import (
	"fmt"
	"math/rand"
	"os"
	"task2/models"
	"task2/operation"
)

// Insert and Get operations of a User (Id, Name, Email, Status)

// When you insert a record, it should be logged into a file log info (
//
//	(date and time, action("insert/get"))

func NewInsert(id int, name string, email string, status string) *models.User {
	return &models.User{Id: id, Name: name, Email: email, Status: status}
}
func main() {
	fmt.Println("Enter a Option of Your Choice\n 1.Insert\n 2.Get\n 3.Exit")
	var n int
	n = 2
	fmt.Scanln(&n)
	for {
		switch n {

		case 1:
			ch := make(chan models.LogInfo, 10)
			var name, email string
			fmt.Print("Enter your name:")
			fmt.Scanln(&name)
			fmt.Print("Enter your email:")
			fmt.Scanln(&email)
			user := &models.User{
				Id:     rand.Int(),
				Name:   name,
				Email:  email,
				Status: "Active",
			}
			err := operation.InsertUser(user, ch)
			if err != nil {
				fmt.Println("Error in inserting file")
			} else {
				fmt.Println("Data Saved Successfully")
			}

			operation.InsertLogDetails(ch)
		case 2:
			ch := make(chan models.LogInfo, 10)
			var email string
			fmt.Print("Enter  Email to find a details:")
			fmt.Scanln(&email)
			operation.GetUserByEmail(email, ch)
			operation.InsertLogDetails(ch)
		default:
			os.Exit(0)
		}
	}
}

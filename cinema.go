package main

import (
	"fmt"
	"os"
	"sync"
)

var once sync.Once
var authenticator *Authenticator

var users = make(map[string]User)
var factory = MovieFactory{}
var movies = []Movie{
	factory.CreateMovie("action", "Leon", 10),
	factory.CreateMovie("comedy", "Man In Black", 8),
	factory.CreateMovie("drama", "Seven Pounds", 7),
}
var currentUser User

func main() {
	authenticator := GetAuthenticator()

	fmt.Println("Welcome to the Cinema Ticket Booking System")

	for {
		fmt.Println("\n1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Please select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login(authenticator)
		case 2:
			register(authenticator)
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

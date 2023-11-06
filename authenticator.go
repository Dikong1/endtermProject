package main

import (
	"fmt"
	"sync"
)

type Authenticator struct {
	observers []Observer
}

type AuthenticationLogger struct{}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) AddUser(observer Observer) {
	a.observers = append(a.observers, observer)
}

var once sync.Once
var authenticator *Authenticator

func GetAuthenticator() *Authenticator {

	once.Do(func() {
		authenticator = NewAuthenticator()
		authenticator.AddUser(AuthenticationLogger{})
	})
	return authenticator
}

func login(authenticator *Authenticator) {
	fmt.Println("\nLogin:")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	var password string
	fmt.Scanln(&password)

	user, exists := users[username]
	if !exists || user.Password != password {
		fmt.Println("Invalid username or password. Please try again.")
		return
	}

	currentUser = user
	authenticator.NotifyUsers(currentUser)
	fmt.Printf("Welcome, %s!\n", currentUser.Username)
	movieListing(ListingStrategy{}, movies)
}

func register(authenticator *Authenticator) {
	fmt.Println("\nRegister:")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	if username == "" {
		fmt.Println("Username cannot be empty")
		return
	}
	if len(username) < 4 || len(username) > 16 {
		fmt.Println("Length of username must be in range 4 to 16 ")
		return
	}

	fmt.Print("Password: ")
	var password string
	fmt.Scanln(&password)

	if password == "" {
		fmt.Println("Password cannot be empty")
		return
	}
	if len(password) < 4 || len(password) > 16 {
		fmt.Println("Length of password must be in range 4 to 16 ")
		return
	}

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists. Please choose a different username.")
		return
	}

	users[username] = User{Username: username, Password: password}
	fmt.Printf("Registration successful. Welcome, %s!\n", username)
	currentUser = users[username]
	login(authenticator)
}

func (a *Authenticator) NotifyUsers(user User) {
	for _, observer := range a.observers {
		observer.Update(user)
	}
}

func (l AuthenticationLogger) Update(user User) {
	fmt.Printf("User %s logged in.\n", user.Username)
}

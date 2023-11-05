package main

import (
	"fmt"
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

func GetAuthenticator() *Authenticator {
	once.Do(func() {
		authenticator = NewAuthenticator()
		authenticator.AddUser(AuthenticationLogger{})
	})
	return authenticator
}

func login() {
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
	fmt.Printf("Welcome, %s!\n", currentUser.Username)
	movieListing(ListingStrategy{}, movies)
}

func register() {
	fmt.Println("\nRegister:")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	var password string
	fmt.Scanln(&password)

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists. Please choose a different username.")
		return
	}

	users[username] = User{Username: username, Password: password}
	fmt.Printf("Registration successful. Welcome, %s!\n", username)
	currentUser = users[username]
	movieListing(ListingStrategy{}, movies)
}

func (a *Authenticator) NotifyUsers(user User) {
	for _, observer := range a.observers {
		observer.Update(user)
	}
}

func (l AuthenticationLogger) Update(user User) {
	fmt.Printf("User %s logged in.\n", user.Username)
}

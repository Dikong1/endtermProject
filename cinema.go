package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Observer Pattern
type Observer interface {
	Update(user User)
}

type User struct {
	Username string
	Password string
}

type Authenticator struct {
	observers []Observer
}

func (a *Authenticator) Attach(observer Observer) {
	a.observers = append(a.observers, observer)
}

func (a *Authenticator) Notify(user User) {
	for _, observer := range a.observers {
		observer.Update(user)
	}
}

type AuthenticationLogger struct{}

func (l AuthenticationLogger) Update(user User) {
	fmt.Printf("User %s logged in.\n", user.Username)
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

// Strategy Pattern
type MovieListing interface {
	ListMovies(movies []Movie)
}

type ByGenreListingStrategy struct{}

func (s ByGenreListingStrategy) ListMovies(movies []Movie) {
	fmt.Println("\nMovie Listing (Grouped by Genre):")
	genres := make(map[string][]Movie)

	for _, movie := range movies {
		genres[movie.Genre] = append(genres[movie.Genre], movie)
	}

	for genre, movies := range genres {
		fmt.Printf("Genre: %s\n", genre)
		for _, movie := range movies {
			fmt.Printf("Title: %s, Tickets Available: %d\n", movie.Title, movie.TicketsAvailable)
		}
	}
}

type SimpleListingStrategy struct{}

func (s SimpleListingStrategy) ListMovies(movies []Movie) {
	fmt.Println("\nMovie Listing:")
	fmt.Println("Title\tGenre\tTickets Available")
	for _, movie := range movies {
		fmt.Printf("%s\t%s\t%d\n", movie.Title, movie.Genre, movie.TicketsAvailable)
	}
}

// Factory Method Pattern
type MovieFactory interface {
	CreateMovie(title, genre string, ticketsAvailable int) Movie
}

type StandardMovieFactory struct{}

func (f StandardMovieFactory) CreateMovie(title, genre string, ticketsAvailable int) Movie {
	return Movie{
		Title:            title,
		Genre:            genre,
		TicketsAvailable: ticketsAvailable,
	}
}

type PremiumMovieFactory struct{}

func (f PremiumMovieFactory) CreateMovie(title, genre string, ticketsAvailable int) Movie {
	return Movie{
		Title:            title,
		Genre:            genre,
		TicketsAvailable: ticketsAvailable,
	}
}

// Singleton Pattern
var once sync.Once
var authenticator *Authenticator

func GetAuthenticator() *Authenticator {
	once.Do(func() {
		authenticator = NewAuthenticator()
		authenticator.Attach(AuthenticationLogger{})
	})
	return authenticator
}

// Movie struct
type Movie struct {
	Title            string
	Genre            string
	TicketsAvailable int
}

// Login function
func login(authenticator *Authenticator) {
	fmt.Println("\nLogin:")
	fmt.Print("Username: ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	username = username[:len(username)-1]

	fmt.Print("Password: ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	password = password[:len(password)-1]

	user, exists := users[username]
	if !exists || user.Password != password {
		fmt.Println("Invalid credentials. Please try again.")
		return
	}

	currentUser = user
	fmt.Printf("Welcome, %s!\n", currentUser.Username)
	movieListing(ByGenreListingStrategy{}, movies)
}

// Register function
func register(authenticator *Authenticator) {
	fmt.Println("\nRegister:")
	fmt.Print("Username: ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	username = username[:len(username)-1]

	fmt.Print("Password: ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	password = password[:len(password)-1]

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists. Please choose a different username.")
		return
	}

	users[username] = User{Username: username, Password: password}
	fmt.Printf("Registration successful. Welcome, %s!\n", username)
	currentUser = users[username]
	movieListing(SimpleListingStrategy{}, movies)
}

// Movie listing function
func movieListing(strategy MovieListing, movies []Movie) {
	for {
		strategy.ListMovies(movies)

		fmt.Print("Enter the title of the movie you want to book tickets for (or type 'exit' to log out): ")
		var movieTitle string
		fmt.Scanln(&movieTitle)

		if movieTitle == "exit" {
			currentUser = User{}
			fmt.Println("Logged out.")
			return
		}

		found := false
		for i, movie := range movies {
			if movie.Title == movieTitle {
				if movie.TicketsAvailable > 0 {
					movies[i].TicketsAvailable--
					fmt.Printf("You have booked a ticket for %s.\n", movie.Title)
					found = true
					break
				} else {
					fmt.Printf("Sorry, no tickets available for %s.\n", movie.Title)
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Printf("Movie with title '%s' not found.\n", movieTitle)
		}
	}
}

var users = make(map[string]User)
var movies = []Movie{
	{"Mission-Impossible", "Action", 10},
	{"Home-Alone", "Comedy", 8},
	{"Titanic", "Drama", 15},
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

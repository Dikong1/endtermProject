package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Movie interface {
	getTitle() string
	getGenre() string
	GetTickets() int
	SetTickets(int)
}

type MovieListing interface {
	ListMovies(movies []Movie)
}

type actionMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *actionMovie) getTitle() string   { return m.Title }
func (m *actionMovie) getGenre() string   { return m.genre }
func (m *actionMovie) GetTickets() int    { return m.TicketsAvailable }
func (m *actionMovie) SetTickets(num int) { m.TicketsAvailable -= num }

type comedyMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *comedyMovie) getTitle() string   { return m.Title }
func (m *comedyMovie) getGenre() string   { return m.genre }
func (m *comedyMovie) GetTickets() int    { return m.TicketsAvailable }
func (m *comedyMovie) SetTickets(num int) { m.TicketsAvailable -= num }

type dramaMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *dramaMovie) getTitle() string   { return m.Title }
func (m *dramaMovie) getGenre() string   { return m.genre }
func (m *dramaMovie) GetTickets() int    { return m.TicketsAvailable }
func (m *dramaMovie) SetTickets(num int) { m.TicketsAvailable -= num }

type MovieFactory struct{}

func (f MovieFactory) CreateMovie(movieType string, title string, tickets int) Movie {
	switch movieType {
	case "action":
		return &actionMovie{Title: title, genre: "Action", TicketsAvailable: tickets}
	case "comedy":
		return &comedyMovie{Title: title, genre: "Comedy", TicketsAvailable: tickets}
	case "drama":
		return &dramaMovie{Title: title, genre: "Drama", TicketsAvailable: tickets}
	default:
		return nil
	}
}

type ListingStrategy struct{}

func (s ListingStrategy) ListMovies(movies []Movie) {
	fmt.Println("\nMovie Listing:")
	fmt.Printf("%-25s\t%-15s\t%-20s\n", "Title", "Genre", "Tickets Available")
	fmt.Println("--------------------------------------------------------------------")
	for _, movie := range movies {
		fmt.Printf("%-25s\t%-15s\t%-20d\n", movie.getTitle(), movie.getGenre(), movie.GetTickets())
	}
}

func movieListing(strategy MovieListing, movies []Movie) {
	for {
		strategy.ListMovies(movies)

		fmt.Print("Enter the title of the movie you want to book tickets for (or type 'exit' to log out): ")
		movieTitle, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		movieTitle = strings.TrimSpace(movieTitle)

		if movieTitle == "exit" {
			currentUser = User{}
			fmt.Println("Logged out.")
			break
		}

		found := false
		for _, movie := range movies {
			if movie.getTitle() == movieTitle {
				var numOfTickets int
				fmt.Print("How many ticket you want to book: ")
				fmt.Scanln(&numOfTickets)
				if movie.GetTickets()-numOfTickets > 0 {
					movie.SetTickets(numOfTickets)
					fmt.Println("--------------------------------------------")
					fmt.Printf("You have booked %d tickets for %s.\n", numOfTickets, movie.getTitle())
					fmt.Println("--------------------------------------------")
					found = true
					break
				} else {
					fmt.Printf("Sorry, only %d tickets available for %s.\n", movie.GetTickets(), movie.getTitle())
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

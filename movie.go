package main

import "fmt"

type Movie interface {
	getTitle() string
	getGenre() string
	GetTickets() int
	SetTickets()
}

type MovieListing interface {
	ListMovies(movies []Movie)
}

type actionMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *actionMovie) getTitle() string { return m.Title }
func (m *actionMovie) getGenre() string { return m.genre }
func (m *actionMovie) GetTickets() int  { return m.TicketsAvailable }
func (m *actionMovie) SetTickets()      { m.TicketsAvailable-- }

type comedyMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *comedyMovie) getTitle() string { return m.Title }
func (m *comedyMovie) getGenre() string { return m.genre }
func (m *comedyMovie) GetTickets() int  { return m.TicketsAvailable }
func (m *comedyMovie) SetTickets()      { m.TicketsAvailable-- }

type dramaMovie struct {
	Title            string
	genre            string
	TicketsAvailable int
}

func (m *dramaMovie) getTitle() string { return m.Title }
func (m *dramaMovie) getGenre() string { return m.genre }
func (m *dramaMovie) GetTickets() int  { return m.TicketsAvailable }
func (m *dramaMovie) SetTickets()      { m.TicketsAvailable-- }

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
	fmt.Println("Title\tGenre\tTickets Available")
	for _, movie := range movies {
		fmt.Printf("%s\t%s\t%d\n", movie.getTitle(), movie.getGenre(), movie.GetTickets())
	}
}

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
		for _, movie := range movies {
			if movie.getTitle() == movieTitle {
				if movie.GetTickets() > 0 {
					movie.SetTickets()
					fmt.Printf("You have booked a ticket for %s.\n", movie.getTitle())
					found = true
					break
				} else {
					fmt.Printf("Sorry, no tickets available for %s.\n", movie.getTitle())
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

package main

type Observer interface {
	Update(user User)
}

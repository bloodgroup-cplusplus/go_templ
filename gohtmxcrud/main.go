package main


import "github.com/go-chi/chi/v5"

type Item struct {
	ID string
	Name string
	Description string
	ImagePath string
}



// In memory database
var items =  make(map[string]Item)

func main() {

// we need a chi router

	r := 




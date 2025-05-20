package main


import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
)

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

// Middleware setup
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// middleware for cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:[]string,
		AllowedMethods:[]string{"GET","POST","PUT","DELETE","OPTIONS"},
	}))

	// create uploads dirctory for storing images
	os.MkdirAll("uploads",os.ModePerm)

	r.Handle("/uploads/*",http.StripPrefix("/uploads",http.FileServer
	((http.Dir("uploads")))))

	// Route definitions

		






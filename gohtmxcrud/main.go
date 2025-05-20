package main


import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"fmt"
	"log"
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
	// 1 GET server the main/home page
	// 2 GET Return list of items
	// 3 POST Create a new item
	// 4 DELTE delte an item with an ID
	// Start the server 
	fmt.Println("Server starting on http: locahost:300")
	log.Fatal(http.ListenAndServe(":3000",r))

}
		






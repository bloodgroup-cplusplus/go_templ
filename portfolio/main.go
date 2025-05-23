package main

import (
	"fmt"
	"html/template" // render html templates
	"log" // logging errors
	"net/http" 
)



func main() {
	// handling the request 
	http.HandleFunc("/",home)
	
}

//global scope
//home function

func home (w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"home.html")
}


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

func renderTemplate(w http.ResponseWriter, tmpl string) ( {
	// Parsing the specified template file being passed as input
	t, err := template.ParseFiles("templates/"+tmpl)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	t.Execute(w,nil)
}


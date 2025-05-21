package main
import (
	"net/http"
	"fmt"
	"html/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl,err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return 
		}
		// anonymous structs
		data := struct {
			Message string
		}{
			Message:"Hello world",
		}
		tmpl.Execute(w,data)
			
	})

	fmt.Println("We started our http server on port :8080")
	http.ListenAndServe(":8080",nil)


}

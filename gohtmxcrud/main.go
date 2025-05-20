package main
import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
		AllowedOrigins: []string{},
		AllowedMethods:[]string{"GET","POST","PUT","DELETE","OPTIONS"},
	}))
	// create uploads dirctory for storing images
	// Route definiteis
	os.MkdirAll("uploads",os.ModePerm)
	r.Handle("/uploads/*",http.StripPrefix("/uploads",http.FileServer((http.Dir("uploads")))))
	r.Get("/",handleIndex)
	r.Get("/items",handleListItems)
	fmt.Println("Server starting on http: locahost:300")
	log.Fatal(http.ListenAndServe(":3000",r))

}
		


//handleindex function to serve the main page of the applciation
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl : template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w,items)
}

// handleList Items
func handleListItems( w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/items-list.html"))
	tmpl.Execute(w,items)
}

//handle create Items 

func handleCreateItem(w http.ResponseWriter, r *http.Request ) {
	r.ParseMultiPartForm(10 <<20) // 10 mb max
	name := r.FromValue("name")
	description := r.FormValue("description")
	id := fmt.Sprintf("%d",len(items)+1)
	// handle image file upload 
	file,handler, err := r.FormFile("image")
	var  imagePath string
	if err == nil {
		defer file.Close()
		filename := filepath.join("uploads",handler.filename)
		f, err := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE,0666)
		if err == nil {
			defer f.close()
			io.Copy(f,file)
			imagePath = "/uploads/" + handler.Filename
		}
	}
}

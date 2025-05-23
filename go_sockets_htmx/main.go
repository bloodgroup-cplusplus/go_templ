package main

import (
	"log"
	"net/http"
	"fmt"
)

func serveIndex( w http.ResponseWriter, r * http.Request) {
	if r.URL.Path != "/" {
		http.Error(w,"not found",http.StatusNotFound)
		return 
	}
	if r.Method!="GET"{
		http.Error(w,"not found",http.StatusNotFound)
		return 
	}

	http.ServeFile(w,r,"templates/index.html")
}


func main()  {
	http.HandleFunc("/",serveIndex)
	http.HandleFunc("/ws",func (w http.ResponseWriter, r *http.Request ) {
		serveWs(hub,w,r)
	})

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000",nil))
}

package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/bloodgroup-cplusplus/go_templ/htmx_gin_notes_app/controllers"
)


func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("views/**/*")
	r.GET("/",controllers.NotesIndex)
	r.POST("/notes", controllers.NotesCreate)
	log.Println("Server started!")
	r.Run() // default port 8080
}

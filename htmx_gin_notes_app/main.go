package main

import (
	"log"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("views/**/*")
	r.GET("/",controllers.NotesIndex)
	log.Println("Server started!")
	r.Run() // default port 8080

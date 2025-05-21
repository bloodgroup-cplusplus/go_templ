package controllers

import (
	"github.com/bloodgroup-cplusplus/go_templ/htmx_gin_notes_app/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func NotesIndex( c * gin.Context) {
	notes :=[] models.Note {
	{
		Name: "Bug squashing",
		Content:"chad"
	},
	{
		Name:"chad",
		Content:"chad",
	},
	{
		Name: "Chad",
		Content: "chad",
	}
	}

	c.HTML (
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes":notes,
		},
	)
}

type FormData struct {
	Name string `form:"name"`
	Content string `form:"content"`
}

func NotesCreate(c *gin.Context) {
	var data FormData
	c.Bind(&data)
}


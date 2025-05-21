package controllers

import (
	"go_htmx/models"
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



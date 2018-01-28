package controllers

import (
	"net/http"
	"twitter-meme-bot/structs"
	"twitter-meme-bot/database"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	thread := structs.Thread{}
	database.DB.Last(&thread)

	type ResponseData struct {
		CreatedAt string
		Thread structs.Thread
	}

	renderView(w, r, "app.html", "home.html", ResponseData{
		CreatedAt: thread.Model.CreatedAt.Format("Jan 2, at 15:04pm"),
		Thread: thread,
	})
}
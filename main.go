package main

import (
	"github.com/joho/godotenv"

	"twitter-meme-bot/database"
	"twitter-meme-bot/reddit"
	"twitter-meme-bot/twitter"

	"net/http"
	"fmt"
	"os"
)

func main() {
	// register .env variables
	err := godotenv.Load()
	if err != nil {
		print("Error loading .env file continuing with normal os env variables")
	}

	database.Connect()
	twitter.Setup();

	println("Bot starting...")

	go reddit.GetThreads()

	// Http server so we can deploy on Heroku
	println("Web server listening on :"+os.Getenv("PORT"));
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}
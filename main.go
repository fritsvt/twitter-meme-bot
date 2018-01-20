package main

import (
	"github.com/joho/godotenv"

	"twitter-meme-bot/database"
	"twitter-meme-bot/reddit"
	"twitter-meme-bot/twitter"

	"log"
	"time"
	"net/http"
	"fmt"
	"os"
)

func main() {
	// register .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	twitter.Setup();

	println("Bot starting...")

	go loopInterval(60) // start the loop

	// Http server so we can deploy on Heroku
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func loopInterval(interval time.Duration) {
	twitter.TweetThreads(*reddit.GetThreads())

	for range time.Tick(time.Second * interval){
		twitter.TweetThreads(*reddit.GetThreads())
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}
package main

import (
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"

	"twitter-meme-bot/database"
	"twitter-meme-bot/reddit"
	"twitter-meme-bot/twitter"
	"twitter-meme-bot/web"

	"net/http"
	"os"
	"time"
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

	go twitter.StartStream()

	go loopInterval(10)

	// http server so we can deploy on Heroku
	println("Web server listening on :"+os.Getenv("PORT"));

	r := mux.NewRouter()
	r = web.RegisterRoutes(*r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/static/")))

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func loopInterval(interval time.Duration) {
	reddit.GetThreads(false)

	for range time.Tick(time.Second * interval){
		reddit.GetThreads(true)
	}
}
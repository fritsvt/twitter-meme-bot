package controllers

import (
	"net/http"
	"os"
	"net/url"
	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
)

func GetNewSchedule(w http.ResponseWriter, r *http.Request) {
	scheduledTweets := []structs.ScheduledTweet{}
	database.DB.Order("created_at desc").Limit(100).Find(&scheduledTweets)

	renderView(w, r, "app.html", "schedule_new_tweet.html", scheduledTweets)
}

func PostNewSchedule(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("password") != os.Getenv("ADMIN_PASSWORD") {
		setFlash("Invalid password", "danger", w, r)
		redirectBack(w, r)
		return
	}
	imageUrl := r.FormValue("img_url")
	tweetTitle := r.FormValue("tweet")
	username := r.FormValue("username")

	if !isValidUrl(imageUrl) {
		setFlash("Image url must be a valid url", "danger", w, r)
		redirectBack(w, r)
		return
	}

	if len(tweetTitle) > 255 {
		setFlash("Tweet title cannot be greater then 255 characters", "danger", w, r)
		redirectBack(w, r)
		return
	}

	if len(username) < 1 || len(username) > 255 {
		setFlash("Username must be between 1 and 255 characters", "danger", w, r)
		redirectBack(w, r)
		return
	}

	database.DB.Create(&structs.ScheduledTweet{
		ImageUrl:imageUrl,
		Title:tweetTitle,
		ToUser:username,
	})

	setFlash("Tweet scheduled to @"+username, "success", w, r)
	http.Redirect(w, r, "/schedule", 301)
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func redirectBack(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/schedule", 301)
}
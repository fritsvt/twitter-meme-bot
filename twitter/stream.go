package twitter

import (
	"net/url"
	"github.com/ChimeraCoder/anaconda"
	"fmt"
	"os"
	"twitter-meme-bot/structs"
	"twitter-meme-bot/database"
	"time"
)

type LastTweet struct {
	ScreenName string
	Timestamp int64
}

func StartStream() {
	lastTweet := LastTweet{}

	if os.Getenv("TWITTER_STREAM") != "true" {
		return
	}

	println("Starting streaming twitter feed...")

	v := url.Values{}

	stream := api.UserStream(v)
	for tweet := range stream.C {
		switch v := tweet.(type) {
		case anaconda.Tweet:
			diff := time.Now().Unix() - lastTweet.Timestamp

			if lastTweet.ScreenName != v.User.ScreenName && diff > 60 {
				lastTweet = LastTweet{
					ScreenName:v.User.ScreenName,
					Timestamp:time.Now().Unix(),
				}
				checkSchedule(v.User.ScreenName, v.IdStr)
			}
			fmt.Printf("%-15s: %s\n", v.User.ScreenName, v.Text)
		case anaconda.EventTweet:
			switch v.Event.Event {
			case "favorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("Favorited by %-15s: %s\n", sn, tw)
			case "unfavorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("UnFavorited by %-15s: %s\n", sn, tw)
			}
		}
	}
}

func checkSchedule(username string, statusId string) {
	scheduledTweet := structs.ScheduledTweet{}
	database.DB.Where("to_user = ?", username).Order("created_at desc").First(&scheduledTweet)

	if scheduledTweet.ID != 0 {
		thread := structs.Thread{
			Title:scheduledTweet.Title,
			Extension:"jpg",
			ImageUrl:scheduledTweet.ImageUrl,
			RedditId:scheduledTweet.Model.CreatedAt.String(),
			ImageHash:"",
		}
		scheduledTweet.StatusId = statusId
		SendTweet(thread, false, &scheduledTweet)
		database.DB.Delete(&scheduledTweet)
	}
}

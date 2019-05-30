package twitter

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
)

type LastTweet struct {
	ID int64
	ScreenName string
	Timestamp int64
}

var lastTweet = LastTweet{}

func StartStream() {
	if os.Getenv("TWITTER_STREAM") != "true" {
		return
	}

	// Fetch feed every 10 secs
	for range time.Tick(time.Second * 61) {
		fetchFeed()
	}
}

func fetchFeed() {
	println("Fetching latest timeline")

	sinceID := "1"
	if lastTweet.ID != 0 {
		sinceID = strconv.FormatInt(lastTweet.ID, 10)
	}
	// fetch from api
	res, _ := api.GetHomeTimeline(url.Values{
		"count": {"50"},
		"since_id": { sinceID },
		"exclude_replies": {"true"},
	})

	for _, tweet := range res {
		layout := "Mon Jan 2 15:04:05 -0700 2006"

		t, err := time.Parse(layout, tweet.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		timestamp := t.UnixNano()

		// new tweet
		if lastTweet.Timestamp < timestamp {
			lastTweet = LastTweet{
				ID: tweet.Id,
				ScreenName: tweet.User.ScreenName,
				Timestamp: timestamp,
			}

			// check new tweet against schedule
			go checkSchedule(strings.ToLower(tweet.User.ScreenName), tweet.IdStr)
		}
	}
}

func checkSchedule(username string, statusId string) {
	scheduledTweet := structs.ScheduledTweet{}
	database.DB.Where("to_user = ?", username).Order("created_at asc").First(&scheduledTweet)

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

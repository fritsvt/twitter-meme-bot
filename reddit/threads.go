package reddit

import (
	"os"
	"fmt"
	"strconv"
	"time"
	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
	"github.com/turnage/graw/reddit"
	"path/filepath"
	"twitter-meme-bot/twitter"
)

var tweetThreads = true
var queuedTweet = structs.Thread{}

func GetThreads(tweet bool) {
	tweetThreads = tweet
	fmt.Println("fetching threads from /r/" + os.Getenv("SUB_REDDIT") + "/" + os.Getenv("REDDIT_SORT"))

	cfg := reddit.BotConfig{
		Agent: "graw:twitter-meme-uouse:1.0.1 by /u/zwembadsniper",
		App: reddit.App{
			ID:     os.Getenv("REDDIT_APP_ID"),
			Secret: os.Getenv("REDDIT_APP_SECRET"),
			Username: os.Getenv("REDDIT_USERNAME"),
			Password: os.Getenv("REDDIT_PASSWORD"),
		},
	}
	bot, err := reddit.NewBot(cfg)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	harvest, err := bot.Listing("/r/"+os.Getenv("SUB_REDDIT")+"/"+os.Getenv("REDDIT_SORT"), "")
	if err != nil {
		fmt.Println("Failed to fetch /r/"+os.Getenv("SUB_REDDIT")+": ", err)
		return
	}

	for _, post := range harvest.Posts[:25] {
		//fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
		filterThread(post)
	}
}

func filterThread (post *reddit.Post) error {
	extension := filepath.Ext(post.URL)

	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" && extension != ".gif" {
		return nil
	}

	threadTitle := post.Title
	if len(threadTitle) > 190 {
		threadTitle = threadTitle[:190]
	}
	thread := structs.Thread{
		ImageUrl: post.URL,
		Title:    threadTitle,
		RedditId:       post.ID,
		Author:   post.Author,
		Extension:extension,
		Upvotes: post.Ups,
	}
	//if queuedTweet.Title != "" || database.GetThreadById(post.ID) == false {
		if tweetThreads {
			// check if this tweet has higher upvotes than the one in memory
			if queuedTweet.Upvotes < thread.Upvotes && database.GetThreadById(post.ID) == false {
				queuedTweet = thread
			}

			lastThread := structs.Thread{}
			database.DB.Last(&lastThread)

			// if it has been an hour
			interval, _ := strconv.ParseInt(os.Getenv("MIN_TWEET_INTERVAL"), 10, 64)
			if (time.Now().Unix() - lastThread.Timestamp) >= interval && queuedTweet.Title != "" {
				twitter.SendTweet(queuedTweet, true, &structs.ScheduledTweet{})
				queuedTweet = structs.Thread{}
			}
		} else {
			if database.GetThreadById(post.ID) == false {
				database.InsertThread(thread)
				println("Inserting but not tweeting " + thread.RedditId)
			}
		}
	//}
	return nil
}
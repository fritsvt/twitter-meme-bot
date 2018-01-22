package reddit

import (
	"os"
	"fmt"
	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
	"github.com/turnage/graw/reddit"
	"path/filepath"
	"twitter-meme-bot/twitter"
)

func GetThreads() {
	fmt.Println("fetching threads from /r/" + os.Getenv("SUB_REDDIT"))

	cfg := reddit.BotConfig{
		Agent: "graw:twitter-meme-uouse:1.0.1 by /u/zwembadsniper",
		// Your registered app info from following:
		// https://github.com/reddit/reddit/wiki/OAuth2
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

	harvest, err := bot.Listing("/r/"+os.Getenv("SUB_REDDIT")+"/controversial", "")
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
	}
	if database.GetThreadById(post.ID) == false {
		database.InsertThread(thread)
		twitter.SendTweet(thread)
	}
	return nil
}
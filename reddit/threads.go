package reddit

import (
	"os"
	"log"
	"fmt"
	"twitter-meme-bot/database"
	"twitter-meme-bot/structs"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"time"
	"path/filepath"
	"twitter-meme-bot/twitter"
)

type filterThreads struct{}

func GetThreads() (*structs.Thread) {
	startRedditStream()

	return &structs.Thread{}
}

func startRedditStream() {
	println("Starting streaming posts from /r/r" + os.Getenv("SUB_REDDIT"))
	// Get an api handle to reddit for a logged out (script) program,
	apiHandle, err := reddit.NewScript("meme-house", 5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Create a configuration specifying what event sources on Reddit graw
	cfg := graw.Config{Subreddits: []string{os.Getenv("SUB_REDDIT")}}

	// launch a graw scan in a goroutine using the bot handle
	_, wait, err := graw.Scan(&filterThreads{}, apiHandle, cfg)
	if err := wait(); err != nil {
		fmt.Printf("graw run encountered an error: %v\n", err)
	}
}

func (a *filterThreads) Post(post *reddit.Post) error {
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
		Id:       post.ID,
		Author:   post.Author,
	}
	if database.GetThreadById(post.ID) == false {
		database.InsertThread(thread)
		twitter.SendTweet(thread)
	}
	return nil
}
package database

import "twitter-meme-bot/structs"

func GetTweetToUser(username string) (tweet structs.ScheduledTweet) {
	scheduledTweet := structs.ScheduledTweet{}

	DB.Where("to_user = ?", username).Order("created_at desc").First(&scheduledTweet)

	return scheduledTweet
}

func InsertTweetToUser(tweet structs.ScheduledTweet) {
	DB.Create(&structs.ScheduledTweet{
		ImageUrl:tweet.ImageUrl,
		Title:tweet.Title,
		ToUser:tweet.ToUser,
	})
}
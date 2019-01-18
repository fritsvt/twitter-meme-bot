package database

import (
	"time"
	"twitter-meme-bot/structs"
)

func GetThreadById(id string) (bool) {
	var thread structs.Thread
	DB.First(&thread, "reddit_id = ?", id)

	if thread.ID != 0 {
		return true
	}
	return false
}

func InsertThread(thread structs.Thread) {
	DB.Create(&structs.Thread{
		RedditId:thread.RedditId,
		ImageUrl:thread.ImageUrl,
		Title:thread.Title,
		Author:thread.Author,
		ImageHash:thread.ImageHash,
		Upvotes: thread.Upvotes,
		Timestamp: time.Now().Unix(),
	})
}

func GetThreadByHash(hash string) (bool) {
	var thread structs.Thread
	DB.First(&thread, "image_hash = ?", hash)
	if thread.ID != 0 {
		return true
	}
	return false;
}
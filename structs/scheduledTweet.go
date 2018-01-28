package structs

import "github.com/jinzhu/gorm"

type ScheduledTweet struct {
	gorm.Model
	ImageUrl string `json:"image_url"`
	StatusId string `json:"status_id"`
	Title string `json:"title"`
	ToUser string `json:"to_user"`
}
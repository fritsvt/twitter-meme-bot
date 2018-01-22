package structs

import "github.com/jinzhu/gorm"

type Thread = struct {
	gorm.Model
	ImageUrl string `json:"image_url"`
	Title string `json:"title"`
	RedditId string `json:"id"`
	Author string `json:"author"`
}
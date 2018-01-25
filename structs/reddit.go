package structs

import "github.com/jinzhu/gorm"

type Thread = struct {
	gorm.Model
	ImageUrl string `json:"image_url"`
	Extension string `json:"extension"`
	Title string `json:"title"`
	RedditId string `json:"id"`
	Author string `json:"author"`
	ImageHash string `json:"image_hash"`
}
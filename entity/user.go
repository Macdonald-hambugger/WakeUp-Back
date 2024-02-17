package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid     string `json:"uuid"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Profile  string `json:"profile_image_url"`
}

package model

import "github.com/jinzhu/gorm"

//Content is a tag to be used in a page
type Content struct {
	gorm.Model
	Slug  string `gorm:"index"`
	Title string `gorm:"size:255"`
	Body  string `gorm:"type:text;"`
}

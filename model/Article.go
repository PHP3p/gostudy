package model

import "github.com/jinzhu/gorm"

type Article struct {
	category Category
	gorm.Model
	Title   string `gorm:"type:varchar(120); not null" json:"title"`
	Desc    string `gorm:"type:varchar(200); not null" json:"desc"`
	Content string `gorm:"type:text; not null" json:"content"`
}

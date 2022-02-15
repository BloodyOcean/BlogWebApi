package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserId uint   `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`

	// Set up one-to-many relationship.
	Comments []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	gorm.Model
	PostId uint   `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var db *gorm.DB
var err error

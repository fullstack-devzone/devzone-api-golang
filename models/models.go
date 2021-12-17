package models

import "time"

type Link struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	CreatedDate time.Time `json:"created_date"`
}

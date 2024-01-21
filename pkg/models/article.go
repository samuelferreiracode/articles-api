package models

type Article struct {
	ID      int64  `json:"id" bson:"id"`
	Title   string `json:"title" bson:"title"`
	Author  string `json:"author" bson:"author"`
	Content string `json:"content" bson:"content"`
}

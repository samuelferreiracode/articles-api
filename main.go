package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var articles = []article{
	{ID: 1, Title: "First Article", Author: "Samuel Ferreira", Content: "Just a test article."},
	{ID: 2, Title: "Second Article", Author: "Samuel Ferreira", Content: "Just a test article."},
}

func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

func createArticle(c *gin.Context) {
	var newArticle article
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, "Article created successfully")
}

func articleById(c *gin.Context) {
	id := c.Param("id")
	number, _ := strconv.ParseInt(id, 10, 0)
	article, err := getArticleById(number)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, article)
}

func getArticleById(id int64) (*article, error) {
	for i, a := range articles {
		if a.ID == id {
			return &articles[i], nil
		}
	}

	return nil, errors.New("Article not found")
}

func main() {
	router := gin.Default()
	router.GET("/articles", getArticles)
	router.GET("/articles/:id", articleById)
	router.POST("/articles", createArticle)
	router.Run("localhost:8080")
}

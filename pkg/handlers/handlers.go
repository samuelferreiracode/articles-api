package handlers

import (
	"example/first-go-api/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRoutes(router *gin.Engine, client *repository.Client) {
	router.GET("/articles", getPaginatedArticles(client))
	router.GET("articles/:id", getArticleByID(client))
}

func getArticleByID(client *repository.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		number, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
			return
		}

		article, err := client.GetArticleByID(number)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
			return
		}

		c.IndentedJSON(http.StatusOK, article)
	}
}

func getPaginatedArticles(client *repository.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		pageSizeStr := c.DefaultQuery("pageSize", "15")

		page, err := strconv.ParseInt(pageStr, 10, 64)
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil || pageSize < 1 {
			pageSize = 15
		}

		articles, err := client.GetPaginatedArticles(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, articles)
	}
}

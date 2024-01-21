package main

import (
	"context"
	"example/first-go-api/pkg/config"
	"example/first-go-api/pkg/handlers"
	"example/first-go-api/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	dbClient := repository.NewClient(cfg.MongoDBURI, cfg.Database)
	defer func() {
		if err := dbClient.DB.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	handlers.RegisterRoutes(r, dbClient)
	r.Run(":8080")
}

//func createArticle(c *gin.Context) {
//	var newArticle repository.Article
//	if err := c.BindJSON(&newArticle); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, err := collection.InsertOne(ctx, newArticle)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.IndentedJSON(http.StatusCreated, "Article created successfully")
//}

//func articleById(c *gin.Context) {
//	id := c.Param("id")
//	number, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
//		return
//	}
//
//	article, err := getArticleById(number)
//	if err != nil {
//		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, article)
//}

//func getArticleById(id int64) (*article, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	var art article
//	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&art)
//	if err != nil {
//		return nil, errors.New("Article not found")
//	}
//
//	return &art, nil
//}

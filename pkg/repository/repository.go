package repository

import (
	"context"
	"example/first-go-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Client struct {
	mongoClient *mongo.Client
	DB          *mongo.Database
}

type Article = models.Article

func NewClient(mongoURI, dbName string) *Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	return &Client{
		DB: client.Database(dbName),
	}
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.mongoClient.Disconnect(ctx)
}

func (c *Client) GetPaginatedArticles(page, pageSize int64) ([]Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSkip((page - 1) * pageSize).SetLimit(pageSize)
	cursor, err := c.DB.Collection("articles").Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []Article
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

func (c *Client) GetArticleByID(id int64) (*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var art Article
	err := c.DB.Collection("articles").FindOne(ctx, bson.M{"id": id}).Decode(&art)
	if err != nil {
		return nil, err
	}

	return &art, nil
}

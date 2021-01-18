package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client    *mongo.Client
	dbConnURL string
)

//MongoDBConfig structure
type MongoDBConfig struct {
	Username     string
	Password     string
	DatabaseName string
	DatabasePort string
	URLDB        string
	Timeout      time.Duration
}

//NewMongoClient create new client
func NewMongoClient(mongoConf *MongoDBConfig) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), mongoConf.Timeout*time.Second)

	dbConnURL = fmt.Sprintf("mongodb://%s:%s", mongoConf.URLDB, mongoConf.DatabasePort)

	client, _ = mongo.NewClient(options.Client().ApplyURI(dbConnURL))

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("Cannot connect to database")
	}

	return client, nil
}

//NewMongoDatabase Create new db
func NewMongoDatabase(cl *mongo.Client, mongoConf *MongoDBConfig) *mongo.Database {
	// Create DB Connection

	db := cl.Database(mongoConf.DatabaseName)

	return db
}

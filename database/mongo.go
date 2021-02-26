package database

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoClient struct {
	instance *mongo.Client
}

const (
	databaseName    = "ingilizce"
	collectionWords = "words"
)

var mongoClient MongoClient

func (MongoClient) setup() {
	dbConfig := config.Get().Database.MongoConfig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("mongodb://%s:%d/", dbConfig.Host, dbConfig.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Panic(ErrDbConnection)
	}

	mongoClient.instance = client

	err = mongoClient.instance.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(ErrPing)
	}
	/*
		i, err := mongoClient.instance.Database(DatabaseName).Collection(CollectionWords).InsertOne(context.TODO(), bson.M{"bismillah": "b"})
		fmt.Println(i.InsertedID)
	*/
}

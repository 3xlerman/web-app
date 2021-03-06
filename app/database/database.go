package database

import (
	"context"
	"fmt"
	"github.com/3xlerman/web-app/app/config"
	"github.com/3xlerman/web-app/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var usersCollection *mongo.Collection

func Connect() {

	// Client connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://3xlerman:" + config.MongoDBPassword + "@lermancluster.3bezk.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Context connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Checking errors
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// Show databases list
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	usersCollection = client.Database("lermanDB").Collection("lermanCollection")
}

func InsertUser(user model.User) (err error) {
	_, err = usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

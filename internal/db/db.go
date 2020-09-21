package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client mongo.Client

func DBconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	u := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbString := "mongodb+srv://" + u + ":" + pass + "@graphql.jl2u2.mongodb.net/" + dbName + "retryWrites=true&w=majority"

	client1, err := mongo.Connect(ctx, options.Client().ApplyURI(dbString))
	Client = *client1
	if err != nil {
		log.Fatal("erro: ", err)
	}
	fmt.Println(Client.ListDatabaseNames, "DB CONNECTED")

	// collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
	// cur, err := collection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(result)
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}

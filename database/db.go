package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://meahamad1346:13464631@cluster0.w3o5xs9.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// important
var Collection *mongo.Collection //get the refrenced db collection

//connect with mango db

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mango db
	client, err := mongo.Connect(context.TODO(), clientOption) //its the context of the connection

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The connection success")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready!")
}

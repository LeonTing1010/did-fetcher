package config

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Client = connectDB()
	db               = MongoDB()
)

func SetUp() {
	fileName := Env("LOG_FILE_NAME")
	// open log file
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Panic(err)
	}
	// defer logFile.Close()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.Output(1, "starting fetchers")
	log.Println("USE DB:", db)
}

func connectDB() *mongo.Client {
	log.Println("Connecting to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(databases)

	return client
}

// getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Database(db).Collection(collectionName)
	return collection
}

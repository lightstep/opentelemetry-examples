package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"mongodb-batch/pkg/generator"
	"sync"
	"time"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	usersCollection := client.Database("otel").Collection("users")

	ticker := time.NewTicker(time.Second * 3)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for range ticker.C {
			n := time.Now()
			// insert the bson object slice using InsertMany()
			dummyProducts := generator.GenerateDummyProducts(10)
			results, err := usersCollection.InsertMany(context.TODO(), dummyProducts)
			// check for errors in the insertion
			if err != nil {
				panic(err)
			}
			// display the ids of the newly inserted objects
			fmt.Printf("Inserting %d took %s\n", 10, time.Since(n))
			fmt.Println(results.InsertedIDs)
		}
	}()
	wg.Wait()
	time.Sleep(time.Second * 4)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

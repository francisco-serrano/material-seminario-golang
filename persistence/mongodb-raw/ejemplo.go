package mongodb_raw

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"math"
	"time"
)

func MongoDisconnect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func RunMongo() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	uri := "mongodb://root:root@localhost:27017/local?authSource=admin"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer MongoDisconnect(client, ctx)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database("local").Collection("sample-collection")

	result, err := collection.InsertOne(ctx, bson.M{"msg": "hello"})
	if err != nil {
		panic(err)
	}

	fmt.Println("result", result)

	values := []interface{}{
		bson.M{"name": "pi", "value": math.Pi},
		bson.M{"name": "pi", "value": math.Pi},
		bson.M{"name": "pi", "value": math.Pi},
		bson.M{"name": "pi", "value": math.Pi},
		bson.M{"name": "pi", "value": math.Pi},
		bson.M{"name": "pi", "value": math.Pi},
	}

	if _, err := collection.InsertMany(ctx, values); err != nil {
		panic(err)
	}

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(ctx) {
		var element bson.D

		if err := cursor.Decode(&element); err != nil {
			panic(err)
		}

		fmt.Println("element", element)
	}
}

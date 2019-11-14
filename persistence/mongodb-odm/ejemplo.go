package mongodb_odm

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

func RunMongoODM() {
	config := &bongo.Config{
		ConnectionString: "mongodb://root:root@localhost:27017/local?authSource=admin",
		Database:         "local",
	}

	connection, err := bongo.Connect(config)
	if err != nil {
		panic(err)
	}

	results := connection.Collection("sample-collection").Find(bson.M{})

	var elem bson.D
	for results.Next(&elem) {
		fmt.Println(elem)
	}
}

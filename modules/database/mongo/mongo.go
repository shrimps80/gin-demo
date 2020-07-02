package mongo

import (
	"fmt"
	"time"
	"context"
	"gin-demo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongodb struct {
	Conn *mongo.Client
}

var Client Mongodb

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	host := fmt.Sprintf("mongodb://%s", config.GetEnv().MongodbHost)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		panic(fmt.Errorf("Connect %s error: %s", host, err))
	}
	cancel()
	Client.Conn = client
}

func (m *Mongodb) SetCollection(name string) *mongo.Collection {
	return m.Conn.Database(config.GetEnv().MongodbName).Collection(name)
}

func (m *Mongodb) InsertOne(name string, val interface{}) string {
	collection := m.SetCollection(name)
	res, err := collection.InsertOne(context.TODO(), val)
	if err != nil {
		panic(err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex()
}

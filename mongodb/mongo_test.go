package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

type User struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name" bson:"name"`
	Age        int       `json:"age" bson:"age"`
	Tags       []string  `json:"tags" bson:"tags"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
}

// TestInsert增加
func TestInsert(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		user := &User{
			Id:         uuid.New().String(),
			Name:       fmt.Sprintf("刘德华%d", i),
			Age:        10 + i,
			Tags:       []string{"科技", "文学"},
			CreateTime: time.Now(),
		}
		result, err := client.Database("testing").Collection("users").InsertOne(context.Background(), &user)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.InsertedID.(string))
	}

}

// TestMongoSearch搜索
func TestMongoSearch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("users")
	query := bson.M{
		"name": bson.M{
			"$gt": "",
		},
	}
	opt := options.Find().SetLimit(10).SetSort(bson.M{"_id": 1})
	cur, err := collection.Find(context.Background(), &query, opt)
	if err != nil {
		panic(err)
	}

	var user []User
	err = cur.All(context.Background(), &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}

// TestDeleteMongo 册除数据
func TestDeleteMongo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("users")

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": "4e1f5ce4-2be9-4084-9277-70d16e0b549a"})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.DeletedCount)

}

// TestUpdateMongo 更新文档
func TestUpdateMongo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("users")

	query := bson.M{"_id": "21d6f563-6656-42d1-b192-1a56938e2379"}
	update := bson.M{
		"$set": bson.M{
			"name": "周华建",
		},
	}
	opt := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(context.Background(), query, update, opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.MatchedCount)

}

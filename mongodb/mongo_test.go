/*
 * :file description:
 * :name: /mongodb/mongo_test.go
 * :author: 张德志
 * :copyright: (c) 2023, Tungee
 * :date created: 2023-10-10 22:07:58
 * :last editor: 张德志
 * :date last edited: 2023-10-11 17:07:54
 */
package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		result, err := client.Database("testing").Collection("users").
			InsertOne(context.Background(), &user)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.InsertedID.(string))
	}

}

//func TestMongoSearch(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		panic(err)
//	}
//	collection := client.Database("testing").Collection("user")
//	query := bson.M{
//		"name": bson.M{
//			"$gt": " ",
//		},
//	}
//	opt := options.Find().SetLimit(10).SetSort(bson.M{"_id": 1})
//	cur, err := collection.Find(context.Background(), &query, opt)
//	if err != nil {
//		panic(err)
//	}
//	var user []User
//	err = cur.All(context.Background(), &user)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(user)
//}

// 单个查询
//func TestMongoSearch(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		panic(err)
//	}
//	collection := client.Database("testing").Collection("users")
//	filter := bson.D{{Key: "_id", Value: "65303cb7-3cf2-474a-a4a5-bcc9c3c46617"}}
//
//	var user User
//	err = collection.FindOne(context.Background(), &filter).Decode(&user)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(user)
//}

// 单个查询
func TestMongoSearch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("users")
	query := bson.D{{Key: "_id", Value: "47b68974-8ed6-44ab-ad4f-0c1c72a03d40"}}
	var user User
	err = collection.FindOne(context.Background(), &query).Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}

func TestDeleteMongo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("testing").Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": "47b68974-8ed6-44ab-ad4f-0c1c72a03d40"})
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
	query := bson.M{"_id": "062a267e-e569-4cba-b4cf-a2ebce4d185d"}
	update := bson.M{"$set": bson.M{
		"name": "hello change info",
	}}
	opt := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(context.Background(), query, update, opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ModifiedCount)
}

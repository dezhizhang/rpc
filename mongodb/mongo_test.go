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

// TestInsert插入数据
//func TestInsert(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		panic(err)
//	}
//
//	err1 := client.Ping(context.Background(), &readpref.ReadPref{})
//	if err1 != nil {
//		panic(err1)
//	}
//
//	// 插入
//	//for i := 0; i < 10; i++ {
//	//	user := &User{
//	//		Id:         uuid.New().String(),
//	//		Name:       fmt.Sprintf("刘德华%d", i),
//	//		Age:        10 + i,
//	//		Tags:       []string{"tag"},
//	//		CreateTime: time.Now(),
//	//	}
//	//	result, err := client.Database("testing").Collection("users").InsertOne(context.Background(), &user)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//	fmt.Println(result.InsertedID.(string))
//	//}
//
//	//----------------查询
//	//coll := client.Database("testing").Collection("users")
//	//query := bson.M{
//	//	"name": bson.M{
//	//		"$gt": "",
//	//	},
//	//}
//	//opt := options.Find().SetLimit(2).SetSort(bson.M{"_id": 1})
//	//cur, err := coll.Find(context.Background(), &query, opt)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//var user []User
//	//err = cur.All(context.Background(), &user)
//	//if err != nil {
//	//	panic(err1)
//	//}
//	//fmt.Println(user)
//	//----------------更新
//	//coll := client.Database("testing").Collection("users")
//	//query := bson.M{"_id": "3e22698b-edde-4c19-bdb8-b8fcf6461710"}
//	//update := bson.M{
//	//	"$set": bson.M{
//	//		"name": "周华建",
//	//	},
//	//}
//	//opt := options.Update().SetUpsert(true)
//	//one, err := coll.UpdateOne(context.Background(), query, update, opt)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Println(one)
//	// -----------------删除数据
//	coll := client.Database("testing").Collection("users")
//	one, err := coll.DeleteOne(context.Background(), bson.M{
//		"_id": "e7d3c61f-2abf-4594-8a9c-38915fa98af6",
//	})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(one.DeletedCount)
//
//}

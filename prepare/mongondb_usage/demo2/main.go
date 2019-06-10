package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {

	var (
		client       *mongo.Client
		err          error
		database     *mongo.Database
		collection   *mongo.Collection
		record       *LogRecord
		insertResult *mongo.InsertOneResult
	)

	if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetConnectTimeout(time.Second * 5)); err != nil {
		fmt.Println(err)
		return
	}

	client.Connect(context.TODO())

	//fmt.Println(client)

	database = client.Database("cron")

	collection = database.Collection("log")

	record = &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Err:     "",
		Content: "hello",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}

	if insertResult, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
	}

	docID := insertResult.InsertedID.(primitive.ObjectID)

	fmt.Println("自增ID：", docID.Hex())

}

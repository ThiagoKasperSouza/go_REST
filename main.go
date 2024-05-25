package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis-14918.c8.us-east-1-3.ec2.redns.redis-cloud.com:14918",
	Username: "default",

	Password: "zCAtUdbQ12&", // no password set
	DB:       0,             // use default DB
})

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	type donation struct {
		Id   string `json:"id"`
		Nome string `json:"nome"`
		Link string `json:"link"`
	}
	ds := []donation{
		{
			Id:   uuid.New().String(),
			Nome: "teste",
			Link: "testelink1",
		},
		{
			Id:   uuid.New().String(),
			Nome: "teste2",
			Link: "testelink2",
		},
	}

	for _, d := range ds {
		res, _ := json.Marshal(d)

		rdb.HSet(context.Background(), "donations", d.Id, res)
	}

	donations := []*donation{}
	val, err := rdb.HGetAll(context.Background(), "donations").Result()
	for _, item := range val {
		donation := &donation{}
		err := json.Unmarshal([]byte(item), donation)
		fmt.Println(donation.Nome)
		if err != nil {
			panic(err)
		}
		donations = append(donations, donation)

	}

}

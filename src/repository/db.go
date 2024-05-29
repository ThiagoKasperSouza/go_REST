package repository

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ADDR"),
		Username: os.Getenv("USR"),
		Password: os.Getenv("DB_P"), // no password set
		DB:       0,                 // use default DB
	})

}

type DbClient struct {
	Instance *redis.Client
}

func (rdb *DbClient) Create(key string, Id string, data []byte) *redis.IntCmd {
	log.Default().Printf("C - %s  %s\n", Id, data)
	return rdb.Instance.HSet(context.Background(), key, Id, data)
}

func (rdb *DbClient) GetAll(key string) *redis.MapStringStringCmd {
	log.Default().Printf("L - %s\n", key)
	return rdb.Instance.HGetAll(context.Background(), key)
}

func (rdb *DbClient) GetItemById(key string, Id string) *redis.StringCmd {
	log.Default().Printf("L - %s by id %s\n", key, Id)
	return rdb.Instance.HGet(context.Background(), key, Id)
}

func (rdb *DbClient) Update(key string, Id string, data []byte) *redis.IntCmd {
	log.Default().Printf("U - %s by id %s\n", key, Id)
	return rdb.Instance.HSet(context.Background(), key, Id, data)
}

func (rdb *DbClient) Delete(key string, Id string) *redis.IntCmd {
	log.Default().Printf("D - %s by id %s\n", key, Id)
	return rdb.Instance.HDel(context.Background(), key, Id)
}

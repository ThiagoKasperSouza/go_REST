package repository

import (
	"context"
	"github.com/segmentio/ksuid"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

/*
Copyright 2024 Thiago Kasper de Souza

This file is part of go_rest.

go_rest is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

go_rest is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with go_rest.  If not, see <https://www.gnu.org/licenses/>
*/

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

var Rdb = DbClient{
	Instance: GetClient(),
}

func (rdb *DbClient) Create(key string, Id ksuid.KSUID, data []byte) *redis.IntCmd {
	log.Default().Printf("C - %s  %s\n", Id, data)
	return rdb.Instance.HSet(context.Background(), key, Id, data)
}

func (rdb *DbClient) GetAll(key string) *redis.MapStringStringCmd {
	log.Default().Printf("L - %s\n", key)
	return rdb.Instance.HGetAll(context.Background(), key)
}

func (rdb *DbClient) GetItemById(key string, Id ksuid.KSUID) *redis.StringCmd {
	log.Default().Printf("L - %s by id %s\n", key, Id)
	return rdb.Instance.HGet(context.Background(), key, Id.String())
}

func (rdb *DbClient) Update(key string, Id ksuid.KSUID, data []byte) *redis.IntCmd {
	log.Default().Printf("U - %s by id %s\n", key, Id)
	return rdb.Instance.HSet(context.Background(), key, Id, data)
}

func (rdb *DbClient) Delete(key string, Id ksuid.KSUID) *redis.IntCmd {
	log.Default().Printf("D - %s by id %s\n", key, Id)
	return rdb.Instance.HDel(context.Background(), key, Id.String())
}

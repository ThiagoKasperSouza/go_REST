package steps

import (
	"context"
	"encoding/json"
	r "newsRestFiber/src/repository"

	"github.com/google/uuid"
)

var rdb = r.DbClient{
	Instance: r.GetClient(),
}

type example struct {
	Id   string `json:"id"`
	Nome string `json:"nome"`
	Link string `json:"link"`
}

var ex = example{
	Id:   uuid.New().String(),
	Nome: "example",
	Link: "examplelink1",
}

func ClientExists() error {
	return rdb.Instance.ClientInfo(context.Background()).Err()
}

func CreateObj() error {

	res, jsonErr := json.Marshal(ex)
	if jsonErr != nil {
		return jsonErr
	}
	_, err := rdb.Create("example", ex.Id, res).Result()
	return err
}

func GetObj() error {
	_, err := rdb.GetItemById("example", ex.Id).Result()
	return err
}

func GetAll() error {
	_, err := rdb.GetAll("example").Result()
	return err
}

func Update() error {
	var up = ex
	up.Nome = "update"
	res, jsonErr := json.Marshal(up)
	if jsonErr != nil {
		return jsonErr
	}
	_, err := rdb.Update("example", up.Id, res).Result()
	return err
}

func Delete() error {
	_, jsonErr := json.Marshal(ex)
	if jsonErr != nil {
		return jsonErr
	}
	_, err := rdb.Delete("example", ex.Id).Result()
	return err
}

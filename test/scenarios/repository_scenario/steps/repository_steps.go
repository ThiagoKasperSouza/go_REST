package steps

import (
	"context"
	"encoding/json"
	r "go_Rest/src/repository"
	"log"
	"time"

	"github.com/segmentio/ksuid"
)

var rdb = r.DbClient{
	Instance: r.GetClient(),
}

type example struct {
	Id   ksuid.KSUID `json:"id"`
	Nome string      `json:"nome"`
	Link string      `json:"link"`
}

func handleError(err error, message string) error {
	if err != nil {
		log.Default().Println(message + " - " + err.Error())
		return err
	}
	return err
}

func getKsUid() ksuid.KSUID {
	id, kErr := ksuid.NewRandomWithTime(time.Now())
	handleError(kErr, "ksUid Error")
	return id
}

var ex = example{
	Id:   getKsUid(),
	Nome: "Teste",
	Link: "https://google.com",
}

func ClientExists() error {
	return rdb.Instance.ClientInfo(context.Background()).Err()
}

func CreateObj() error {

	res, jsonErr := json.Marshal(ex)
	handleError(jsonErr, "json Marshal err")
	_, err := rdb.Create("example", ex.Id, res).Result()
	return handleError(err, "create err")
}

func GetObj() error {
	_, err := rdb.GetItemById("example", ex.Id).Result()
	return handleError(err, "getObj err")
}

func GetAll() error {
	_, err := rdb.GetAll("example").Result()
	return handleError(err, "getAll err")
}

func Update() error {
	var up = ex
	up.Nome = "update"
	res, jsonErr := json.Marshal(up)
	handleError(jsonErr, "jsonErr")
	_, err := rdb.Update("example", up.Id, res).Result()
	return handleError(err, "Update err")
}

func Delete() error {
	_, jsonErr := json.Marshal(ex)
	handleError(jsonErr, "jsonErr")

	_, err := rdb.Delete("example", ex.Id).Result()
	return handleError(err, "delete err")
}

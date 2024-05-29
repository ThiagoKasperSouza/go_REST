package steps

import (
	"encoding/json"
	r "newsRestFiber/repository"

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

func CreateObj() error {

	res, error := json.Marshal(ex)
	rdb.Create("example", ex.Id, res)
	return error
}

func GetObj() error {
	_, err := rdb.GetItemById("example", ex.Id)
	return err
}

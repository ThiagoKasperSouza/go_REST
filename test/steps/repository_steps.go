package steps

import (
	"context"
	"encoding/json"
	r "newsRestFiber/repository"

	"github.com/google/uuid"
)

type example struct {
	Id   string `json:"id"`
	Nome string `json:"nome"`
	Link string `json:"link"`
}

var rdb = r.DbClient{
	Instance: r.GetClient(),
}

func CreateObj(ctx context.Context) error {
	ex := example{
		Id:   uuid.New().String(),
		Nome: "example",
		Link: "examplelink1",
	}
	res, error := json.Marshal(ex)
	rdb.Create("example", ex.Id, res)
	return error
}

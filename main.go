package main

import (
	"encoding/json"
	r "newsRestFiber/src/repository"
	m "newsRestFiber/src/repository/models"

	"github.com/google/uuid"
)

func main() {
	rdb := r.DbClient{
		Instance: r.GetClient(),
	}
	d := m.Donation{
		Id:   uuid.New().String(),
		Nome: "TESTE",
		Link: "link1",
	}
	res, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	d.Create(rdb, "donations", d.Id, res)

}

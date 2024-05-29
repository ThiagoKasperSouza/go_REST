package main

import (
	"encoding/json"
	"fmt"
	r "newsRestFiber/repository"

	"github.com/google/uuid"
)

func main() {
	rdb := r.DbClient{
		Instance: r.GetClient(),
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

		rdb.Create("donations", d.Id, res)
	}

	donations := []*donation{}
	var lastId string
	val, _ := rdb.GetAll("donations")
	for _, item := range val {
		donation := &donation{}
		err := json.Unmarshal([]byte(item), donation)
		lastId = donation.Id
		fmt.Println(donation.Nome)
		if err != nil {
			panic(err)
		}
		donations = append(donations, donation)

	}
	res, _ := rdb.GetItemById("donations", lastId)
	fmt.Print(res)

	updateD := donation{
		Id:   lastId,
		Nome: "NomeUpdated",
		Link: "LinkUpdated",
	}
	jsonUp, _ := json.Marshal(updateD)
	var upRes = rdb.Update("donations", lastId, jsonUp)
	fmt.Print(upRes)

	var delRes = rdb.Delete("donations", lastId)
	fmt.Print(delRes)

}

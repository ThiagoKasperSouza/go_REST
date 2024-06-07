package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	middlewares "newsRestFiber/src/middlewares"
	repository "newsRestFiber/src/repository"
	models "newsRestFiber/src/repository/models"

	"github.com/google/uuid"
)

/*
Copyright 2024 Thiago Kasper de Souza

This file is part of rsNews_blogApi.

rsNews_blogApi is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

rsNews_blogApi is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with rsNews_blogApi.  If not, see <https://www.gnu.org/licenses/>
*/

const (
	GET  = "GET "
	POST = "POST "
)

func main() {
	rdb := repository.DbClient{
		Instance: repository.GetClient(),
	}

	router := http.NewServeMux()

	stack := middlewares.CreateStack(
		middlewares.Headers,
		middlewares.Logging,
	)

	router.HandleFunc(POST+"/donations/create", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var t models.Donation
		t.Id = uuid.New().String()
		decoder.Decode(&t)

		res, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		t.Create(rdb, "donations", t.Id, res)

		json.NewEncoder(w).Encode(t)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on http://localhost:8080")
	server.ListenAndServe()

	// d := m.Donation{
	// 	Id:   uuid.New().String(),
	// 	Nome: "TESTE",
	// 	Link: "link1",
	// }
	// res, err := json.Marshal(d)
	// if err != nil {
	// 	panic(err)
	// }
	// d.Create(rdb, "donations", d.Id, res)

}

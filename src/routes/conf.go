package routes

import (
	"encoding/json"
	"errors"
	"go_Rest/src/repository/models"
	"log"
	"net/http"
	"regexp"
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

const (
	GET     = "GET "
	POST    = "POST "
	PUT     = "PUT "
	PATCH   = "PATCH "
	DELETE  = "DELETE "
	OPTIONS = "OPTIONS "
)

type Router struct {
	mux             *http.ServeMux
	collection_name string
}

type ErrorResponse struct {
	err    error
	status int
}

func ValidateLinkRegex(w http.ResponseWriter, t models.Info) {
	match, regex_err := regexp.Match("^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/|\\/|\\/\\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$", []byte(t.Link))
	if regex_err != nil || match == false {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Invalid Link error"),
			status: http.StatusBadRequest,
		})
	}
}

func HandleError(err error, w http.ResponseWriter, message string) {
	if err != nil {
		log.Default().Println(message + " - " + err.Error())
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Could not finish request"),
			status: http.StatusBadRequest,
		})
	}
}

package models

import "github.com/segmentio/ksuid"

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

type Info struct {
	CrudModel
	Id   ksuid.KSUID `json:"id"`
	Nome string      `json:"nome"`
	Link string      `json:"link"`
}

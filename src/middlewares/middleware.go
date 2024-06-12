package middlewares

import "net/http"

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

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		//retorna o proximo da lista
		for i := len(middlewares) - 1; i >= 0; i-- {
			current := middlewares[i]
			next = current(next)
		}
		return next
	}

}

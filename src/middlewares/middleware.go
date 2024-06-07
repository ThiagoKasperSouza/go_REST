package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		//retorna o proximo da lista
		for i := len(middlewares); i >= 0; i-- {
			current := middlewares[i]
			next = current(next)
		}
		return next
	}

}

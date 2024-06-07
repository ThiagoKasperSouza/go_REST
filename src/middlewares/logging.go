package middlewares

import (
	"log"
	"net/http"
	"time"
)

type wrapper struct {
	http.ResponseWriter
	statuscode int
}

func (w *wrapper) WriteHeader(statuscode int) {
	w.ResponseWriter.WriteHeader(statuscode)
	w.statuscode = statuscode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &wrapper{
			ResponseWriter: w,
			statuscode:     http.StatusOK,
		}
		next.ServeHTTP(w, r)
		log.Println(r.Method, wrapper.statuscode, r.URL.Path, time.Since(start))
	})
}

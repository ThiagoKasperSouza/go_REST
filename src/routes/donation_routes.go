package routes

import (
	"encoding/json"
	"net/http"
	"newsRestFiber/src/repository"
	"newsRestFiber/src/repository/models"

	"github.com/google/uuid"
)

type DonationRouter struct {
	mux *http.ServeMux
}

func (router *DonationRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Donation
	t.Id = uuid.New().String()
	decoder.Decode(&t)

	res, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	t.Create(repository.Rdb, "donations", t.Id, res)

	json.NewEncoder(w).Encode(t)
}

func RegisterDonationRoutes(prefix string, mux *http.ServeMux) {
	dr := &DonationRouter{
		mux: mux,
	}
	dr.mux.HandleFunc(prefix+"/create", dr.Create)
}

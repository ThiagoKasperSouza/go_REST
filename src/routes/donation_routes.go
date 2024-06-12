package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"newsRestFiber/src/conf"
	"newsRestFiber/src/repository"
	"newsRestFiber/src/repository/models"

	"github.com/google/uuid"
)

type DonationRouter struct {
	mux *http.ServeMux
}
type ErrorResponse struct {
	err    error
	status int
}

const collection_name = "donations"

func (router *DonationRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Donation
	t.Id = uuid.New().String()
	decoder.Decode(&t)

	res, err := json.Marshal(t)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Serialization error"),
			status: http.StatusBadRequest,
		})
	}
	t.Create(repository.Rdb, collection_name, t.Id, res)

	json.NewEncoder(w).Encode(t)
}

func (router *DonationRouter) GetItemById(w http.ResponseWriter, r *http.Request) {
	var t models.Donation

	res, err := t.GetItemById(repository.Rdb, collection_name, r.PathValue("id"))
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Bad req error"),
			status: http.StatusBadRequest,
		})
	}
	json.NewEncoder(w).Encode(res)
}

func (router *DonationRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	list, err := t.GetAll(repository.Rdb, collection_name)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Can't Get any list"),
			status: http.StatusBadRequest,
		})
	}
	json.NewEncoder(w).Encode(list)
}

func (router *DonationRouter) Update(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&t)
	jsonD, err := json.Marshal(t)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Serialization error"),
			status: http.StatusBadRequest,
		})
	}

	res, err := t.Update(repository.Rdb, collection_name, r.PathValue("id"), jsonD)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    err,
			status: http.StatusBadRequest,
		})
	}
	json.NewEncoder(w).Encode(res)
}

func (router *DonationRouter) Delete(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	res, err := t.GetItemById(repository.Rdb, collection_name, r.PathValue("id"))
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    err,
			status: http.StatusBadRequest,
		})
	}
	json.NewEncoder(w).Encode(res)
}

func RegisterDonationRoutes(prefix string, mux *http.ServeMux) {
	dr := &DonationRouter{
		mux: mux,
	}
	dr.mux.HandleFunc(conf.POST+prefix+"/create", dr.Create)
	dr.mux.HandleFunc(conf.POST+prefix+"/{id}", dr.GetItemById)
	dr.mux.HandleFunc(conf.GET+prefix+"/", dr.GetAll)
	dr.mux.HandleFunc(conf.PUT+prefix+"/update/{id}", dr.Update)
	dr.mux.HandleFunc(conf.DELETE+prefix+"/delete/{id}", dr.Delete)
}

package routes

import (
	"encoding/json"
	"errors"
	"github.com/segmentio/ksuid"
	"log"
	"net/http"
	"newsRestFiber/src/repository"
	"newsRestFiber/src/repository/models"
	"regexp"
	"time"
)

type DonationRouter struct {
	mux *http.ServeMux
}
type ErrorResponse struct {
	err    error
	status int
}

const collection_name = "donations"

func validateLinkRegex(w http.ResponseWriter, t models.Donation) {
	match, regex_err := regexp.Match("^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/|\\/|\\/\\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$", []byte(t.Link))
	if regex_err != nil || match == false {
		json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Invalid Link error"),
			status: http.StatusBadRequest,
		})
	}
}

func handleError(err error, w http.ResponseWriter, message string) {
	if err != nil {
		log.Default().Println(message + " - " + err.Error())
		_ = json.NewEncoder(w).Encode(&ErrorResponse{
			err:    errors.New("Could not finish request"),
			status: http.StatusBadRequest,
		})
	}
}

func (router *DonationRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Donation
	id, kErr := ksuid.NewRandomWithTime(time.Now())
	handleError(kErr, w, "ksUid Error")
	t.Id = id
	decoder.Decode(&t)

	res, err := json.Marshal(t)
	handleError(err, w, "Serialization Error")

	validateLinkRegex(w, t)

	_, cErr := t.Create(repository.Rdb, collection_name, t.Id, res)
	handleError(cErr, w, "create Error")

	jErr := json.NewEncoder(w).Encode(t)
	handleError(jErr, w, "jsonEncode Error")
}

func (router *DonationRouter) GetItemById(w http.ResponseWriter, r *http.Request) {
	var t models.Donation

	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	handleError(err, w, "Invalid ksuid conversion")

	res, err := t.GetItemById(repository.Rdb, collection_name, id)
	handleError(err, w, "getItemByIdError")

	jErr := json.NewEncoder(w).Encode(res)
	handleError(jErr, w, "jsonEncode Error")
}

func (router *DonationRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	list, err := t.GetAll(repository.Rdb, collection_name)
	handleError(err, w, "getAllError")

	jErr := json.NewEncoder(w).Encode(list)
	handleError(jErr, w, "jsonEncode Error")
}

func (router *DonationRouter) Update(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&t)
	jsonD, err := json.Marshal(t)
	handleError(err, w, "Serialization Error")

	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	handleError(err, w, "ksuid frombytes Error")

	res, err := t.Update(repository.Rdb, collection_name, id, jsonD)
	handleError(err, w, "Update Error")

	jErr := json.NewEncoder(w).Encode(res)
	handleError(jErr, w, "jsonEconde Error")

}

func (router *DonationRouter) Delete(w http.ResponseWriter, r *http.Request) {
	var t models.Donation
	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	handleError(err, w, "ksuid frombytes Error")

	res, err := t.Delete(repository.Rdb, collection_name, id)
	handleError(err, w, "delete Error")

	jErr := json.NewEncoder(w).Encode(res)
	handleError(jErr, w, "jsonEncode Error")
}

func RegisterDonationRoutes(prefix string, mux *http.ServeMux) {
	dr := &DonationRouter{
		mux: mux,
	}
	dr.mux.HandleFunc(POST+prefix+"/create", dr.Create)
	dr.mux.HandleFunc(POST+prefix+"/{id}", dr.GetItemById)
	dr.mux.HandleFunc(GET+prefix+"/", dr.GetAll)
	dr.mux.HandleFunc(PUT+prefix+"/update/{id}", dr.Update)
	dr.mux.HandleFunc(DELETE+prefix+"/delete/{id}", dr.Delete)
}

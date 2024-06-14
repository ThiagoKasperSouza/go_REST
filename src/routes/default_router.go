package routes

import (
	"encoding/json"
	"go_Rest/src/repository"
	"go_Rest/src/repository/models"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
)

func (router *Router) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Info
	id, kErr := ksuid.NewRandomWithTime(time.Now())
	HandleError(kErr, w, "ksUid Error")
	t.Id = id
	decoder.Decode(&t)

	res, err := json.Marshal(t)
	HandleError(err, w, "Serialization Error")

	ValidateLinkRegex(w, t)

	_, cErr := t.Create(repository.Rdb, router.collection_name, t.Id, res)
	HandleError(cErr, w, "create Error")

	jErr := json.NewEncoder(w).Encode(t)
	HandleError(jErr, w, "jsonEncode Error")
}

func (router *Router) GetItemById(w http.ResponseWriter, r *http.Request) {
	var t models.Info

	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	HandleError(err, w, "Invalid ksuid conversion")

	res, err := t.GetItemById(repository.Rdb, router.collection_name, id)
	HandleError(err, w, "getItemByIdError")

	jErr := json.NewEncoder(w).Encode(res)
	HandleError(jErr, w, "jsonEncode Error")
}

func (router *Router) GetAll(w http.ResponseWriter, r *http.Request) {
	var t models.Info
	list, err := t.GetAll(repository.Rdb, router.collection_name)
	HandleError(err, w, "getAllError")

	jErr := json.NewEncoder(w).Encode(list)
	HandleError(jErr, w, "jsonEncode Error")
}

func (router *Router) Update(w http.ResponseWriter, r *http.Request) {
	var t models.Info
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&t)
	jsonD, err := json.Marshal(t)
	HandleError(err, w, "Serialization Error")

	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	HandleError(err, w, "ksuid frombytes Error")

	res, err := t.Update(repository.Rdb, router.collection_name, id, jsonD)
	HandleError(err, w, "Update Error")

	jErr := json.NewEncoder(w).Encode(res)
	HandleError(jErr, w, "jsonEconde Error")

}

func (router *Router) Delete(w http.ResponseWriter, r *http.Request) {
	var t models.Info
	byteId := []byte(r.PathValue("id"))
	id, err := ksuid.FromBytes(byteId)
	HandleError(err, w, "ksuid frombytes Error")

	res, err := t.Delete(repository.Rdb, router.collection_name, id)
	HandleError(err, w, "delete Error")

	jErr := json.NewEncoder(w).Encode(res)
	HandleError(jErr, w, "jsonEncode Error")
}

func RegisterRoutes(prefix string, mux *http.ServeMux) {
	/*
		Se houver necessidade, é possível incluir modelos
		customizados como parametro adicional, incluindo em seguida na struct Router,
		dentro de conf.go
	*/
	dr := &Router{
		mux:             mux,
		collection_name: prefix,
	}
	dr.mux.HandleFunc(POST+prefix+"/create", dr.Create)
	dr.mux.HandleFunc(POST+prefix+"/{id}", dr.GetItemById)
	dr.mux.HandleFunc(GET+prefix+"/", dr.GetAll)
	dr.mux.HandleFunc(PUT+prefix+"/update/{id}", dr.Update)
	dr.mux.HandleFunc(DELETE+prefix+"/delete/{id}", dr.Delete)
}

package handlers

import (
	"encoding/json"
	"github.com/ur0/pokeintel-report/db"
	"github.com/ur0/pokeintel-report/models"
	"io/ioutil"
	"net/http"
)

type encountersRequest struct {
	Encounters []models.Encounter
}

func EncountersHandler(w http.ResponseWriter, r *http.Request) {
	req := encountersRequest{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	go db.StoreEncounters(req.Encounters)

	w.WriteHeader(http.StatusCreated)
}

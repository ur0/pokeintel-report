package handlers

import(
  "github.com/ur0/pokeintel-report/models"
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
)

type EncountersRequest struct {
  Encounters []models.Encounter
}

func EncountersHandler(w http.ResponseWriter, r *http.Request) {
  req := EncountersRequest{}
  body, _ := ioutil.ReadAll(r.Body)

  err := json.Unmarshal(body, &req)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
  }

  //TODO: Store
}

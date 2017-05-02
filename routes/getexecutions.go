package routes

import (
	"canary/models"
	"encoding/json"
	"net/http"
)

func GetExecutions(w http.ResponseWriter, r *http.Request) {
	var m models.Execution

	models, err := m.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(&models)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
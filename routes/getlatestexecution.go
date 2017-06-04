package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetLatestExecutions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var m models.Execution

	execution, err := m.FindLatest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(&execution)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
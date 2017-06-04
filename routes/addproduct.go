package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const ContentType = "Content-Type"
const JSON = "application/json"

func AddProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product models.Product
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	saved, err := product.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(&saved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
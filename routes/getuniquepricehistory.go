package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func GetUniquePriceHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ph models.PriceHistory

	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	pricehistory, err := ph.GetUniquePrices(productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(&pricehistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
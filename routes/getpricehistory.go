package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/husobee/vestigo"
	"net/http"
	"strconv"
)

func GetPriceHistory(w http.ResponseWriter, r *http.Request) {
	var ph models.PriceHistory

	id := vestigo.Param(r, "id")
	if id == "" {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	productId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	pricehistory, err := ph.FindByProductId(productId)
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
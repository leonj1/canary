package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var p models.Product
	products, err := p.FindByStatus("ACTIVE")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ph models.PriceHistory
	for _, product := range products {
		currentPrice, err := ph.FindLatestByProductId(product.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product.CurrentPrice = currentPrice.Price
	}

	js, err := json.Marshal(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
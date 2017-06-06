package routes

import (
	"canary/models"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func EditProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product models.Product

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

	products, err := product.FindByProductId(productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var saved *models.Product
	for _, p := range *products {
		p.Name = product.Name
		p.Url = product.Url
		p.TargetPrice = product.TargetPrice
		p.Status = product.Status
		p.Website = product.Website
		saved, err = p.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	js, err := json.Marshal(&saved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, JSON)
	w.Write(js)
}

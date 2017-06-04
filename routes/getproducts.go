package routes

import (
	"canary/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	log.Print(fmt.Sprint("Fetching active products"))
	products, err := p.FindByStatus("ACTIVE")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print(fmt.Sprintf("Number of products returned: %d", len(products)))
	js, err := json.Marshal(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print(fmt.Sprintf("Returning products as json: %s", js))
	w.Header().Set(ContentType, JSON)
	w.Write(js)
}
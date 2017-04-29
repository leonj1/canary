package services

import (
	"canary/models"
	"fmt"
	"github.com/kataras/go-errors"
)

type Amazon struct {
	Price 		string
	AlternatePrice 	string
}

// Website specific implementation on how to parse it's product page
func (a Amazon) Fetch(url string) (*models.CurrentPrice, error) {
	if url == "" {
		return nil, errors.New("Please provide a url to fetch")
	}

	// TODO implement getting the price for the Product here

	p := models.CurrentPrice{
		Price: "1.00",
		AlternatePrice: "0.99",
	}

	fmt.Println("Fetched")

	return &p, nil
}

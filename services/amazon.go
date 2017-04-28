package services

import (
	"canary/models"
	"github.com/kataras/go-errors"
)

type Amazon struct {
	Price 		string
	AlternatePrice 	string
}

// Website specific implementation on how to parse it's product page
func (a Amazon) Fetch(url string) (models.CurrentPrice, error) {
	if url == "" {
		return nil, errors.New("Please provide a url to fetch")
	}
}

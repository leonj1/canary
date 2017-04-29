package services

import "canary/models"

type ProductPage interface {
	Fetch(url string) (models.CurrentPrice, error)
}

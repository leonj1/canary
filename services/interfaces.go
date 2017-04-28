package services

import "canary/models"

type ProductPage interface {
	fetch(url string) (models.CurrentPrice, error)
}

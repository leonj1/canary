package main

import (
	"canary/models"
	"canary/services"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func sendAlert(products []models.ProductOnSale) {
	log.Printf("%d number of sales found", len(products))
}

func min(a, b float64) (float64) {
	if a < b {
		return a
	}
	return b
}

func main() {

	//var delay = flag.Int("delay", 5, "interval check between product fetches in seconds")
	var userName = flag.String("user", "", "db username")
	var password = flag.String("pass", "", "db password")
	var databaseName = flag.String("db", "", "db name")
	flag.Parse()

	// open connection to db
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", *userName, *password, *databaseName)
	models.InitDB(connectionString)

	var p models.Product
	products, err := p.FindByStatus("ACTIVE")
	if err != nil {
		panic(err.Error())
	}

	var sales []models.ProductOnSale

	for _, product := range products {
		switch website := product.Website; website {
		case "amazon":
			a := services.Amazon{Name: product.Name}
			currentPrice, err := a.Fetch(product.Url)
			if err != nil {
				panic(err.Error())
			}

			priceHistory := new(models.PriceHistory)
			priceHistory.ProductId = product.Id
			priceHistory.Price = currentPrice.Price
			priceHistory.AlternatePrice = currentPrice.AlternatePrice

			ph, err := priceHistory.Save()
			if err != nil {
				panic(err.Error())
			}

			targetPrice, err := strconv.ParseFloat(product.TargetPrice, 64)
			if err != nil {
				panic(err.Error())
			}

			currentPriceInt, err := strconv.ParseFloat(ph.Price, 64)
			if err != nil {
				panic(err.Error())
			}
			alternatePriceInt, err := strconv.ParseFloat(ph.AlternatePrice, 64)
			if err != nil {
				panic(err.Error())
			}

			if currentPriceInt < targetPrice || alternatePriceInt < targetPrice {
				sales = append(sales, models.ProductOnSale{
					Name: product.Name,
					Url: product.Url,
					Price: fmt.Sprintf("%.6f", min(currentPriceInt, alternatePriceInt)),
				})
			}
		}
	}

	if len(sales) > 0 {
		sendAlert(sales)
	}
}

package main

import (
	"canary/models"
	"canary/services"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var delay = flag.Int("delay", 5, "interval check between product fetches in seconds")
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

	for _, product := range products {
		switch website := product.Website; website {
		case "amazon":
			a := new(services.Amazon)
			currentPrice := a.fetch(product.Website)
			if err != nil {
				panic(err.Error())
			}

			priceHistory := new(models.PriceHistory)
			priceHistory.ProductId = product.Id
			priceHistory.Price = currentPrice.Price
			priceHistory.AlternatePrice = currentPrice.AlternatePrice
		}
	}
}

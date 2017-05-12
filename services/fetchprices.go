package services

import (
	"bytes"
	"canary/models"
	"fmt"
	"github.com/kataras/go-errors"
	"github.com/sourcegraph/go-ses"
	"log"
	"strconv"
)

const ACTIVE = "ACTIVE"

func FetchPrices(envelope models.Envelope) {

	var p models.Product
	products, err := p.FindByStatus(ACTIVE)
	if err != nil {
		panic(err.Error())
	}

	var execution models.Execution
	_, err = execution.Save()
	if err != nil {
		panic(err.Error())
	}

	var sales []models.ProductOnSale

	for _, product := range products {
		switch website := product.Website; website {
		case "amazon":
			a := Amazon{Name: product.Name}
			currentPrice, err := a.Fetch(product.Url)
			if err != nil {
				continue
			}

			priceHistory := new(models.PriceHistory)
			priceHistory.ProductId = product.Id
			priceHistory.Price = currentPrice.Price
			priceHistory.AlternatePrice = currentPrice.AlternatePrice

			ph, err := priceHistory.Save()
			if err != nil {
				continue
			}

			targetPrice, err := strconv.ParseFloat(product.TargetPrice, 64)
			if err != nil {
				continue
			}

			currentPriceInt, err := strconv.ParseFloat(ph.Price, 64)
			if err != nil {
				continue
			}
			//alternatePriceInt, err := strconv.ParseFloat(ph.AlternatePrice, 64)
			//if err != nil {
			//	panic(err.Error())
			//}

			//if currentPriceInt < targetPrice || alternatePriceInt < targetPrice {
			if currentPriceInt < targetPrice {
				sales = append(sales, models.ProductOnSale{
					Name: product.Name,
					Url: product.Url,
					Price: fmt.Sprintf("%.2f", min(currentPriceInt, targetPrice)),
				})
			}
		}
	}

	if len(sales) > 0 {
		_ = sendAlert(envelope, sales)
	}
}

// privates

func sendAlert(envelope models.Envelope, products []models.ProductOnSale) (error) {
	log.Printf("%d number of sales found", len(products))

	// Change the From address to a sender address that is verified in your Amazon SES account.
	from := envelope.From
	to := envelope.To
	subject := envelope.Subject

	var contents bytes.Buffer
	for _, p := range products {
		contents.WriteString(fmt.Sprintf("Current Price: %s Product: %s\n", p.Price, p.Name))
	}

	_, err := ses.EnvConfig.SendEmail(from, to, subject, contents.String())
	if err != nil {
		return errors.New(fmt.Sprintf("Error sending email: %s", err))
	}

	log.Print("Email sent with products that reached target price")

	return nil
}

func min(a, b float64) (float64) {
	if a < b {
		return a
	}
	return b
}

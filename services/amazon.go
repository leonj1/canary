package services

import (
	"canary/models"
	"fmt"
	"github.com/kataras/go-errors"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

type Amazon struct {
	Name 		string
	Price 		string
	AlternatePrice 	string
}

// Website specific implementation on how to parse it's product page
func (a Amazon) Fetch(url string) (*models.CurrentPrice, error) {
	if url == "" {
		return nil, errors.New("Please provide a url to fetch")
	}

	// TODO implement getting the price for the Product here
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// Get the current Price
	startingPoint, ok := Find(root, ById("priceblock_ourprice"))
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find current Amazon price for %s", a.Name))
	}

	price := Text(startingPoint)
	if len(price)>0 {
		price = price[1:]
	}

	// Get the lowest price from other vendors
	//moreBuyingChoices_feature_div, ok := Find(root, ById("moreBuyingChoices_feature_div"))
	//if !ok {
	//	return nil, errors.New("Unable to find moreBuyingChoices_feature_div"))
	//}
	//mbc, ok := Find(moreBuyingChoices_feature_div, ById("mbc"))
	//if !ok {
	//	return nil, errors.New("Unable to find mbc")
	//}
	//pa_mbc_on_amazon_offer, ok := Find(mbc, ByClass("pa_mbc_on_amazon_offer"))
	//if !ok {
	//	return nil, errors.New("Unable to find pa_mbc_on_amazon_offer")
	//}
	//
	//vendorPrice := Text(second)
	//if len(vendorPrice)>0 {
	//	vendorPrice = vendorPrice[1:]
	//}

	// Sample result
	p := models.CurrentPrice{
		Price: price,
		AlternatePrice: "0.00",
	}

	log.Print(fmt.Sprintf("Current price: %s for %s", a.Name, p.Price))

	return &p, nil
}

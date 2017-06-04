package services

import (
	"bytes"
	"canary/models"
	"fmt"
	"github.com/kataras/go-errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Amazon struct {
	Name 		string
	Price 		string
	AlternatePrice 	string
}

func (a Amazon) amazonPrimePrice(root *html.Node) (string, error) {
	// Get the current Price
	startingPoint, ok := Find(root, ById("priceblock_ourprice"))
	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find current Amazon price for %s", a.Name))
	}

	price := Text(startingPoint)
	if len(price)>0 {
		price = price[1:]
	}

	// TODO Find a reliable way to get prices from other vendors in Amazon's product page
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
	return price, nil
}

func (a Amazon) amazonSalePrice(root *html.Node) (string, error) {
	// Get the current Price
	startingPoint, ok := Find(root, ById("priceblock_saleprice"))
	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find current Sale price for %s", a.Name))
	}

	price := Text(startingPoint)
	if len(price)>0 {
		price = price[1:]
	}

	return price, nil
}

func (a Amazon) amazonPreOrderSalePrice1(root *html.Node) (string, error) {
	// Get the current Price
	startingPoint, ok := Find(root, ByClass("header-price"))
	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find pre-order Sale price for %s", a.Name))
	}

	price := Text(startingPoint)
	if len(price)>0 {
		price = price[1:]
	}

	return price, nil
}

func (a Amazon) amazonPreOrderSalePrice2(root *html.Node) (string, error) {
	// Get the current Price
	// TODO Iffy on using this class as looking since its so common, but seems to work for now
	startingPoint, ok := Find(root, ByClass("a-color-price"))
	if !ok {
		return "", errors.New(fmt.Sprintf("Unable to find pre-order1 Sale price for %s", a.Name))
	}

	price := Text(startingPoint)
	if len(price)>0 {
		price = price[1:]
	}

	return price, nil
}

// Website specific implementation on how to parse it's product page
func (a Amazon) Fetch(url string) (*models.CurrentPrice, error) {
	if url == "" {
		return nil, errors.New("Please provide a url to fetch")
	}

	// TODO implement getting the price for the Product here
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.89 Safari/537.36")
	log.Printf("Fetching url %s", url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 300 {
		return nil, errors.New(fmt.Sprintf("Unexpected http status code: %d", resp.StatusCode))
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err
	}

	root, err := html.Parse(bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	price, err := a.amazonPrimePrice(root)
	if err != nil {
		price, err = a.amazonSalePrice(root)
		if err != nil {
			price, err = a.amazonPreOrderSalePrice1(root)
			if err != nil {
				price, err = a.amazonPreOrderSalePrice2(root)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	// Sample result
	p := models.CurrentPrice{
		Price: price,
		AlternatePrice: "0.00",
	}

	log.Print(fmt.Sprintf("Current price: %s for %s", a.Name, p.Price))

	return &p, nil
}

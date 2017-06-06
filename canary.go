package main

import (
	"canary/models"
	"canary/routes"
	"canary/services"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {

	var userName = flag.String("user", "", "Database username")
	var password = flag.String("pass", "", "Database password")
	var databaseHost = flag.String("dbHost", "", "Database host")
	var databaseName = flag.String("db", "", "Database name")
	var databasePort = flag.Int("dbPort", 3306, "Database port")
	var serverPort = flag.String("port", "", "Web Server port")
	var emailTo = flag.String("emailTo", "", "To Email Address")
	var emailFrom = flag.String("emailFrom", "", "From Email Address")
	var emailSubject = flag.String("emailSubject", "", "Email Subject")
	flag.Parse()

	// open connection to db
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", *userName, *password, *databaseHost, *databasePort, *databaseName)
	models.InitDB(connectionString)

	router := httprouter.New()

	router.POST("/public/products", routes.AddProduct)
	router.PUT("/public/products/:id", routes.EditProduct)
	router.GET("/public/products", routes.GetProducts)
	router.GET("/public/pricehistory/:id", routes.GetPriceHistory)
	router.GET("/public/pricehistory/:id/unique", routes.GetUniquePriceHistory)
	router.GET("/public/executions", routes.GetExecutions)
	router.GET("/public/executions/latest", routes.GetLatestExecutions)

	envelope := models.Envelope{
		To: *emailTo,
		From: *emailFrom,
		Subject: *emailSubject,
	}

	// Fetch prices in a background thread
	s := func() {
		for {
			services.FetchPrices(envelope)
			log.Print("Done Fetching prices for products")
			time.Sleep(time.Hour * 1)
		}
	}
	go s()

	log.Println("Starting web server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), router))
}

package main

import (
	"canary/models"
	"canary/routes"
	"canary/services"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"time"
)

func main() {

	//var delay = flag.Int("delay", 5, "interval check between product fetches in seconds")
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

	router := vestigo.NewRouter()

	router.Post("/products", routes.AddProduct)
	router.Put("/products/:id", routes.EditProduct)
	router.Get("/products", routes.GetProducts)
	router.Get("/pricehistory/:id", routes.GetPriceHistory)
	router.Get("/executions", routes.GetExecutions)

	envelope := models.Envelope{
		To: *emailTo,
		From: *emailFrom,
		Subject: *emailSubject,
	}

	// Fetch prices in a background thread
	s := func() {
		for {
			services.FetchPrices(envelope)
			time.Sleep(time.Hour * 1)
		}
	}
	go s()

	log.Println("Starting web server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), router))
}

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
	var userName = flag.String("user", "", "db username")
	var password = flag.String("pass", "", "db password")
	var databaseName = flag.String("db", "", "db name")
	var serverPort = flag.String("port", "", "server port")
	flag.Parse()

	// open connection to db
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", *userName, *password, *databaseName)
	models.InitDB(connectionString)

	router := vestigo.NewRouter()

	router.Post("/products", routes.AddProduct)
	router.Put("/products/:id", routes.EditProduct)
	router.Get("/products", routes.GetProducts)
	router.Get("/pricehistory/:id", routes.GetPriceHistory)
	router.Get("/executions", routes.GetExecutions)

	// Fetch prices in a background thread
	s := func() {
		for {
			services.FetchPrices()
			time.Sleep(time.Hour * 1)
		}
	}
	go s()

	log.Println("Starting web server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), router))
}

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
	router.Get("/products", routes.GetProducts)
	router.Get("/pricehistory", routes.GetPriceHistory)

	go services.FetchPrices()

	log.Println("Starting web server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), router))

}

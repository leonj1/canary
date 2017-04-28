package main

import (
	"canary/models"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var interval = flag.Int("interval", 5, "interval between checks")
	var userName = flag.String("user", "", "db username")
	var password = flag.String("pass", "", "db password")
	var databaseName = flag.String("db", "", "db name")
	flag.Parse()

	// open connection to db
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", *userName, *password, *databaseName)
	models.InitDB(connectionString)


}

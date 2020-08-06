package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/health"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/http/rest"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/models"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/users"
	"log"
	"net/http"
)

func main() {

	// Connect to database
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	// Migrate the schema - figure out how to do proper migrations
	db.AutoMigrate(&models.User{})

	// init all services needed to handle requests
	healthChecker := health.NewService()
	userManagement := users.NewService(db)
	// pass all the services into a handler
	router := rest.Handler(healthChecker, userManagement)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}

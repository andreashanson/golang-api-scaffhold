package main

import (
	"log"

	"example.com/internal/database"
	"example.com/internal/handlers"
	"example.com/internal/postgres"
	"example.com/internal/server"
)

func main() {

	dbRepo := postgres.NewPostgresDB()
	// dbRepo := mongo.NewMongoDB()

	db := database.NewDatabase(dbRepo)

	apiHandlers := handlers.NewHandlers(db)

	srv := server.NewServer(apiHandlers)

	if err := srv.Start(); err != nil {
		log.Println("error", err.Error())
		panic(err)
	}

}

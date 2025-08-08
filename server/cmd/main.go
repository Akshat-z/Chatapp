package main

import (
	"log"

	"github.com/Akshat-z/Chat-app/db"
	"github.com/Akshat-z/Chat-app/internal/user"
	"github.com/Akshat-z/Chat-app/router"
)

func main() {
	dbConnection, err := db.New()
	if err != nil {
		log.Fatalf("could not initialize db connection: %s", err)
	}
	userRepo := user.NewRepository(dbConnection.GetDB())
	userServ := user.NewService(userRepo)
	userHandler := user.NewHandler(userServ)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:4000")
}

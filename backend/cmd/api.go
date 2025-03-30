package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pogle-form/backend/cmd/repos"
	"pogle-form/backend/cmd/router"
	"pogle-form/backend/utils"

	"github.com/pecet3/logger"
)

func runAPI() {
	logger.Info("Starting...")
	utils.LoadEnv()
	app := repos.NewApp()
	address := os.Getenv("ADDRESS")
	router.Run(app)

	server := &http.Server{
		Addr:    address,
		Handler: app.Srv,
	}

	logger.Info(fmt.Sprintf("Server is listening on: [%s]", address))
	log.Fatal(server.ListenAndServe())

}

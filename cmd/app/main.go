package main

import (
	"fmt"
	"net/http"
	dbConn "versla/adapter/gorm"
	"versla/app"
	"versla/config"
	"versla/router"
	lr "versla/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}

	application := app.New(logger, db)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	srv := &http.Server{
		Handler: appRouter,
	}

	logger.Info().Msgf("Starting server %v", address)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failure")
	}
}

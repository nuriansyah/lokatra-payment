package main

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/shared/logger"
	"github.com/nuriansyah/lokatra-payment/shared/server"
)

var configService *configs.Config

// @securityDefinitions.apikey LokatraAuth
// @in header
// @name Authorization
func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize config
	configService = configs.Get()

	// Set desired log level
	logger.SetLogLevel(configService)

	// Wire everything up
	httpService := InitializeServiceService()

	// Run server
	go httpService.SetupAndServe()

	gracefulShutdown := server.GetGracefulShutdown()
	gracefulShutdown.AddShutdownFunc(httpService.GracefulShutdown)

	waitCh, close := gracefulShutdown.Listen()
	defer close()
	<-waitCh
	gracefulShutdown.Shutdown()
}

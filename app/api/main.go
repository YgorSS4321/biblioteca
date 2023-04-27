package main

import (
	"fmt"
	"treinamento/app/api/config"
	"treinamento/app/api/endpoints/router"
	"treinamento/infra/postgres"
)

func main() {
	setUpPostgres()
	startServer(config.ServerHost, config.ServerPort)
}

func setUpPostgres() {
	err := postgres.SetUpCredentials(
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresHost,
		config.PostgresPort,
	)
	if err != nil {
		panic(err)
	}
}

func startServer(host string, port int) {
	server := router.Start()
	server.Debug = true
	address := fmt.Sprintf("%s:%d", host, port)
	server.Logger.Info(server.Start(address))
}

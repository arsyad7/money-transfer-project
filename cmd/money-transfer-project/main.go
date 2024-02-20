package main

import (
	"fmt"
	"money-transfer-project/internal/app"
	"money-transfer-project/internal/app/controller"
	"money-transfer-project/internal/app/infra"
	"money-transfer-project/internal/app/repo/postgres"
	"money-transfer-project/internal/app/repo/rest"
	"money-transfer-project/internal/app/service"
	"money-transfer-project/pkg/di"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	err = infra.InitTimezone()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading timezone")
	}

	err = LoadApplicationConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationPackage()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationRepository()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationService()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationController()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app.Start()
}

func LoadApplicationConfig() error {
	err := di.Provide(infra.LoadPgDatabaseCfg)
	if err != nil {
		return fmt.Errorf("LoadPgDatabaseCfg: %s", err.Error())
	}

	err = di.Provide(infra.LoadAppCfg)
	if err != nil {
		return fmt.Errorf("LoadAppCfg: %s", err.Error())
	}

	return nil
}

func LoadApplicationPackage() error {
	err := di.Provide(infra.NewEcho)
	if err != nil {
		return fmt.Errorf("NewEcho: %s", err.Error())
	}

	err = di.Provide(infra.NewDatabases)
	if err != nil {
		return fmt.Errorf("NewDatabases: %s", err.Error())
	}

	return nil
}

func LoadApplicationRepository() error {
	err := di.Provide(postgres.NewMoneyTransferRepo)
	if err != nil {
		return fmt.Errorf("NewMoneyTransferRepo: %s", err.Error())
	}

	err = di.Provide(rest.NewMockAPIRepo)
	if err != nil {
		return fmt.Errorf("NewMockAPIRepo: %s", err.Error())
	}

	return nil
}

func LoadApplicationService() error {
	err := di.Provide(service.NewMoneyTransferSvc)
	if err != nil {
		return fmt.Errorf("NewMoneyTransferService: %s", err.Error())
	}

	return nil
}

func LoadApplicationController() error {
	err := di.Provide(controller.NewMoneyTransferCtrl)
	if err != nil {
		return fmt.Errorf("NewMoneyTransferController: %s", err.Error())
	}

	return nil
}

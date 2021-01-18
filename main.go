package main

import (
	"context"
	"os"
	"time"

	log "github.com/ihsan-husaeri/employee-service/common/logger"
	"github.com/ihsan-husaeri/employee-service/config"
	"github.com/ihsan-husaeri/employee-service/config/database"
	"github.com/ihsan-husaeri/employee-service/modules/employee/handler"
	"github.com/ihsan-husaeri/employee-service/modules/employee/repository"
	"github.com/ihsan-husaeri/employee-service/modules/employee/service"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitEnv()
	mongoConf := &database.MongoDBConfig{
		Username:     config.DbUsername,
		Password:     config.DbPassword,
		DatabaseName: config.DbName,
		DatabasePort: config.DbPort,
		URLDB:        config.URLDb,
		Timeout:      10,
	}

	// Create log file
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}
	// Echo instance
	e := echo.New()

	client, err := database.NewMongoClient(mongoConf)

	if err != nil {
		log.MakeLogEntry(nil).Panic(err)
		panic(err)
	}
	db := database.NewMongoDatabase(client, mongoConf)
	log.MakeLogEntry(nil).Info("Connected to database...")
	ctx, cancel := context.WithTimeout(context.Background(), mongoConf.Timeout*time.Second)

	defer cancel()

	emplRepository := repository.NewEmployeeRepository(db)
	emplService := service.NewEmployeeService(ctx, emplRepository)
	handler.ApplyEmployeeHandler(e, emplService)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.AppPort))

}

package database

import (
	"fmt"
	config_global "norht-gowind/src/common/configs"
	model2 "norht-gowind/src/modules/customers/models"
	model1 "norht-gowind/src/modules/echo/models"
	model3 "norht-gowind/src/modules/employees/models"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	host := config_global.GetEnvVar("HOST")
	dbPort, _ := strconv.ParseUint(config_global.GetEnvVar("DB_PORT"), 10, 32)
	dbName := config_global.GetEnvVar("DB_NAME")
	dbUser := config_global.GetEnvVar("DB_USER")
	dbPass := config_global.GetEnvVar("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, dbUser, dbPass, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&model1.Echo{}, &model2.Customer{}, &model3.Employee{})
}

func CloseConnectio() {
	connection, _ := DB.DB()
	connection.Close()
}

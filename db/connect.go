package db

import (
	"fmt"
	"github.com/ngdangkietse/ndk-go-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var Error error

func ConnectDB(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresDB,
		config.PostgresPort,
		config.PostgresSslMode,
		config.TimeZone)

	Instance, Error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Error != nil {
		log.Fatal("Cannot connect to Postgres:", Error)
	}

	log.Println("Connected to db:", Instance)
}

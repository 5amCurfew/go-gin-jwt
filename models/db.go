package models

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect
func ConnectToAuthDatabase() {
	var err error

	databaseName := os.Getenv("DATABASE_NAME")
	DB, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to connect database %s", databaseName))
	} else {
		log.Infof("%s connection successful", databaseName)
	}

	DB.AutoMigrate(&User{})
}

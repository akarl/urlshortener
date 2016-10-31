package database

import (
	"urlshortener/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	DB, err = gorm.Open("sqlite3", config.DB_CONNECTION)

	return err
}

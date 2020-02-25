package models

import (
	"fmt"
	"go-pet/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func Init() {
	con := config.GetConfig()

	// db configuration
	port := con.GetString("psql.port")
	host := con.GetString("psql.host")
	database := con.GetString("psql.database")
	username := con.GetString("psql.username")
	password := con.GetString("psql.password")
	sslmode := con.GetString("psql.sslmode")
	dbconfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, username, database, password, sslmode)
	db, err = gorm.Open("postgres", dbconfig)
	if err != nil {
		log.Fatalf("failed to connect to database. Err %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

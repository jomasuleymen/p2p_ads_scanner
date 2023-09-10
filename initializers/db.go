package initializers

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {

	if db != nil {
		return db
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dbDSN := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", dbHost, dbUser, dbName, dbPassword, dbPort)
	tempDB, err := gorm.Open(
		postgres.New(postgres.Config{DSN: dbDSN, PreferSimpleProtocol: true}),
		&gorm.Config{
			CreateBatchSize: 900,
		},
	)

	if err != nil {
		log.Println("Couldn't connect to database:", err.Error())
		os.Exit(1)
	}

	sqlDb, _ := tempDB.DB()

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(15)

	db = tempDB
	db.Config.Logger = myLogger

	return db
}

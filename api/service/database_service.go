package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
)

type Database struct {
	DB            *gorm.DB
	IsInitialized bool
}

type DatabaseService interface {
	Initialize() error
	Migrate() error
}

type databaseService struct {

}

func NewDatabaseService() DatabaseService {
	return databaseService{}
}

var database Database

func (d databaseService) Initialize() error {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DbUser, DbPassword, DbPort, DbHost, DbName := os.Getenv("mysql_user"), os.Getenv("mysql_password"), os.Getenv("mysql_port"), os.Getenv("mysql_host"), os.Getenv("mysql_database")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	database = Database{DB:db, IsInitialized:true}

	return nil
}

func (d databaseService) Migrate() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	if database.IsInitialized {
		dbConn := database.DB

		fmt.Println("Starting migration")
		result, err := migrate.Exec(dbConn.DB(), "mysql", migrations, migrate.Up)

		if err != nil {
			log.Fatalf("Migration failed: %s", err)
		}

		log.Printf("Applied %d migrations successfully", result)

		return nil
	} else {
		return errors.New("database not initialized")
	}
}
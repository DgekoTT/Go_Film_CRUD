package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var TestDB *gorm.DB

func InitTestDB(string2 string) (*gorm.DB, error) {
	var err error
	dsn := string2
	TestDB, err = gorm.Open(postgres.Open(dsn), Config)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Printf("Connected to database")
	return TestDB, nil
}

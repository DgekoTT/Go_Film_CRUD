package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var Config = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	NowFunc: func() time.Time {
		return time.Now().Round(time.Millisecond)
	},
}

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), Config)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Printf("Connected to database")
}

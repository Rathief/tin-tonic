package config

import (
	"log"
	"tin-tonic/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// dsn := os.Getenv("DB_ADDR")
	dsn := "host=localhost user=postgres password=admin dbname=tin-tonic port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&entity.Store{}); err != nil {
		log.Panic(err)
	}
	return db
}

package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func Init() {
	dsn := "host=localhost user=postgres password=qwaszx dbname=gorm_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}
	
	_ = db
}

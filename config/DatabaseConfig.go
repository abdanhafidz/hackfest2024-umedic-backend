package config

import (
	"github.com/quzuu-be/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

func init() {
	dsn := "host=localhost user=postgres password=558585 dbname=umedic-db port=8899 sslmode=disable TimeZone=Asia/Jakarta"
	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic(&Err)
	}
	if DB != nil {
		DB.AutoMigrate(&models.Dokter{})
		DB.AutoMigrate(&models.Keluhan{})
	}
}

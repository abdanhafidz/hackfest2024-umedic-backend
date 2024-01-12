package models

import (
	"time"
)

type Keluhan struct {
	ID              uint      `gorm:"column:id_keluhan;primaryKey" json:"id_keluhan"`
	KategoriKeluhan string    `gorm:"column:kategori_keluhan" json:"kategori_keluhan"`
	UserID          uint      `gorm:"column:id_user" json:"id_user"`
	Keluhan         string    `gorm:"column:keluhan" json:"keluhan"`
	DokterHandlerID uint      `gorm:"column:id_dokter_handler" json:"id_dokter_handler"`
	Status          string    `gorm:"column:status;type:ENUM('baru','penting','serius','pengawasan','ditangani');default:'baru'" json:"status"`
	CreatedAt       time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (Keluhan) TableName() string {
	return "data_keluhan"
}

type Dokter struct {
	ID             uint      `gorm:"column:id_dokter;primaryKey" json:"id_dokter"`
	Name           string    `gorm:"column:name" json:"name"`
	Occupation     string    `gorm:"column:occupation" json:"occupation"`
	WorkExperience string    `gorm:"column:workexperience" json:"workexperience"`
	PhoneNumber    uint      `gorm:"column:phonenumber" json:"phonenumber"`
	CreatedAt      time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (Dokter) TableName() string {
	return "data_dokter"
}

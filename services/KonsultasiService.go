// handlers/auth_handlers.go

package services

import (
	"time"

	"github.com/quzuu-be/config"
	"github.com/quzuu-be/models"
)

// LoginHandler handles user login
func KonsultasiService(dataReq *models.KonsultasiRequest) (err any, kelid uint) {
	var db = config.DB
	konsultasi := models.Keluhan{
		KategoriKeluhan: dataReq.Kategori_keluhan,
		UserID:          dataReq.Id_user,
		Keluhan:         dataReq.Keluhan,
		DokterHandlerID: dataReq.Id_dokter,
		Status:          "baru",
		CreatedAt:       time.Now(),
	}
	result := db.Create(&konsultasi)
	return result.Error, konsultasi.ID

}

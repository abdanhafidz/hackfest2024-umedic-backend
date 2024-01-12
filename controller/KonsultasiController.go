package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func KonsultasiController(c *gin.Context) {
	var dataReq models.KonsultasiRequest
	c.Bind(&dataReq)
	err, konsultasiID := services.KonsultasiService(&dataReq)
	if err != nil {
		panic(err)
	}
	data := strconv.FormatUint(uint64(konsultasiID), 10)
	respon := "Saat ini umedic sedang dalam tahap pengembangan dan anda akan mendapatkan tanggapan awal lewat email / whatsapp langsung dari dokter. Terima kasih sudah mengirim keluhan!"
	c.JSON(200, gin.H{
		"data":           gin.H{"id_konsultasi": data},
		"message":        "create konsultasi succeed!",
		"tanggapan_awal": respon,
	})
}

package models

type KonsultasiRequest struct {
	Id_user          uint   `form:"id_user"`
	Kategori_keluhan string `form:"kategori_keluhan"`
	Keluhan          string `form:"keluhan"`
	Id_dokter        uint   `form:"id_dokter"`
}

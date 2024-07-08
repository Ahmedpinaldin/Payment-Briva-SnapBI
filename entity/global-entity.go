package entity

type ParamsApprovalCabang struct {
	Wil_Rmd string `form:"wil_rmd"`
}
type ResponseCabang struct {
	Id_Cabang   *string `json:"id_cabang"`
	Inisial_Cab *string `json:"inisial_cab"`
	Nama        *string `json:"nama"`
	Alamat      *string `json:"alamat"`
}
type ResponseUnit struct {
	Kode_Unit    *string `json:"kode_unit"`
	Nama_Unit    *string `json:"nama_unit"`
	Inisial_Unit *string `json:"inisial_unit"`
}
type ResponseCabangApproval struct {
	Inisial_Cab *string `json:"inisial_cab"`
	Nama        *string `json:"nama"`
	Status      *string `json:"status"`
}

type ResponseWilayah struct {
	Wilayah_Rmd *string `json:"wilayah_rmd"`
	Keterangan  *string `json:"keterangan"`
}
type ResponseCaraPenanganan struct {
	Cara_Penanganan *string `json:"cara_penanganan"`
}
type ResponseRencananPenyelesaian struct {
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
}
type ResponsePic struct {
	Id_Pic      *string `json:"id_pic"`
	Inisial_Pic *string `json:"inisial_pic"`
	Nama        *string `json:"nama"`
}
type ResponseDetailPic struct {
	Id_Pic      *string `json:"id_pic"`
	Inisial_Pic *string `json:"inisial_pic"`
	Nama        *string `json:"nama"`
	Jabatan     *string `json:"jabatan"`
	No_Hp       *string `json:"no_hp"`
}

type ResponsePenyelesaian struct {
	Id_Status           *string `json:"id_status"`
	Status_Penyelesaian *string `json:"status_penyelesaian"`
}

type ResponseHasilPenyelesaian struct {
	Id_Hasil           *string `json:"id_hasil"`
	Id_Status          *string `json:"id_status"`
	Hasil_Penyelesaian *string `json:"hasil_penyelesaian"`
}

package entity

import "time"

type LoginDTO struct {
	Username string `validate:"required,min=6,max=50"`
	Password string `validate:"required,min=6,max=15"`
}

type ResponseLogin struct {
	Login []Datas `json:"login"`
}

type Datas struct {
	Data []DataLogin `json:"data"`
}
type DataLogin struct {
	USERNAME    string `json:"USERNAME"`
	Password    string `json:"Password"`
	Perusahaan  string `json:"Perusahaan"`
	Kode_cabang string `json:"Kode_cabang"`
	Kode_unit   string `json:"Kode_unit"`
	Posisi_nama string `json:"Posisi_nama"`
}

type Param struct {
	Name  string
	Value interface{}
}
type LoginResponse struct {
	Login []LoginData `json:"login"`
}
type LoginData struct {
	Data     []Data `json:"data"`
	Response string `json:"response"`
	Message  string `json:"message"`
}

type Data struct {
	IDSDM           string `json:"idsdm"`
	NIK             string `json:"nik"`
	Perusahaan      string `json:"perusahaan"`
	NomorInduk      string `json:"nomor_induk"`
	Username        string `json:"username"`
	Nama            string `json:"nama"`
	Email           string `json:"email"`
	PosisiNama      string `json:"posisi_nama"`
	PosisiSSO       string `json:"posisi_sso"`
	PosisiSingkatan string `json:"posisi_singkatan"`
	LokasiKerja     string `json:"lokasi_kerja"`
	KodeCabang      string `json:"kode_cabang"`
	Cabang          string `json:"cabang"`
	Unit            string `json:"unit"`
	KodeUnit        string `json:"kode_unit"`
	Mekaar          string `json:"mekaar"`
	KodeMekaar      string `json:"kode_mekaar"`
	KodeLokasi      string `json:"kode_lokasi"`
	Foto            string `json:"foto"`
}

type Response struct {
	User_ID      *int64     `json:"user_id"`
	Username     *string    `json:"username"`
	Name         *string    `json:"name"`
	Password     *string    `json:"password"`
	Status       *bool      `json:"status"`
	Id_Level     *int64     `json:"id_level"`
	Inisial_Cab  *string    `json:"inisial_cab"`
	Kode_Unit    *string    `json:"kode_unit"`
	Kode_Cabang  *string    `json:"kode_cabang"`
	Add_by       *string    `json:"add_by"`
	Date_add     *time.Time `json:"date_add"`
	Time_Login   *time.Time `json:"time_login"`
	Group_Level  *string    `json:"group_level"`
	Nama_Unit    *string    `json:"nama_unit"`
	Map_Unit     *string    `json:"map_unit"`
	Level        *string    `json:"level"`
	Wilayah_RMD  *string    `json:"wilayah_rmd"`
	Wilayah      *string    `json:"wilayah"`
	Bucket       *string    `json:"bucket"`
	Unit         *string    `json:"unit"`
	Cabang       *string    `json:"cabang"`
	Alternate    *string    `json:"alternate"`
	Foto         *string    `json:"foto"`
	Token        *string    `json:"token"`
	Lokasi_Kerja *string    `json:"lokasi_kerja"`
	Nomor_Induk  *string    `json:"nomor_induk"`
}
type GetProfile struct {
	Unit string `valid:"unit"`
}

package entity

type ResponseProfile struct {
	Nama_Unit      *string `json:"nama_unit"`
	Kode_Unit      *string `json:"kode_unit"`
	Inisial        *string `json:"inisial"`
	Alamat_Unit    *string `json:"alamat_unit"`
	Inisial_Cabang *string `json:"inisial_cabang"`
	Nama_Cabang    *string `json:"nama_cabang"`
	Alamat_Cabang  *string `json:"alamat_cabang"`
} 
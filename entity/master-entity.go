package entity

type ParamsNasabah struct {
	PAGE        *string `json:"PAGE"`
	LIMIT       *string `json:"LIMIT"`
	TYPE        *string `json:"TYPE"`
	GROUP_ROLE  *string `json:"GROUP_ROLE"`
	ROLE        *string `json:"ROLE"`
	WIL_RMD     *string `json:"WIL_RMD"`
	INISIAL_CAB *string `json:"INISIAL_CAB"`
	KODE_UNIT   *string `json:"KODE_UNIT"`
	SEARCH      *string `json:"SEARCH"`
	BUCKET      *string `json:"BUCKET"`
}

type ResponseNasabah struct {
	ID              *string `json:"id"`
	Nama_Unit       *string `json:"nama_unit"`
	Cabang          *string `json:"cabang"`
	No_Rekening     *string `json:"no_rekening"`
	Kolektibilitas  *string `json:"kolektibilitas"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	Tipe_Kredit     *string `json:"tipe_kredit"`
	Ft              *string `json:"ft"`
	Angs_Ke         *string `json:"angs_ke"`
	Jml_Pinjaman    *string `json:"jml_pinjaman"`
	Saldo_Nominatif *string `json:"saldo_nominatif"`
}
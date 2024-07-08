package response

type BulanResponse struct {
	Id                   int64   `json:"id"`
	Kode_Unit            string  `json:"kode_unit"`
	No_Rekening          string  `json:"no_rekening"`
	Nama_Nasabah         string  `json:"nama_nasabah"`
	Tipe_Kredit          string  `json:"tipe_kredit"`
	Kolektibilitas       string  `json:"kolektibilitas"`
	Ft                   int64   `json:"ft"`
	Angs_Ke              int64   `json:"angs_ke"`
	Saldo_Nominatif      float64 `json:"saldo_nominatif"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
}

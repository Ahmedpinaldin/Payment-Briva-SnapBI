package response

type PlanBulanWo struct {
	Id                   int64   `json:"id"`
	Kode_Unit            string  `json:"kode_unit"`
	No_Rekening          string  `json:"no_rekening"`
	Nama_Nasabah         string  `json:"nama_nasabah"`
	Tgl_Wo               string  `json:"tgl_wo"`
	Jml_wo               float64 `json:"jml_wo"`
	Saldo_Nominatif      float64 `json:"saldo_nominatif"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
}

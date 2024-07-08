package response

type PlanBulanById struct {
	No_Rekening       string  `json:"no_rekening"`
	Kode_Unit         string  `json:"kode_unit"`
	Inisial_Cab       string  `json:"inisial_cab"`
	Nama_Unit         string  `json:"nama_unit"`
	Nama_Nasabah      string  `json:"nama_nasabah"`
	Tipe_Kredit       string  `json:"tipe_kredit"`
	Tgl_Realisasi     string  `json:"tgl_realisasi"`
	Tgl_Jatuh_Tempo   string  `json: "tgl_jatuh_tempo"`
	Bln_Jatuh_Tempo   string  `json: "bln_jatuh_tempo"`
	Tgl_Angs          string  `json: "tgl_angs"`
	Jml_Pinjaman      float64 `json: "jml_pinjaman"`
	Kolektibilitas    string  `json:"kolektibilitas"`
	Saldo_Nominatif   float64 `json:"saldo_nominatif"`
	Ft                int64   `json:"ft"`
	Jenis_Usaha       string  `json:"jenis_usaha"`
	Kondisi_Usaha     string  `json:"kondisi_usaha"`
	Sp_1              string  `json:"sp_1"`
	Sp_2              string  `json:"sp_2"`
	Sp_3              string  `json:"sp_3"`
	Obyek_Agunan      string  `json:"obyek_agunan"`
	Dokumen_Agunan    string  `json:"dokumen_agunan"`
	Pengikatan        string  `json:"pengikatan"`
	Jenis_Pengikatan  string  `json:"jenis_pengikatan"`
	Status_Dokumen    string  `json:"status_dokumen"`
	Nilai_Saat_Ini    string  `json:"nilai_saat_ini"`
	Mudah_Jual        string  `json:"mudah_jual"`
	Keberadaan_Agunan string  `json:"keberadaan_agunan"`
	Sengketa          string  `json:"sengketa"`
	S_Diputuskan_Kpp  string  `json:"s_diputuskan_kpp"`
	Status_Lelang     string  `json:"status_lelang"`
	Lelang_Ke         string  `json:"lelang_ke"`
	Permasalahan      string  `json:"permasalahan"`
	Kategori_Recovery *int    `json:"kategori_recovery"`
	Ket_Recovery      *string `json:"ket_recovery"`
}

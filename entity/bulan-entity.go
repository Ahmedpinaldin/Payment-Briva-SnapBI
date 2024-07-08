package entity

type PlanBulan struct {
	Id                     int64       `json:"id"`
	Tgl_Plan               string      `json:"tgl_plan" `
	Bln_Plan               int64       `json:"bln_plan"`
	Thn_Plan               int64       `json:"thn_plan"`
	No_Rekening            string      `json:"no_rekening"`
	Nama_Nasabah           *string     `json:"nama_nasabah"`
	Kode_Unit              string      `json:"kode_unit"`
	Inisial_Cab            string      `json:"inisial_cab"`
	Jenis_Usaha            string      `json:"jenis_usaha"`
	Kondisi_Usaha          string      `json:"kondisi_usaha"`
	Sp_1                   string      `json:"sp_1"`
	Sp_2                   string      `json:"sp_2"`
	Sp_3                   string      `json:"sp_3"`
	Obyek_Agunan           string      `json:"obyek_agunan"`
	Dokumen_Agunan         string      `json:"dokumen_agunan"`
	Pengikatan             string      `json:"pengikatan"`
	Jenis_Pengikatan       string      `json:"jenis_pengikatan"`
	Status_Dokumen         string      `json:"status_dokumen"`
	Nilai_Saat_Ini         string      `json:"nilai_saat_ini"`
	Mudah_Jual             string      `json:"mudah_jual"`
	Keberadaan_Agunan      string      `json:"keberadaan_agunan"`
	Sengketa               string      `json:"sengketa"`
	S_Diputuskan_Kpp       string      `json:"s_diputuskan_kpp"`
	Status_Lelang          string      `json:"status_lelang"`
	Lelang_Ke              string      `json:"lelang_ke"`
	Permasalahan           string      `json:"permasalahan"`
	Cara_Penanganan        string      `json:"cara_penanganan"`
	Rencana_Penyelesaian   *string     `json:"rencana_penyelesaian"`
	Input_Pokok            float32     `json:"input_pokok"`
	Nominal_Lunas_Bertahap float32     `json:"nominal_lunas_bertahap"`
	Tgl_Jb                 string      `json:"tgl_jb"`
	Ket_Rinci              string      `json:"ket_rinci"`
	Id_Pic                 string      `json:"id_pic"`
	Kategori_Recovery      interface{} `json:"kategori_recovery"`
	Ket_Recovery           *string     `json:"ket_recovery"`
}

type PlanBulanWo struct {
	Id                     int64       `json:"id"`
	Tgl_Plan               string      `json:"tgl_plan" `
	Bln_Plan               int64       `json:"bln_plan"`
	Thn_Plan               int64       `json:"thn_plan"`
	No_Rekening            string      `json:"no_rekening"`
	Nama_Nasabah           *string     `json:"nama_nasabah"`
	Kode_Unit              string      `json:"kode_unit"`
	Inisial_Cab            string      `json:"inisial_cab"`
	Lunas_Blm_Lunas        string      `json:"lunas_blm_lunas"`
	Nilai_Jamkrindo        string      `json:"nilai_jamkrindo"`
	Jaminan_Ada_Tidak      string      `json:"jaminan_ada_tidak"`
	Sp_1                   string      `json:"sp_1"`
	Sp_2                   string      `json:"sp_2"`
	Sp_3                   string      `json:"sp_3"`
	Obyek_Agunan           string      `json:"obyek_agunan"`
	Dokumen_Agunan         string      `json:"dokumen_agunan"`
	Pengikatan             string      `json:"pengikatan"`
	Jenis_Pengikatan       string      `json:"jenis_pengikatan"`
	Status_Dokumen         string      `json:"status_dokumen"`
	Nilai_Saat_Ini         string      `json:"nilai_saat_ini"`
	Mudah_Jual             string      `json:"mudah_jual"`
	Keberadaan_Agunan      string      `json:"keberadaan_agunan"`
	Sengketa               string      `json:"sengketa"`
	S_Diputuskan_Kpp       string      `json:"s_diputuskan_kpp"`
	Status_Lelang          string      `json:"status_lelang"`
	Lelang_Ke              string      `json:"lelang_ke"`
	Permasalahan           string      `json:"permasalahan"`
	Cara_Penanganan        string      `json:"cara_penanganan"`
	Rencana_Penyelesaian   *string     `json:"rencana_penyelesaian"`
	Input_Pokok            float32     `json:"input_pokok"`
	Nominal_Lunas_Bertahap float32     `json:"nominal_lunas_bertahap"`
	Tgl_Jb                 string      `json:"tgl_jb"`
	Ket_Rinci              string      `json:"ket_rinci"`
	Id_Pic                 string      `json:"id_pic"`
	Nominal_Pembayaran     string      `json:"nominal_pembayaran"`
	Kategori_Recovery      interface{} `json:"kategori_recovery"`
	Ket_Recovery           *string     `json:"ket_recovery"`
}

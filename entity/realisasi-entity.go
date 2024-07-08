package entity

type ParamsTransaksi struct {
	Page           string `form:"page"`
	Limit          string `form:"limit"`
	Wilayah        string `form:"wilayah"`
	Wilayah_Cabang string `form:"wilayah_cabang"`
	UnitKode       string `form:"unit_kode"`
	Date_Awal      string `form:"date_awal"`
	Date_Akhir     string `form:"date_akhir"`
	Kode           string `form:"kode"`
}

type ResponseTransaksiResumeWidget struct {
	Total_Noa       *string `json:"TOTAL_NOA"`
	Transaksi_Pokok *string `json:"TRANSAKSI_POKOK"`
	Transaksi_Bunga *string `json:"TRANSAKSI_BUNGA"`
	Transaksi_Denda *string `json:"TRANSAKSI_DENDA"`
	Total_Transaksi *string `json:"TOTAL_TRANSAKSI"`
}
type ResponseTransaksiTable struct {
	Cabang          *string `json:"CABANG"`
	Unit            *string `json:"UNIT"`
	Wil             *string `json:"WIL"`
	Total_Noa       *string `json:"TOTAL_NOA"`
	Transaksi_Pokok *string `json:"TRANSAKSI_POKOK"`
	Transaksi_Bunga *string `json:"TRANSAKSI_BUNGA"`
	Transaksi_Denda *string `json:"TRANSAKSI_DENDA"`
	Total_Transaksi *string `json:"TOTAL_TRANSAKSI"`
}

type ResponseDetailTransaksi struct {
	No_Rekening     *string `json:"NO_REKENING"`
	Nama_Nasabah    *string `json:"NAMA_NASABAH"`
	Cabang          *string `json:"CABANG"`
	Unit            *string `json:"UNIT"`
	Transaksi_Pokok *string `json:"TRANSAKSI_POKOK"`
	Transaksi_Bunga *string `json:"TRANSAKSI_BUNGA"`
	Transaksi_Denda *string `json:"TRANSAKSI_DENDA"`
}
type ResponseDetailTransaksiById struct {
	Kode_Unit       *string `json:"KODE_UNIT"`
	TglTrans        *string `json:"TGLTRANS"`
	No_Rekening     *string `json:"NO_REKENING"`
	Nama_Nasabah    *string `json:"NAMA_NASABAH"`
	Kode_Trans      *string `json:"KODE_TRANS"`
	Transaksi_Pokok *string `json:"TRANSAKSI_POKOK"`
	Transaksi_Bunga *string `json:"TRANSAKSI_BUNGA"`
	Transaksi_Denda *string `json:"TRANSAKSI_DENDA"`
	Unit            *string `json:"UNIT"`
	Id_Nasabah      *string `json:"ID_NASABAH"`
	Tipe_Kredit     *string `json:"TIPE_KREDIT"`
	Ft              *string `json:"FT"`
	Angs_Ke         *string `json:"ANGS_KE"`
	Tgl_Angsuran    *string `json:"TGL_ANGSURAN"`
	Tgl_Realisasi   *string `json:"TGL_REALISASI"`
	Tgl_Jth_Tempo   *string `json:"TGL_JTH_TEMPO"`
	Kol             *string `json:"KOL"`
	Angs_Total      *string `json:"ANGS_TOTAL"`
	Jml_Pinjaman    *string `json:"JML_PINJAMAN"`
	Os              *string `json:"OS"`
	Rek_Dca         *string `json:"REK_DCA"`
	Saldo_Dca       *string `json:"SALDO_DCA"`
}

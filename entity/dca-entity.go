package entity

type ParamListMppd struct {
	No_Memo string `form:"no-memo"`
}
type ParamListLmpd struct {
	From_Date string `form:"from-date"`
	To_Date   string `form:"to-date"`
}
type ResponseMonitoringDca struct {
	JMLH int64  `json:"JMLH"`
	Data string `json:"DATA"`
}
type DataDCA struct {
	Janji_Pertemuan *string `json:"janji_pemenuhan"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	No_Memo         *string `json:"no_memo"`
	Angs_Ke         *string `json:"angs_ke"`
	Angsuran_Total  *string `json:"angsuran_total"`
	Saldo_Dca       *string `json:"saldo_dca"`
	Alasan          *string `json:"alasan"`
	Debet_DCA       *string `json:"debet_dca"`
}
type DataLMPD struct {
	No_Memo         *string `json:"no_memo"`
	Tgl_Memo        *string `json:"tgl_memo"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	No_Rekening     *string `json:"no_rekening"`
	Angsuran_Total  *string `json:"angsuran_total"`
	Angs_Ke         *string `json:"angs_ke"`
	Saldo_Dca       *string `json:"saldo_dca"`
	Janji_Pertemuan *string `json:"janji_pertemuan"`
	Debet_DCA       *string `json:"debet_dca"`
	Tgl_Trans       *string `json:"tgl_trans"`
	Setoran         *string `json:"setoran"`
}
type DataMPPD struct {
	Saldo_Dca       *string `json:"saldo_dca"`
	Debet_DCA       *string `json:"debet_dca"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	Alasan_Debet    *string `json:"alasan_Debet"`
	Tipe_Kredit     *string `json:"tipe_kredit"`
	No_Memo         *string `json:"no_memo"`
	Tgl_Memo        *string `json:"tgl_memo"`
	Nomer_Rek       *string `json:"nomer_rek"`
	Janji_Pemenuhan *string `json:"janji_pemenuhan"`
	Angs_Ke         *string `json:"angs_ke"`
	Angsuran        *string `json:"angsuran"`
	Id              *string `json:"id"`
}
type InserMPPD struct {
	Angsuran       string `valid:"required"`
	Debet_DCA      string `valid:"required"`
	Sisa_DCA       string `valid:"required"`
	Kode_Unit      string `valid:"required"`
	Kode_Cabang    string `valid:"required"`
	Bulan          string `valid:"required"`
	Tahun          string `valid:"required"`
	DateEntry_Time string `valid:"required"`
	Date_Entry     string `valid:"required"`
	Saldo_DCA      string `valid:"required"`
}
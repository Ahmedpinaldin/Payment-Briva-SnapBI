package entity

type ParamDeskCall struct {
	No_Rekening string `form:"no-rekening"`
}
type ResponseListDeskCall struct {
	Kode_Unit       *string `json:"kode_unit"`
	Inisial_Cab     *string `json:"inisial_cab"`
	Tgl_DeskCall    *string `json:"tgl_deskcall"`
	Tgl_Jb          *string `json:"tgl_jb"`
	Status_Telpon   *string `json:"stats_telp"`
	No_Hp           *string `json:"no_hp"`
	Keterangan      *string `json:"ket"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	No_Rekening     *string `json:"no_rekening"`
	Tgl_Jatuh_Tempo *string `json:"tgl_jatuh_tempo"`
	Tgl_Angsuran    *string `json:"tgl_angsuran"`
	Angsuran_Ke     *string `json:"angs_ke"`
	Tipe_Kredit     *string `json:"tipe_kredit"`
	Angsuran_Total  *string `json:"angsuran_total"`
	Nama_Karyawan   *string `json:"nama_karyawan"`
	Tgl_Realisasi   *string `json:"tgl_realisasi"`
}
type ResponseDetailDeskCall struct {
	Nama_Nasabah   *string `json:"nama_nasabah"`
	No_Rekening    *string `json:"no_rekening"`
	Jml_Pimjaman   *string `json:"jml_pinjaman"`
	Angsuran_Total *string `json:"angsuran_total"`
	Angsuran_Ke    *string `json:"angs_ke"`
	No_Hp          *string `json:"no_hp"`
	Tipe_Kredit    *string `json:"tipe_kredit"`
	Tgl_Realisasi  *string `json:"tgl_realisasi"`
	Nama_Karyawan  *string `json:"nama_karyawan"`
}
type ResponseRiwayatDeskCall struct {
	Tgl_DeskCall *string `json:"tgl_deskcall"`
	Keterangan   *string `json:"ket"`
}
type ResponseListNasabahDeskCall struct {
	Nama_Nasabah         *string `json:"NAMA_NASABAH"`
	No_Rekening          *string `json:"NO_REKENING"`
	Dpd_Clos             *string `json:"dpd_clos"`
	Angsuran             *string `json:"angsuran"`
	Dca                  *string `json:"dca"`
	Plafond              *string `json:"plafond"`
	OutStanding          *string `json:"OUTSTANDING"`
	Angsuran_Ke          *string `json:"angsuran_ke"`
	No_Hp                *string `json:"no_hp"`
	Temp_Tgl_Jatuh_Tempo *string `json:"temp_TGL_JTH_TEMPO"`
	Nama_Aom             *string `json:"nama_aom"`
	Plan_Penanganan      *string `json:"plan_penanganan"`
	Tgl_Jb               *string `json:"tgl_jb"`
	Tgl_Angsuran         *string `json:"tgl_angsuran"`
	Tgl_Jatuh_Tempo      *string `json:"TGL_JTH_TEMPO"`
}

type InsertDeskCall struct {
	No_Rekening  string `valid:"required"`
	Status_Telp  string `valid:"required"`
	Keterangan   string `valid:"required"`
	Tgl_DeskCall string `valid:"required"`
}

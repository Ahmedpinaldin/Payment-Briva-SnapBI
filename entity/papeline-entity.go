package entity

type ResponseBucketBulan struct {
	Inisial_cab      *string `json:"inisial_cab"`
	Nama_Cabang      *string `json:"nama_cabang"`
	Noa_Coll         *string `json:"noa_coll"`
	Coll             *string `json:"coll"`
	Noa_3r_Baru      *string `json:"noa_3r_baru"`
	N_3r_Baru        *string `json:"_3r_baru"`
	Noa_3r           *string `json:"noa_3r"`
	N_3r             *string `json:"_3r"`
	Noa_Po_Sekaligus *string `json:"noa_po_sekaligus"`
	Po_Sekaligus     *string `json:"po_sekaligus"`
	Noa_Po_Bertahap  *string `json:"noa_po_bertahap"`
	Po_Bertahap      *string `json:"po_bertahap"`
	Noa_Lelang       *string `json:"noa_lelang"`
	Lelang           *string `json:"lelang"`
	Noa_Back_To_0    *string `json:"noa_back_to_0"`
	Back_To_0        *string `json:"back_to_0"`
	Noa_Back_1_30    *string `json:"noa_back_1_30"`
	Back_1_30        *string `json:"back_1_30"`
	Noa_Back_31_60   *string `json:"noa_back_31_60"`
	Back_31_60       *string `json:"back_31_60"`
	Noa_Rembes_Npl   *string `json:"noa_rembes_npl"`
	Rembes_Npl       *string `json:"rembes_npl"`
}
type ParamsStatus struct {
	Wilayah_Cabang *string `form:"wilayah_cabang"`
}
type ResponseStatus struct {
	Bulan *string `json:"approve_bulan"`
	Hari  *string `json:"approve_hari"`
}
type CBResumeTotal struct {
	Kol         string `form:"kol"`
	Inisial_Cab string `form:"inisial_cab"`
}
type ParmsListBucketBulan struct {
	Tahun   string `form:"tahun"`
	Bulan   string `form:"bulan"`
	Wilayah string `form:"wilayah"`
}

type ParamsPenyelesaian struct {
	Data []DataPenyelesaian `valid:"required"`
}
type DataPenyelesaian struct {
	NoRekening        string `form:"no_rekening"`
	NoRekeningWo      string `form:"no_rekening_wo"`
	NamaNasabah       string `form:"nama_nasabah"`
	KodeUnit          string `form:"kode_unit"`
	InisialCab        string `form:"inisial_cab"`
	CaraPenanganan    string `form:"cara_penanganan"`
	RencanaPenanganan string `form:"rencana_penyelesaian"`
	KetPlan           string `form:"ket_rinci_plan"`
	Status            string `form:"status_penyelesaian"`
	Hasil             string `form:"hasil_penyelesaian"`
	Ket               string `form:"ket_rinci"`
	NominalPembayaran string `form:"nominal_pembayaran"`
	Os                string `form:"os_penyelesaian"`
	TglPlan           string `form:"tgl_plan"`
	TglJb             string `form:"tgl_jb"`
	IdPic             string `form:"id_pic"`
}
type ParamsSubmit struct {
	Data []DataSubmit `valid:"required"`
}
type ParamsKirim struct {
	Tipe         *string `form:"tipe"`
	Inisial_Cab  *string `form:"inisial_cab"`
	Status       *string `form:"status"`
	User_Approve *string `form:"user_approve"`
	User_Level   *string `form:"user_level"`
	Komen        *string `form:"komen"`
}
type DataSubmit struct {
	NoRekening        string `form:"no_rekening"`
	NoRekeningWo      string `form:"no_rekening_wo"`
	NamaNasabah       string `form:"nama_nasabah"`
	KodeUnit          string `form:"kode_unit"`
	InisialCab        string `form:"inisial_cab"`
	Kolektibilitas    string `form:"kolektibilitas"`
	Ft                string `form:"ft"`
	SaldoNominatif    string `form:"saldo_nominatif"`
	NominalPembayaran string `form:"nominal_pembayaran"`
	OsPenyelesaian    string `form:"os_penyelesaian"`
	InputPokok        string `form:"input_pokok"`
	NominalLunas      string `form:"nominal_lunas_bertahap"`
	CaraPenanganan    string `form:"cara_penanganan"`
	RencanaPenanganan string `form:"rencana_penyelesaian"`
	TglAngsuran       string `form:"tgl_angsuran"`
	KetRinci          string `form:"ket_rinci"`
	IdPic             string `form:"id_pic"`
}
type ParamsTable struct {
	Page        *string `form:"page"`
	Limit       *string `form:"limit"`
	Wilayah     *string `form:"wilayah"`
	Inisial_Cab *string `form:"cabang"`
	Search      *string `form:"search"`
	Id          *string `form:"id"`
	No_Rekening *string `form:"no_rekening"`
}
type ParamsPapelineBulanResume struct {
	Page           *string `form:"page"`
	Limit          *string `form:"limit"`
	Wilayah        *string `form:"wilayah"`
	Wilayah_Cabang *string `form:"wilayah_cabang"`
}
type ParamsDetailPlan struct {
	Page           *string `form:"page"`
	Limit          *string `form:"limit"`
	Wilayah_Cabang *string `form:"wilayah_cabang"`
	Wilayah        *string `form:"wilayah"`
	Kode_Unit      *string `form:"kode_unit"`
	Kode_Bucket    *string `form:"kode_bucket"`
}
type ParamsDetailPlanBulanById struct {
	Kode_Plan string `form:"kode_plan"`
}

type ParamsCekPlan struct {
	Page           *string `form:"page"`
	Limit          *string `form:"limit"`
	Wilayah_Cabang *string `form:"wilayah_cabang"`
	Wilayah        *string `form:"wilayah"`
}

type ParamsRecovery struct {
	Page           *string `form:"page"`
	Limit          *string `form:"limit"`
	Wilayah_Cabang *string `form:"wilayah_cabang"`
	Wilayah        *string `form:"wilayah"`
	Kode_Unit      *string `form:"kode_unit"`
	Kode_Bucket    *string `form:"kode_bucket"`
}
type ResponseResumePlanCol struct {
	Cabang          *string `json:"cabang"`
	Keterangan      *string `json:"keterangan"`
	Os              *string `json:"os"`
	Noa_Coll        *string `json:"noa_coll"`
	Os_Coll         *string `json:"os_coll"`
	Noa_Input_Pokok *string `json:"noa_input_pokok"`
	Os_Input_Pokok  *string `json:"os_input_pokok"`
	Noa_Po          *string `json:"noa_po"`
	Os_Po           *string `json:"os_po"`
	Noa_Luhap       *string `json:"noa_luhap"`
	Os_Luhap        *string `json:"os_luhap"`
	Noa_Lelang      *string `json:"noa_lelang"`
	Os_Lelang       *string `json:"os_lelang"`
	Noa_3r          *string `json:"noa_3r"`
	Os_3r           *string `json:"os_3r"`
	Status          *string `json:"status"`
}

type ParamsCbBulan struct {
	Page           *string `form:"page"`
	Limit          *string `form:"limit"`
	Wilayah_Cabang *string `form:"wilayah_cabang"`
	Search         *string `form:"search"`
}
type ResponseKirimPlan struct {
	Status *string `json:"status"`
}
type ResponseCBResumePlanCol struct {
	Unit            *string `json:"unit"`
	Os              *string `json:"os"`
	Noa_Coll        *string `json:"noa_coll"`
	Os_Coll         *string `json:"os_coll"`
	Noa_Input_Pokok *string `json:"noa_input_pokok"`
	Os_Input_Pokok  *string `json:"os_input_pokok"`
	Noa_Po          *string `json:"noa_po"`
	Os_Po           *string `json:"os_po"`
	Noa_Luhap       *string `json:"noa_luhap"`
	Os_Luhap        *string `json:"os_luhap"`
	Noa_Lelang      *string `json:"noa_lelang"`
	Os_Lelang       *string `json:"os_lelang"`
	Noa_3r          *string `json:"noa_3r"`
	Os_3r           *string `json:"os_3r"`
	Status          *string `json:"status"`
}

type ResponseResumePlanWo struct {
	Cabang            *string `json:"cabang"`
	Keterangan        *string `json:"keterangan"`
	Os                *string `json:"os"`
	Noa_Coll          *string `json:"noa_coll"`
	Os_Coll           *string `json:"os_coll"`
	Noa_Po_Sesuai_Os  *string `json:"noa_po_sesuai_os"`
	Os_Po_Sesuai_Os   *string `json:"os_po_sesuai_os"`
	Noa_Po_Dibawah_Os *string `json:"noa_po_dibawah_os"`
	Os_Po_Dibawah_OS  *string `json:"os_po_dibawah_os"`
	Noa_Lelang        *string `json:"noa_lelang"`
	Os_Lelang         *string `json:"os_lelang"`
	Status            *string `json:"status"`
}

type ResponseCbResumePlanWo struct {
	Unit              *string `json:"unit"`
	Os                *string `json:"os"`
	Noa_Coll          *string `json:"noa_coll"`
	Os_Coll           *string `json:"os_coll"`
	Noa_Po_Sesuai_Os  *string `json:"noa_po_sesuai_os"`
	Os_Po_Sesuai_Os   *string `json:"os_po_sesuai_os"`
	Noa_Po_Dibawah_Os *string `json:"noa_po_dibawah_os"`
	Os_Po_Dibawah_OS  *string `json:"os_po_dibawah_os"`
	Noa_Lelang        *string `json:"noa_lelang"`
	Os_Lelang         *string `json:"os_lelang"`
	Status            *string `json:"status"`
}
type ResponseResumeRealisasiCol struct {
	Cabang          *string `json:"cabang"`
	Noa_Baseon      *string `json:"noa_baseon"`
	Os_Baseon       *string `json:"os_baseon"`
	Noa_Coll        *string `json:"noa_coll"`
	Os_Coll         *string `json:"os_coll"`
	Noa_Input_Pokok *string `json:"noa_input_pokok"`
	Os_Input_Pokok  *string `json:"os_input_pokok"`
	Noa_Po          *string `json:"noa_po"`
	Os_Po           *string `json:"os_po"`
	Noa_Luhap       *string `json:"noa_luhap"`
	Os_Luhap        *string `json:"os_luhap"`
	Noa_Lelang      *string `json:"noa_lelang"`
	Os_Lelang       *string `json:"os_lelang"`
	Noa_3r          *string `json:"noa_3r"`
	Os_3r           *string `json:"os_3r"`
	Noa_No_Plan     *string `json:"noa_no_plan"`
	Os_No_Plan      *string `json:"os_no_plan"`
}
type ResponseCbResumeRealisasiCol struct {
	Unit            *string `json:"unit"`
	Noa_Baseon      *string `json:"noa_baseon"`
	Os_Baseon       *string `json:"os_baseon"`
	Noa_Coll        *string `json:"noa_coll"`
	Os_Coll         *string `json:"os_coll"`
	Noa_Input_Pokok *string `json:"noa_input_pokok"`
	Os_Input_Pokok  *string `json:"os_input_pokok"`
	Noa_Po          *string `json:"noa_po"`
	Os_Po           *string `json:"os_po"`
	Noa_Luhap       *string `json:"noa_luhap"`
	Os_Luhap        *string `json:"os_luhap"`
	Noa_Lelang      *string `json:"noa_lelang"`
	Os_Lelang       *string `json:"os_lelang"`
	Noa_3r          *string `json:"noa_3r"`
	Os_3r           *string `json:"os_3r"`
	Noa_No_Plan     *string `json:"noa_no_plan"`
	Os_No_Plan      *string `json:"os_no_plan"`
}
type ResponseResumeRealisasiWo struct {
	Cabang            *string `json:"cabang"`
	Noa_Baseon        *string `json:"noa_baseon"`
	Os_Baseon         *string `json:"os_baseon"`
	Noa_Coll          *string `json:"noa_coll"`
	Os_Coll           *string `json:"os_coll"`
	Noa_Po_Sesuai_Os  *string `json:"noa_po_sesuai_os"`
	Os_Po_Sesuai_Os   *string `json:"os_po_sesuai_os"`
	Noa_Po_Dibawah_Os *string `json:"noa_po_dibawah_os"`
	Os_Po_Dibawah_Os  *string `json:"os_po_dibawah_os"`
	Noa_Lelang        *string `json:"noa_lelang"`
	Os_Lelang         *string `json:"os_lelang"`
	Noa_No_Plan       *string `json:"noa_no_plan"`
	Os_No_Plan        *string `json:"os_no_plan"`
}

type ResponseCbResumeRealisasiWo struct {
	Unit              *string `json:"unit"`
	Noa_Baseon        *string `json:"noa_baseon"`
	Os_Baseon         *string `json:"os_baseon"`
	Noa_Coll          *string `json:"noa_coll"`
	Os_Coll           *string `json:"os_coll"`
	Noa_Po_Sesuai_Os  *string `json:"noa_po_sesuai_os"`
	Os_Po_Sesuai_Os   *string `json:"os_po_sesuai_os"`
	Noa_Po_Dibawah_Os *string `json:"noa_po_dibawah_os"`
	Os_Po_Dibawah_Os  *string `json:"os_po_dibawah_os"`
	Noa_Lelang        *string `json:"noa_lelang"`
	Os_Lelang         *string `json:"os_lelang"`
	Noa_No_Plan       *string `json:"noa_no_plan"`
	Os_No_Plan        *string `json:"os_no_plan"`
}

type ResponseNasabahWoPlanHari struct {
	No_Rekening          *string `json:"no_rekening"`
	Kode_Unit            *string `json:"kode_unit"`
	Unit                 *string `json:"unit"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_penyelesaian *string `json:"rencana_penyelesaian"`
	Tgl_Jb               *string `json:"tgl_jb"`
}
type ResponseNasabahPlanHari struct {
	Id                   *string `json:"id"`
	Unit                 *string `json:"unit"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Ft                   *string `json:"ft"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_penyelesaian *string `json:"rencana_penyelesaian"`
	Tgl_Angsuran         *string `json:"tgl_angsuran"`
	Tgl_Jb               *string `json:"tgl_jb"`
}

type ResponseDetailNasabahPlanHari struct {
	Id                     *string `json:"id"`
	Unit                   *string `json:"unit"`
	Kode_Unit              *string `json:"kode_unit"`
	Inisial_Cab            *string `json:"inisial_cab"`
	No_Rekening            *string `json:"no_rekening"`
	Nama_Nasabah           *string `json:"nama_nasabah"`
	Kolektibilitas         *string `json:"kolektibilitas"`
	Ft                     *string `json:"ft"`
	Saldo_Nominatif        *string `json:"saldo_nominatif"`
	Tgl_Angsuran           *string `json:"tgl_angsuran"`
	Cara_Penanganan        *string `json:"cara_penanganan"`
	Rencana_penyelesaian   *string `json:"rencana_penyelesaian"`
	Input_pokok            *string `json:"input_pokok"`
	Nominal_lunas_bertahap *string `json:"nominal_lunas_bertahap"`
	Ket_rinci              *string `json:"ket_rinci"`
	Tgl_Jb                 *string `json:"tgl_jb"`
	Id_pic                 *string `json:"id_pic"`
	Inisial_Pic            *string `json:"inisial_pic"`
	Nama                   *string `json:"nama"`
	Jabatan                *string `json:"jabatan"`
	No_Hp                  *string `json:"no_hp"`
	Inisial_Assigment      *string `json:"inisial_account_assigment"`
	Nama_Assigment         *string `json:"nama_account_assigment"`
	Posisi_Assigment       *string `json:"posisi_account_assigment"`
	Hp_Assigment           *string `json:"hp_account_assigment"`
	No_Rekening_r          *string `json:"no_rekening_r"`
}

type ResponseDetailNasabahWoPlanHari struct {
	No_Rekening          *string `json:"no_rekening"`
	Kode_Unit            *string `json:"kode_unit"`
	Unit                 *string `json:"unit"`
	Inisial_Cab          *string `json:"inisial_cab"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_penyelesaian *string `json:"rencana_penyelesaian"`
	Nama_Pic             *string `json:"nama_pic"`
	No_Hp                *string `json:"no_hp"`
	Jabatan              *string `json:"jabatan"`
}
type ResponseDetailPlanNpl struct {
	Id                   *string `json:"id"`
	Cabang               *string `json:"cabang"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Ft                   *string `json:"ft"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Os_Penyelesaian      *string `json:"os_penyelesaian"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Wilayah_Rmd          *string `json:"wilayah_rmd"`
	Nama                 *string `json:"nama"`
	Bln_Plan             *string `json:"bln_plan"`
	Thn_Plan             *string `json:"thn_plan"`
	S_Realisasi          *string `json:"s_realisasi"`
}
type ResponseCbDetailPlanNpl struct {
	Id                      *string `json:"id"`
	No_Rekening             *string `json:"no_rekening"`
	Kode_Unit               *string `json:"kode_unit"`
	Unit                    *string `json:"unit"`
	Inisial_Cab             *string `json:"inisial_cab"`
	Nama_Nasabah            *string `json:"nama_nasabah"`
	Kolektibilitas          *string `json:"kolektibilitas"`
	Ft                      *string `json:"ft"`
	Saldo_Nominatif         *string `json:"saldo_nominatif"`
	Tgl_Angsuran            *string `json:"tgl_angsuran"`
	Tgl_Jb                  *string `json:"tgl_jb"`
	Bln_Plan                *string `json:"bln_plan"`
	Thn_Plan                *string `json:"thn_plan"`
	Cara_Penanganan         *string `json:"cara_penanganan"`
	Input_pokok             *string `json:"input_pokok"`
	Nominal_lunas_bertahap  *string `json:"nominal_lunas_bertahap"`
	Ket_rinci               *string `json:"ket_rinci"`
	Id_pic                  *string `json:"id_pic"`
	Tgl_angs_jb             *string `json:"tgl_angs_jb"`
	Os_penyelesaian         *string `json:"os_penyelesaian"`
	Rencana_penyelesaian    *string `json:"rencana_penyelesaian"`
	S_realisasi             *string `json:"s_realisasi"`
	Status                  *string `json:"status"`
	S_penyelesaian          *string `json:"s_penyelesaian"`
	S_penyelesaian_hari_ini *string `json:"s_penyelesaian_hari_ini"`
}

type ResponseDetailPlanWo struct {
	Id                   *string `json:"id"`
	Cabang               *string `json:"cabang"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Nominal_Pembayaran   *string `json:"nominal_pembayaran"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Wilayah_Rmd          *string `json:"wilayah_rmd"`
	Bln_Plan             *string `json:"bln_plan"`
	Thn_Plan             *string `json:"thn_plan"`
	S_Realisasi          *string `json:"s_realisasi"`
}
type ResponseCbDetailPlanWo struct {
	Id                      *string `json:"id"`
	No_Rekening             *string `json:"no_rekening"`
	Kode_Unit               *string `json:"kode_unit"`
	Unit                    *string `json:"unit"`
	Inisial_Cab             *string `json:"inisial_cab"`
	Nama_Nasabah            *string `json:"nama_nasabah"`
	Kolektibilitas          *string `json:"kolektibilitas"`
	Saldo_Nominatif         *string `json:"saldo_nominatif"`
	Tgl_Jb                  *string `json:"tgl_jb"`
	Bln_Plan                *string `json:"bln_plan"`
	Thn_Plan                *string `json:"thn_plan"`
	Cara_Penanganan         *string `json:"cara_penanganan"`
	Rencana_penyelesaian    *string `json:"rencana_penyelesaian"`
	Nominal_Pembayaran      *string `json:"nominal_pembayaran"`
	Ket_Rinci               *string `json:"ket_rinci"`
	Id_pic                  *string `json:"id_pic"`
	S_Realisasi             *string `json:"s_realisasi"`
	Status                  *string `json:"status"`
	S_penyelesaian          *string `json:"s_penyelesaian"`
	S_penyelesaian_hari_ini *string `json:"s_penyelesaian_hari_ini"`
}

type ResponseDetailRealisasi struct {
	Id              *string `json:"id"`
	Cabang          *string `json:"cabang"`
	No_Rekening     *string `json:"no_rekening"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	Saldo_Nominatif *string `json:"saldo_nominatif"`
	Os_Plan         *string `json:"os_plan"`
	Os_Selesai      *string `json:"os_selesai"`
	Os_Realisasi    *string `json:"os_realisasi"`
	Status          *string `json:"status"`
}

type ResponseCbDetailRealisasi struct {
	Id              *string `json:"id"`
	Unit            *string `json:"unit"`
	No_Rekening     *string `json:"no_rekening"`
	Nama_Nasabah    *string `json:"nama_nasabah"`
	Saldo_Nominatif *string `json:"saldo_nominatif"`
	Os_Plan         *string `json:"os_plan"`
	Os_Selesai      *string `json:"os_selesai"`
	Os_Realisasi    *string `json:"os_realisasi"`
	Status          *string `json:"status"`
}

type ResponseHariDetailNplCB struct {
	Id                     *string `json:"id"`
	Tgl_Plan               *string `json:"tgl_plan"`
	Unit                   *string `json:"unit"`
	Kode_Unit              *string `json:"kode_unit"`
	Inisial_Cab            *string `json:"inisial_cab"`
	No_Rekening            *string `json:"no_rekening"`
	Nama_Nasabah           *string `json:"nama_nasabah"`
	Kolektibilitas         *string `json:"kolektibilitas"`
	Ft                     *string `json:"ft"`
	Saldo_Nominatif        *string `json:"saldo_nominatif"`
	Tgl_Angsuran           *string `json:"tgl_angsuran"`
	Cara_Penanganan        *string `json:"cara_penanganan"`
	Rencana_Penyelesaian   *string `json:"rencana_penyelesaian"`
	Input_Pokok            *string `json:"input_pokok"`
	Nominal_Lurus_Bertahap *string `json:"nominal_lurus_bertahap"`
	Ket_Rinci              *string `json:"ket_rinci"`
	Id_Pic                 *string `json:"id_pic"`
	Os_Penyelesaian        *string `json:"os_penyelesaian"`
	Inisial_Pic            *string `json:"inisial_pic"`
	Nama                   *string `json:"nama"`
	Jabatan                *string `json:"jabatan"`
	No_Hp                  *string `json:"no_hp"`
	As_inisial             *string `json:"as_inisial"`
	As_nama                *string `json:"as_nama"`
	As_posisi              *string `json:"as_posisi"`
	As_no_hp               *string `json:"as_no_hp"`
	No_rekening_r          *string `json:"no_rekening_r"`
}

type ResponseHariDetailWoCB struct {
	Id                   *string `json:"id"`
	Unit                 *string `json:"unit"`
	Tgl_Plan             *string `json:"tgl_plan"`
	Kode_Unit            *string `json:"kode_unit"`
	Inisial_Cab          *string `json:"inisial_cab"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Nominal_Pembayaran   *string `json:"nominal_pembayaran"`
	Ket_Rinci            *string `json:"ket_rinci"`
	Id_Pic               *string `json:"id_pic"`
	Inisial_Pic          *string `json:"inisial_pic"`
	Nama                 *string `json:"nama"`
	Jabatan              *string `json:"jabatan"`
	No_Hp                *string `json:"no_hp"`
}
type ResponseDetailRealisasiNplById struct {
	Id                     *string `json:"id"`
	Tgl_Plan               *string `json:"tgl_plan"`
	Unit                   *string `json:"unit"`
	Kode_Unit              *string `json:"kode_unit"`
	Inisial_Cab            *string `json:"inisial_cab"`
	No_Rekening            *string `json:"no_rekening"`
	Nama_Nasabah           *string `json:"nama_nasabah"`
	Kolektibilitas         *string `json:"kolektibilitas"`
	Ft                     *string `json:"ft"`
	Saldo_Nominatif        *string `json:"saldo_nominatif"`
	Tgl_Angsuran           *string `json:"tgl_angsuran"`
	Cara_Penanganan        *string `json:"cara_penanganan"`
	Rencana_Penyelesaian   *string `json:"rencana_penyelesaian"`
	Input_Pokok            *string `json:"input_pokok"`
	Nominal_Lurus_Bertahap *string `json:"nominal_lurus_bertahap"`
	Ket_Rinci              *string `json:"ket_rinci"`
	Id_Pic                 *string `json:"id_pic"`
	Os_Penyelesaian        *string `json:"os_penyelesaian"`
	Inisial_Pic            *string `json:"inisial_pic"`
	Nama                   *string `json:"nama"`
	Jabatan                *string `json:"jabatan"`
	No_Hp                  *string `json:"no_hp"`
	Status_Penyelesaian    *string `json:"status_penyelesaian"`
	Hasil_Penyelesaian     *string `json:"hasil_penyelesaian"`
	Ket_Penyelesaian       *string `json:"ket_penyelesaian"`
	Id_Pic_Selesai         *string `json:"id_pic_selesai"`
	Inisial_Pic_Selesai    *string `json:"inisial_pic_selesai"`
	Nama_Pic_Selesai       *string `json:"nama_pic_selesai"`
	Jabatan_Selesai        *string `json:"jabatan_selesai"`
	No_Hp_Selesai          *string `json:"no_hp_selesai"`
	Tgl_Jb                 *string `json:"tgl_jb"`
	Transaksi              *string `json:"transaksi"`
}
type ResponseDetailRealisasiWoById struct {
	Id                         *string `json:"id"`
	Tgl_Plan                   *string `json:"tgl_plan"`
	Unit                       *string `json:"unit"`
	Kode_Unit                  *string `json:"kode_unit"`
	Inisial_Cab                *string `json:"inisial_cab"`
	No_Rekening                *string `json:"no_rekening"`
	Nama_Nasabah               *string `json:"nama_nasabah"`
	Kolektibilitas             *string `json:"kolektibilitas"`
	Saldo_Nominatif            *string `json:"saldo_nominatif"`
	Cara_Penanganan            *string `json:"cara_penanganan"`
	Rencana_Penyelesaian       *string `json:"rencana_penyelesaian"`
	Nominal_Pembayaran         *string `json:"nominal_pembayaran"`
	Nominal_Pembayaran_Selesai *string `json:"nominal_pembayaran_selesai"`
	Ket_Rinci                  *string `json:"ket_rinci"`
	Id_Pic                     *string `json:"id_pic"`
	Inisial_Pic                *string `json:"inisial_pic"`
	Nama                       *string `json:"nama"`
	Jabatan                    *string `json:"jabatan"`
	No_Hp                      *string `json:"no_hp"`
	Status_Penyelesaian        *string `json:"status_penyelesaian"`
	Hasil_Penyelesaian         *string `json:"hasil_penyelesaian"`
	Ket_Penyelesaian           *string `json:"ket_penyelesaian"`
	Id_Pic_Selesai             *string `json:"id_pic_selesai"`
	Inisial_Pic_Selesai        *string `json:"inisial_pic_selesai"`
	Nama_Pic_Selesai           *string `json:"nama_pic_selesai"`
	Jabatan_Selesai            *string `json:"jabatan_selesai"`
	No_Hp_Selesai              *string `json:"no_hp_selesai"`
	Tgl_Jb                     *string `json:"tgl_jb"`
	Transaksi                  *string `json:"transaksi"`
}
type ResponsePapelineBulanResume struct {
	Inisial_Cabang   *string `json:"inisial_cabang"`
	Nama_Cabang      *string `json:"nama_cabang"`
	Keterangan       *string `json:"keterangan"`
	Noa_Coll         *string `json:"noa_coll"`
	Coll             *string `json:"coll"`
	Noa_3r_Baru      *string `json:"noa_3r_baru"`
	R3_Baru          *string `json:"_3r_baru"`
	Noa_3r           *string `json:"noa_3r"`
	R3               *string `json:"_3r"`
	Noa_Po_Sekaligus *string `json:"noa_po_sekaligus"`
	Po_Sekaligus     *string `json:"po_sekaligus"`
	Noa_po_bertahap  *string `json:"noa_po_bertahap"`
	Po_bertahap      *string `json:"po_bertahap"`
	Noa_lelang       *string `json:"noa_lelang"`
	Lelang           *string `json:"lelang"`
	Noa_back_to_0    *string `json:"noa_back_to_0"`
	Back_to_0        *string `json:"back_to_0"`
	Noa_back_1_30    *string `json:"noa_back_1_30"`
	Back_1_30        *string `json:"back_1_30"`
	Noa_back_31_60   *string `json:"noa_back_31_60"`
	Back_31_60       *string `json:"back_31_60"`
	Noa_rembes_npl   *string `json:"noa_rembes_npl"`
	Rembes_Npl       *string `json:"rembes_npl"`
}

type ResponsePapelineBulanResumeNpl struct {
	Inisial_Cabang   *string `json:"inisial_cabang"`
	Nama_Cabang      *string `json:"nama_cabang"`
	Noa_Coll         *string `json:"noa_coll"`
	Coll             *string `json:"coll"`
	Noa_Input_Pokok  *string `json:"noa_input_pokok"`
	Input_Pokok      *string `json:"input_pokok"`
	Noa_3r_Baru      *string `json:"noa_3r_baru"`
	R3_Baru          *string `json:"_3r_baru"`
	Noa_3r           *string `json:"noa_3r"`
	R3               *string `json:"_3r"`
	Noa_Po_Sekaligus *string `json:"noa_po_sekaligus"`
	Po_Sekaligus     *string `json:"po_sekaligus"`
	Noa_po_bertahap  *string `json:"noa_po_bertahap"`
	Po_bertahap      *string `json:"po_bertahap"`
	Noa_lelang       *string `json:"noa_lelang"`
	Lelang           *string `json:"lelang"`
	Noa_back_to_0    *string `json:"noa_back_to_0"`
	Back_to_0        *string `json:"back_to_0"`
	Noa_back_1_30    *string `json:"noa_back_1_30"`
	Back_1_30        *string `json:"back_1_30"`
	Noa_back_31_60   *string `json:"noa_back_31_60"`
	Back_31_60       *string `json:"back_31_60"`
	Noa_Stay_Npl     *string `json:"noa_stay_npl"`
	Stay_Npl         *string `json:"stay_npl"`
}

type ResponsePapelineBulanResumeTotal struct {
	Inisial_Cabang         *string `json:"inisial_cabang"`
	Nama_Cabang            *string `json:"nama_cabang"`
	Noa_Coll_Npl           *string `json:"noa_coll_npl"`
	Coll_Npl               *string `json:"col_npll"`
	Noa_Input_Pokok_Npl    *string `json:"noa_input_pokok_npl"`
	Input_Pokok_Npl        *string `json:"input_pokok_npl"`
	Os_Input_Pokok_Npl     *string `json:"os_input_pokok_npl"`
	Noa_3r_Npl_Baru        *string `json:"noa_3r_npl_baru"`
	R3_Npl_Baru            *string `json:"_3r_npl_baru"`
	Noa_3r_Npl             *string `json:"noa_3r_npl"`
	R3_Npl                 *string `json:"_3r_npl"`
	Noa_Po_Sekaligus_Npl   *string `json:"noa_po_sekaligus_npl"`
	Po_Sekaligus_Npl       *string `json:"po_sekaligus_npl"`
	Noa_Po_Bertahap_Npl    *string `json:"noa_po_bertahap_npl"`
	Po_Bertahap_Npl        *string `json:"po_bertahap_npl"`
	Os_Po_Bertahap_Npl     *string `json:"os_po_bertahap_npl"`
	Noa_Stay_Npl_Npl       *string `json:"noa_stay_npl_npl"`
	Stay_Npl_Npl           *string `json:"stay_npl_npl"`
	Noa_Coll_61_90         *string `json:"noa_coll_61_90"`
	Coll_61_90             *string `json:"coll_61_90"`
	Noa_3r_61_90_Baru      *string `json:"noa_3r_61_90_baru"`
	R3_61_90_Baru          *string `json:"_3r_61_90_baru"`
	Noa_R3_61_90           *string `json:"noa_3r_61_90"`
	R3_61_90               *string `json:"_3r_61_90"`
	Noa_Po_Sekaligus_61_90 *string `json:"noa_po_sekaligus_61_90"`
	Po_Sekaligus_61_90     *string `json:"po_sekaligus_61_90"`
	Noa_Po_Bertahap_61_90  *string `json:"noa_po_bertahap_61_90"`
	Po_Bertahap_61_90      *string `json:"po_bertahap_61_90"`
	Os_Po_Bertahap_61_90   *string `json:"os_po_bertahap_61_90"`
	Noa_Rembes_Npl_61_90   *string `json:"noa_rembes_npl_61_90"`
	Rembes_Npl_61_90       *string `json:"rembes_npl_61_90"`
}

type ResponsePapelineBulanResumeWo struct {
	Inisial_Cabang *string `json:"inisial_cabang"`
	Nama_Cabang    *string `json:"nama_cabang"`
	Noa_Coll       *string `json:"noa_coll"`
	Coll           *string `json:"coll"`
	Noa_Po         *string `json:"noa_po"`
	Po             *string `json:"po"`
	Noa_Lelang     *string `json:"noa_lelang"`
	Lelang         *string `json:"lelang"`
}

type ResponseBulanDetailPlanNpl struct {
	Id                   *string `json:"id"`
	Nama_Unit            *string `json:"nama_unit"`
	Cabang               *string `json:"cabang"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Tipe_Kredit          *string `json:"tipe_kredit"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Ft                   *string `json:"ft"`
	Angs_Ke              *string `json:"angs_ke"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Tgl_Realisasi        *string `json:"tgl_realisasi"`
	Tgl_Jatuh_Tempo      *string `json:"tgl_jatuh_tempo"`
	Jml_Pinjaman         *string `json:"jml_pinjaman"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Nama_Pic             *string `json:"nama_pic"`
}

type ResponseBulanDetailPlanNplById struct {
	Id                   *string `json:"id"`
	Nama_Unit            *string `json:"nama_unit"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Tipe_Kredit          *string `json:"tipe_kredit"`
	Tgl_Realisasi        *string `json:"tgl_realisasi"`
	Tgl_Jatuh_Tempo      *string `json:"tgl_jatuh_tempo"`
	Jml_Pinjaman         *string `json:"jml_pinjaman"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Ft                   *string `json:"ft"`
	Jenis_Usaha          *string `json:"jenis_usaha"`
	Kondisi_Usaha        *string `json:"kondisi_usaha"`
	Sp_1                 *string `json:"sp_1"`
	Sp_2                 *string `json:"sp_2"`
	Sp_3                 *string `json:"sp_3"`
	Obyek_Agunan         *string `json:"obyek_angunan"`
	Dokumen_Agunan       *string `json:"dukumen_angunan"`
	Pengikatan           *string `json:"pengikatan"`
	Jenis_Pengikatan     *string `json:"jenis_pengikatan"`
	Status_Dokumen       *string `json:"status_dokumen"`
	Nilai_Saat_Ini       *string `json:"nilai_saat_ini"`
	Mudah_Jual           *string `json:"mudah_jual"`
	Keberadaan_Agunan    *string `json:"keberadaan_agunan"`
	Sengketa             *string `json:"sengketa"`
	S_Diputuskan_Kpp     *string `json:"s_diputuskan_kpp"`
	Status_Lelang        *string `json:"status_lelang"`
	Lelang_Ke            *string `json:"lelang_ke"`
	Permasalahan         *string `json:"permasalahan"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Tgl_Jb               *string `json:"tgl_jb"`
	Ket_Rinci            *string `json:"ket_rinci"`
	Inisial_Pic          *string `json:"inisial_pic"`
	Nama_Pic             *string `json:"nama_pic"`
	Jabatan_Pic          *string `json:"jabatan_pic"`
	No_Hp_Pic            *string `json:"no_hp_pic"`
}
type ResponseBulanDetailPlanWo struct {
	Id                   *string `json:"id"`
	Nama_Unit            *string `json:"nama_unit"`
	Cabang               *string `json:"cabang"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Tgl_Wo               *string `json:"tgl_Wo"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Nominal_Pembayaran   *string `json:"nominal_pembayaran"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
}

type ResponseBulanDetailPlanWoById struct {
	Id                   *string `json:"id"`
	Nama_Unit            *string `json:"nama_unit"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Tipe_Kredit          *string `json:"tipe_kredit"`
	Tgl_Realisasi        *string `json:"tgl_realisasi"`
	Tgl_Jatuh_Tempo      *string `json:"tgl_jatuh_tempo"`
	Jml_Pinjaman         *string `json:"jml_pinjaman"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Sp_1                 *string `json:"sp_1"`
	Sp_2                 *string `json:"sp_2"`
	Sp_3                 *string `json:"sp_3"`
	Obyek_Agunan         *string `json:"obyek_angunan"`
	Dokumen_Agunan       *string `json:"dukumen_angunan"`
	Pengikatan           *string `json:"pengikatan"`
	Jenis_Pengikatan     *string `json:"jenis_pengikatan"`
	Status_Dokumen       *string `json:"status_dokumen"`
	Nilai_Saat_Ini       *string `json:"nilai_saat_ini"`
	Mudah_Jual           *string `json:"mudah_jual"`
	Keberadaan_Agunan    *string `json:"keberadaan_agunan"`
	Sengketa             *string `json:"sengketa"`
	S_Diputuskan_Kpp     *string `json:"s_diputuskan_kpp"`
	Status_Lelang        *string `json:"status_lelang"`
	Lelang_Ke            *string `json:"lelang_ke"`
	Permasalahan         *string `json:"permasalahan"`
	Cara_Penanganan      *string `json:"cara_penanganan"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Nominal_Pembayaran   *string `json:"nominal_pembayaran"`
	Tgl_Jb               *string `json:"tgl_jb"`
	Ket_Rinci            *string `json:"ket_rinci"`
	Inisial_Pic          *string `json:"inisial_pic"`
	Nama_Pic             *string `json:"nama_pic"`
	Jabatan_Pic          *string `json:"jabatan_pic"`
	No_Hp_Pic            *string `json:"no_hp_pic"`
}
type ResponseCekPlan struct {
	Inisial_Cab    *string `json:"inisial_cab"`
	Nama           *string `json:"nama"`
	Id_Approve     *string `json:"id_approve"`
	Id_Approve_Pbu *string `json:"id_approve_pbu"`
}

type ResponseRecoveryKol struct {
	Cabang            *string `json:"cabang"`
	Unit              *string `json:"unit"`
	No_Rekening       *string `json:"no_rekening"`
	Nama_Nasabah      *string `json:"nama_nasabah"`
	Tipe_Kredit       *string `json:"tipe_kredit"`
	Ft                *string `json:"ft"`
	Kolektibilitas    *string `json:"kolektibilitas"`
	Saldo_Nominatif   *string `json:"saldo_nominatif"`
	Kategori_Recovery *string `json:"kategori_recovery"`
}
type ResponseRecoveryWo struct {
	Cabang            *string `json:"cabang"`
	Unit              *string `json:"unit"`
	No_Rekening       *string `json:"no_rekening"`
	Nama_Nasabah      *string `json:"nama_nasabah"`
	Jml_Wo            *string `json:"jml_wo"`
	Kolektibilitas    *string `json:"kolektibilitas"`
	Saldo_Nominatif   *string `json:"saldo_nominatif"`
	Kategori_Recovery *string `json:"kategori_recovery"`
}
type CBResponseBulan struct {
	Id                   *string `json:"id"`
	Nama_Unit            *string `json:"nama_unit"`
	No_Rekening          *string `json:"no_rekening"`
	Nama_Nasabah         *string `json:"nama_nasabah"`
	Tipe_Kredit          *string `json:"tipe_kredit"`
	Kolektibilitas       *string `json:"kolektibilitas"`
	Ft                   *string `json:"ft"`
	Angs_Ke              *string `json:"angs_ke"`
	Saldo_Nominatif      *string `json:"saldo_nominatif"`
	Rencana_Penyelesaian *string `json:"rencana_penyelesaian"`
	Id_Plan              *string `json:"id_plan"`
}

type CBResponseBulanResume struct {
	Noa *string `json:"noa_baseon"`
	Os  *string `json:"os_baseon"`
}
type CBResponseBulanResumeTable struct {
	Noa_back_to_61_90      *string `json:"noa_back_to_61_90"`
	Os_back_to_61_90       *string `json:"os_back_to_61_90"`
	Noa_back_to_31_60      *string `json:"noa_back_to_31_60"`
	Os_back_to_31_60       *string `json:"os_back_to_31_60"`
	Noa_back_to_1_30       *string `json:"noa_back_to_1_30"`
	Os_back_to_1_30        *string `json:"os_back_to_1_30"`
	Noa_back_to_current    *string `json:"noa_back_to_current"`
	Os_back_to_current     *string `json:"os_back_to_current"`
	Noa_3r                 *string `json:"noa_3r"`
	Os_3r                  *string `json:"os_3r"`
	Noa_lunas_sekaligus    *string `json:"noa_lunas_sekaligus"`
	Os_lunas_sekaligus     *string `json:"os_lunas_sekaligus"`
	Noa_lunas_bertahap     *string `json:"noa_lunas_bertahap"`
	Os_lunas_bertahap      *string `json:"os_lunas_bertahap"`
	Nominal_lunas_bertahap *string `json:"nominal_lunas_bertahap"`
	Noa_lelang_terjual     *string `json:"noa_lelang_terjual"`
	Os_lelang_terjual      *string `json:"os_lelang_terjual"`
	Noa_stay_npl           *string `json:"noa_stay_npl"`
	Os_stay_npl            *string `json:"os_stay_npl"`
}
type ResponseHistoriPlanHari struct {
	Nama      *string `json:"nama"`
	Jabatan   *string `json:"jabatan"`
	Komen     *string `json:"komen"`
	Status    *string `json:"status"`
	Tgl_kirim *string `json:"tgl_kirim"`
}

package response

import "database/sql"

type ResumePlanKol1 struct {
	// Kode_Unit             string          `json:"kode_unit"`
	NoaTtpPdBucket130     int             `json:"noa_ttp_pd_bucket_1_30"`
	OsTtpPdBucket130      sql.NullFloat64 `json:"os_ttp_pd_bucket_1_30"`
	NoaBackToCurrent      int             `json:"noa_back_to_current"`
	OsBackToCurrent       sql.NullFloat64 `json:"os_back_to_current"`
	Noa3r                 int             `json:"noa_3r"`
	Os3r                  sql.NullFloat64 `json:"os_3r"`
	NoaPelunasanSekaligus int             `json:"noa_pelunasan_sekaligus"`
	OsPelunasanSekaligus  sql.NullFloat64 `json:"os_pelunasan_sekaligus"`
	NoaLelangTerjual      int             `json:"noa_lelang_terjual"`
	OsLelangTerjual       sql.NullFloat64 `json:"os_lelang_terjual"`
	NoaRembes3160         int             `json:"noa_rembes_31_60"`
	OsRembes3160          sql.NullFloat64 `json:"os_rembes_31_60"`
	NoaRembes130          int             `json:"noa_rembes_1_30"`
	OsRembes130           sql.NullFloat64 `json:"os_rembes_1_30"`
}

type ResumePlanKol2 struct {
	NoaTtpPdBucket6190     int             `json:"noa_ttp_pd_bucket_61_90"`
	OsTtpPdBucket6190      sql.NullFloat64 `json:"os_ttp_pd_bucket_61_90"`
	NoaTtpPdBucket130      int             `json:"noa_ttp_pd_bucket_1_30"`
	OsTtpPdBucket130       sql.NullFloat64 `json:"os_ttp_pd_bucket_1_30"`
	NoaTtpPdBucket3160     int             `json:"noa_ttp_pd_bucket_31_60"`
	OsTtpPdBucket3160      sql.NullFloat64 `json:"os_ttp_pd_bucket_31_60"`
	NoaBackTo3160          int             `json:"noa_back_to_31_60"`
	OsBackTo3160           sql.NullFloat64 `json:"os_back_to_31_60"`
	NoaBackTo130           int             `json:"noa_back_to_1_30"`
	OsBackTo130            sql.NullFloat64 `json:"os_back_to_1_30"`
	NoaBackToCurrent       int             `json:"noa_back_to_current"`
	OsBackToCurrent        sql.NullFloat64 `json:"os_back_to_current"`
	Noa3r                  int             `json:"noa_3r"`
	Os3r                   sql.NullFloat64 `json:"os_3r"`
	NoaPelunasanSekaligus  int             `json:"noa_pelunasan_sekaligus"`
	OsPelunasanSekaligus   sql.NullFloat64 `json:"os_pelunasan_sekaligus"`
	NoaLelangTerjual       int             `json:"noa_lelang_terjual"`
	OsLelangTerjual        sql.NullFloat64 `json:"os_lelang_terjual"`
	NoaPelunasanBertahap   sql.NullInt64   `json:"noa_pelunasan_bertahap"`
	OsPelunasanBertahap    sql.NullInt64   `json:"os_pelunasan_bertahap"`
	OsNominalLunasBertahap sql.NullFloat64 `json:"os_nominal_lunas_bertahap"`
	NoaRembes6190          int             `json:"noa_rembes_61_90"`
	OsRembes6190           sql.NullFloat64 `json:"os_rembes_61_90"`
	NoaRembesNPL           int             `json:"noa_rembes_npl"`
	OsRembesNPL            sql.NullFloat64 `json:"os_rembes_npl"`
}

type ResumePlanKol34 struct {
	NoaBackTo6190        int             `json:"noa_back_to_61_90"`
	OsBackTo6190         sql.NullFloat64 `json:"os_back_to_61_90"`
	NoaBackTo3160        int             `json:"noa_back_to_31_60"`
	OsBackTo3160         sql.NullFloat64 `json:"os_back_to_31_60"`
	NoaBackTo130         int             `json:"noa_back_to_1_30"`
	OsBackTo130          sql.NullFloat64 `json:"os_back_to_1_30"`
	NoaBackToCurrent     int             `json:"noa_back_to_current"`
	OsBackToCurrent      sql.NullFloat64 `json:"os_back_to_current"`
	Noa3r                int             `json:"noa_3r"`
	Os3r                 sql.NullFloat64 `json:"os_3r"`
	NoaLunasSekaligus    int             `json:"noa_pelunasan_sekaligus"`
	OsLunasSekaligus     sql.NullFloat64 `json:"os_pelunasan_sekaligus"`
	NoaLunasBertahap     sql.NullInt64   `json:"noa_pelunasan_bertahap"`
	OsLunasBertahap      sql.NullInt64   `json:"os_pelunasan_bertahap"`
	NominalLunasBertahap sql.NullFloat64 `json:"nominal_lunas_bertahap"`
	NoaLelangTerjual     int             `json:"noa_lelang_terjual"`
	OsLelangTerjual      sql.NullFloat64 `json:"os_lelang_terjual"`
	NoaStayNPL           int             `json:"noa_stay_npl"`
	OsStayNPL            sql.NullFloat64 `json:"os_stay_npl"`
}

type ResumePlanKol5 struct {
	NoaBackTo6190        int             `json:"noa_back_to_61_90"`
	OsBackTo6190         sql.NullFloat64 `json:"os_back_to_61_90"`
	NoaBackTo3160        int             `json:"noa_back_to_31_60"`
	OsBackTo3160         sql.NullFloat64 `json:"os_back_to_31_60"`
	NoaBackTo130         int             `json:"noa_back_to_1_30"`
	OsBackTo130          sql.NullFloat64 `json:"os_back_to_1_30"`
	NoaBackToCurrent     int             `json:"noa_back_to_current"`
	OsBackToCurrent      sql.NullFloat64 `json:"os_back_to_current"`
	NoaInputPokok        int             `json:"noa_input_pokok"`
	OsInputPokok         sql.NullFloat64 `json:"os_input_pokok"`
	Noa3r                int             `json:"noa_3r"`
	Os3r                 sql.NullFloat64 `json:"os_3r"`
	NoaLunasSekaligus    int             `json:"noa_pelunasan_sekaligus"`
	OsLunasSekaligus     sql.NullFloat64 `json:"os_pelunasan_sekaligus"`
	NoaLunasBertahap     sql.NullInt64   `json:"noa_pelunasan_bertahap"`
	OsLunasBertahap      sql.NullInt64   `json:"os_pelunasan_bertahap"`
	NominalLunasBertahap sql.NullFloat64 `json:"nominal_lunas_bertahap"`
	NoaLelangTerjual     int             `json:"noa_lelang_terjual"`
	OsLelangTerjual      sql.NullFloat64 `json:"os_lelang_terjual"`
	NoaStayNPL           int             `json:"noa_stay_npl"`
	OsStayNPL            sql.NullFloat64 `json:"os_stay_npl"`
}

type ResumeWo struct {
	NoaPickupAngsuran     int             `json:"noa_pickup_angsuran"`
	OsPickupAngsuran      sql.NullFloat64 `json:"os_pickup_angsuran"`
	NoaPelunasanSesuaiOS  int             `json:"noa_pelunasan_sesuai_os"`
	OsPelunasanSesuaiOS   sql.NullFloat64 `json:"os_pelunasan_sesuai_os"`
	NoaPelunasanDibawahOS int             `json:"noa_pelunasan_dibawah_os"`
	OsPelunasanDibawahOS  sql.NullFloat64 `json:"os_pelunasan_dibawah_os"`
	NoaLelangTerjual      int             `json:"noa_lelang_terjual"`
	OsLelangTerjual       sql.NullFloat64 `json:"os_lelang_terjual"`
}

type ResumeRembesNpl struct {
	NoaRembesNpl         int             `json:"noa_rembes_npl"`
	OsRembesNpl          sql.NullFloat64 `json:"os_rembes_npl"`
	NoaLunasBertahap     int             `json:"noa_lunas_bertahap"`
	OsLunasBertahap      sql.NullFloat64 `json:"os_lunas_bertahap"`
	NominalLunasBertahap sql.NullFloat64 `json:"nominal_lunas_bertahap"`
}

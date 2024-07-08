package entity

type ParamsGranTotal struct {
	WILAYAH string `json:"WILAYAH"`
}
type ParamsPenururnanNPL struct {
	PAGE    string `json:"PAGE"`
	LIMIT   string `json:"LIMIT"`
	WILAYAH string `json:"WILAYAH"`
}
type ResponsePenurunanNPL struct {
	Wil          *string `json:"wil"`
	Cabang       *string `json:"cabang"`
	Inisial_Cab  *string `json:"inisial_cab"`
	Baki_Debet   *string `json:"baki_debet"`
	Baseon_Npl   *string `json:"baseon_npl"`
	Npl_Saat_Ini *string `json:"npl_saat_ini"`
	Os_Target    *string `json:"os_target"`
	RMBS         *string `json:"rmbsn_sudah_flow"`
	Sisa_Potensi *string `json:"sisa_potensi"`
	Col_Npl      *string `json:"col_npl"`
	Col_Btb      *string `json:"col_btb"`
	Col_Btc      *string `json:"col_btc"`
	Phase_Out    *string `json:"phase_out"`
	Restruktur   *string `json:"restruktur"`
}
type ResponseGrandTotal struct {
	Baki_debet           *string `json:"baki_debet"`
	Baki_debet_baseon    *string `json:"baki_debet_baseon"`
	Baseon_npl           *string `json:"baseon_npl"`
	Os_Target            *string `json:"os_target"`
	Npl_saat_ini         *string `json:"npl_saat_ini"`
	Rmbsn_sudah_flow     *string `json:"rmbsn_sudah_flow"`
	Sisa_potensi_npl     *string `json:"sisa_potensi_npl"`
	Rencana_penyelesaian *string `json:"rencana_penyelesaian"`
	Col_npl              *string `json:"col_npl"`
	Col_btb_btc          *string `json:"col_btb_btc"`
	Phase_out            *string `json:"phase_out"`
	Restruktur           *string `json:"restruktur"`
}

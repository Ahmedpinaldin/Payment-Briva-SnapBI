package response

type GetPenangananId struct {
	Id_Pic      int    `json:"id_pic"`
	Inisial_Pic string `json:"inisial_pic"`
	Nama        string `json:"nama"`
	Jabatan     string `json:"jabatan"`
	No_Hp       string `json:"no_hp"`
}

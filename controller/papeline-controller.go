package controller

import (
	"log"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"
	"example.com/rest-api-recoll-mobile/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Define a helper function to return nil if the input is an empty string
func nilIfEmpty(s *string) *string {
	if s != nil && *s == "" || s != nil && *s == "0" || s != nil && *s == "ALL" {
		return nil
	}
	return s
}
func ListTableBulan(c *gin.Context) {
	var data entity.ParmsListBucketBulan
	c.Bind(&data)
	Wilayah := "'" + data.Wilayah + "'"
	Bulan := "'" + data.Bulan + "'"
	Tahun := "'" + data.Tahun + "'"

	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := `
	SELECT C.inisial_cab AS inisial_cab,
		C.nama AS nama_cabang,
		COUNT(CASE WHEN P.cara_penanganan='COLL' THEN P.cara_penanganan ELSE NULL END) AS noa_coll,
		SUM(CASE WHEN P.cara_penanganan='COLL' THEN N.saldo_nominatif ELSE 0 END) AS coll,
		COUNT(CASE WHEN P.cara_penanganan='3R' AND N.restruktur_ke=0 THEN P.id ELSE NULL END) AS noa_3r_baru,
		SUM(CASE WHEN P.cara_penanganan='3R' AND N.restruktur_ke=0 THEN N.saldo_nominatif ELSE 0 END) AS _3r_baru,
		COUNT(CASE WHEN P.cara_penanganan='3R' AND N.restruktur_ke>0 THEN P.id ELSE NULL END) AS noa_3r,
		SUM(CASE WHEN P.cara_penanganan='3R' AND N.restruktur_ke>0 THEN N.saldo_nominatif ELSE 0 END) AS _3r,
		COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id ELSE NULL END) AS noa_po_sekaligus,
		SUM(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif ELSE 0 END) AS po_sekaligus,
		COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.id ELSE NULL END) AS noa_po_bertahap,
		SUM(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.nominal_lunas_bertahap ELSE 0 END) AS po_bertahap,
		COUNT(CASE WHEN P.cara_penanganan='LELANG' THEN P.cara_penanganan ELSE NULL END) AS noa_lelang,
		SUM(CASE WHEN P.cara_penanganan='LELANG' THEN N.saldo_nominatif ELSE 0 END) AS lelang,
		COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id ELSE NULL END) AS noa_back_to_0,
		SUM(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif ELSE 0 END) AS back_to_0,
		COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN P.id ELSE NULL END) AS noa_back_1_30,
		SUM(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN N.saldo_nominatif ELSE 0 END) AS back_1_30,
		COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN P.id ELSE NULL END) AS noa_back_31_60,
		SUM(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN N.saldo_nominatif ELSE 0 END) AS back_31_60,
		COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN P.id ELSE NULL END) AS noa_rembes_npl,
		SUM(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN N.saldo_nominatif ELSE 0 END) AS rembes_npl
	FROM plan_bulan AS P
	LEFT JOIN nominatif_baseon AS N ON P.no_rekening = N.no_rekening
	LEFT JOIN wilayah_rmd AS W ON P.inisial_cab = W.inisial_cab
	LEFT JOIN cabang AS C ON P.inisial_cab = C.inisial_cab
	LEFT JOIN approve_bulan AS A ON P.inisial_cab = A.inisial_cab AND A.tipe = 'RMD'`

	if Bulan != "''" {
		query += " AND A.bln = " + Bulan
	}
	if Tahun != "''" {
		query += " AND A.thn = " + Tahun
	}
	query += " WHERE N.ft >= 61 AND N.ft <= 90 "
	if Bulan != "''" {
		query += " AND P.bln_plan = " + Bulan
	}
	if Tahun != "''" {
		query += " AND P.thn_plan = " + Tahun
	}
	query += " AND N.kolektibilitas IN ('L', 'PK')"
	if Wilayah != "''" {
		query += " AND (W.wilayah_rmd = 0 OR W.wilayah_rmd = " + Wilayah + ")"
	}
	query += " GROUP BY C.inisial_cab, C.nama ORDER BY C.nama"

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseBucketBulan

	for resp.Next() {
		var result entity.ResponseBucketBulan
		err = resp.Scan(
			&result.Inisial_cab,
			&result.Nama_Cabang,
			&result.Noa_Coll,
			&result.Coll,
			&result.Noa_3r_Baru,
			&result.N_3r_Baru,
			&result.Noa_3r,
			&result.N_3r,
			&result.Noa_Po_Sekaligus,
			&result.Po_Sekaligus,
			&result.Noa_Po_Bertahap,
			&result.Po_Bertahap,
			&result.Noa_Lelang,
			&result.Lelang,
			&result.Noa_Back_To_0,
			&result.Back_To_0,
			&result.Noa_Back_1_30,
			&result.Back_1_30,
			&result.Noa_Back_31_60,
			&result.Back_31_60,
			&result.Noa_Rembes_Npl,
			&result.Rembes_Npl,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
			return
		}
		results = append(results, result)
	}
	if err := resp.Err(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
		return
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func StatusPlan(c *gin.Context) {
	var limiter = rate.NewLimiter(rate.Limit(2), 5)
	if !limiter.Allow() {
		res := helper.BuildErrorResponse(400, "Too many requests. Try again later", "error", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var data entity.ParamsStatus
	c.Bind(&data)

	if err := validateParamsStatus(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string

	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_StatusPlan", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseStatus
	for rows.Next() {
		var result entity.ResponseStatus
		err = rows.Scan(
			&result.Bulan,
			&result.Hari,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
func ListResumePlanCol(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumePlanHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseResumePlanCol
	for rows.Next() {
		var result entity.ResponseResumePlanCol
		err = rows.Scan(
			&result.Cabang,
			&result.Keterangan,
			&result.Os,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Input_Pokok,
			&result.Os_Input_Pokok,
			&result.Noa_Po,
			&result.Os_Po,
			&result.Noa_Luhap,
			&result.Os_Luhap,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_3r,
			&result.Os_3r,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
func ListCbResumePlanCol(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()

	var cabang *string
	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "ResumePlanHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error executing stored procedure", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCBResumePlanCol
	for rows.Next() {
		var result entity.ResponseCBResumePlanCol
		err = rows.Scan(
			&result.Unit,
			&result.Os,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Input_Pokok,
			&result.Os_Input_Pokok,
			&result.Noa_Po,
			&result.Os_Po,
			&result.Noa_Luhap,
			&result.Os_Luhap,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_3r,
			&result.Os_3r,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error scanning row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListResumePlanWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH", Value: wilayah},
	}
	rows, err := service.ExecuteSP(db, "RMD_ResumePlanHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseResumePlanWo
	for rows.Next() {
		var result entity.ResponseResumePlanWo
		err = rows.Scan(
			&result.Cabang,
			&result.Keterangan,
			&result.Os,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Po_Sesuai_Os,
			&result.Os_Po_Sesuai_Os,
			&result.Noa_Po_Dibawah_Os,
			&result.Os_Po_Dibawah_OS,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListNasabahPlanHari(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string
	var search *string

	cabang = nilIfEmpty(data.Inisial_Cab)
	search = nilIfEmpty(data.Search)

	// ...

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_AddNasabah_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusTooManyRequests, res)
		return
	}

	var results []entity.ResponseNasabahPlanHari
	for rows.Next() {
		var result entity.ResponseNasabahPlanHari
		err = rows.Scan(
			&result.Id,
			&result.Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Saldo_Nominatif,
			&result.Cara_Penanganan,
			&result.Rencana_penyelesaian,
			&result.Tgl_Angsuran,
			&result.Tgl_Jb,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListNasabahWoPlanHari(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string
	var search *string

	cabang = nilIfEmpty(data.Inisial_Cab)
	search = nilIfEmpty(data.Search)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_AddNasabahWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusTooManyRequests, res)
		return
	}

	var results []entity.ResponseNasabahWoPlanHari
	for rows.Next() {
		var result entity.ResponseNasabahWoPlanHari
		err = rows.Scan(
			&result.No_Rekening,
			&result.Kode_Unit,
			&result.Unit,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Cara_Penanganan,
			&result.Rencana_penyelesaian,
			&result.Tgl_Jb,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailPlanHariCB(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	// IdUnit :=  "'"+ Kode +"'"
	// fromDate :=  "'"+ data.From_Date +"'"
	// toDate :=  "'"+ data.To_Date +"'"

	if err := validation.ValidateNoRekening(*data.No_Rekening); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var Id *string

	Id = nilIfEmpty(data.Id)
	params := []entity.Param{
		{Name: "ID", Value: Id},
		{Name: "NO_REKENING", Value: data.No_Rekening},
	}

	rows, err := service.ExecuteSP(db, "PlanHariKol_AddOrEdit_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var results []entity.ResponseHariDetailNplCB
	for rows.Next() {
		var result entity.ResponseHariDetailNplCB
		if Id != nil {
			err = rows.Scan(

				&result.Id,
				&result.Tgl_Plan,
				&result.Unit,
				&result.Kode_Unit,
				&result.Inisial_Cab,
				&result.No_Rekening,
				&result.Nama_Nasabah,
				&result.Kolektibilitas,
				&result.Ft,
				&result.Saldo_Nominatif,
				&result.Tgl_Angsuran,
				&result.Cara_Penanganan,
				&result.Rencana_Penyelesaian,
				&result.Input_Pokok,
				&result.Nominal_Lurus_Bertahap,
				&result.Ket_Rinci,
				&result.Id_Pic,
				&result.Os_Penyelesaian,
				&result.Inisial_Pic,
				&result.Nama,
				&result.Jabatan,
				&result.No_Hp,
				&result.As_inisial,
				&result.As_nama,
				&result.As_posisi,
				&result.As_no_hp,
				&result.No_rekening_r,
			)
		} else {
			err = rows.Scan(

				&result.Id,
				// &result.Tgl_Plan,
				&result.Unit,
				&result.Kode_Unit,
				&result.Inisial_Cab,
				&result.No_Rekening,
				&result.Nama_Nasabah,
				&result.Kolektibilitas,
				&result.Ft,
				&result.Saldo_Nominatif,
				&result.Tgl_Angsuran,
				&result.Cara_Penanganan,
				&result.Rencana_Penyelesaian,
				&result.Input_Pokok,
				&result.Nominal_Lurus_Bertahap,
				&result.Ket_Rinci,
				&result.Id_Pic,
				// &result.Os_Penyelesaian,
				&result.Inisial_Pic,
				&result.Nama,
				&result.Jabatan,
				&result.No_Hp,
				&result.As_inisial,
				&result.As_nama,
				&result.As_posisi,
				&result.As_no_hp,
				&result.No_rekening_r,
			)
		}

		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailPlanWoHariCB(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	if err := validation.ValidateNoRekening(*data.No_Rekening); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var Id *string

	Id = nilIfEmpty(data.Id)
	params := []entity.Param{
		{Name: "ID", Value: Id},
		{Name: "NO_REKENING", Value: data.No_Rekening},
	}

	rows, err := service.ExecuteSP(db, "PlanHariWo_AddOrEdit_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var results []entity.ResponseHariDetailWoCB
	for rows.Next() {
		var result entity.ResponseHariDetailWoCB
		if Id != nil {
			err = rows.Scan(
				&result.Id,
				&result.Unit,
				&result.Kode_Unit,
				&result.Tgl_Plan,
				&result.Inisial_Cab,
				&result.No_Rekening,
				&result.Nama_Nasabah,
				&result.Kolektibilitas,
				&result.Saldo_Nominatif,
				&result.Cara_Penanganan,
				&result.Rencana_Penyelesaian,
				&result.Nominal_Pembayaran,
				&result.Ket_Rinci,
				&result.Id_Pic,
				&result.Inisial_Pic,
				&result.Nama,
				&result.Jabatan,
				&result.No_Hp,
			)
		} else {
			err = rows.Scan(
				&result.Id,
				&result.Unit,
				&result.Kode_Unit,
				// &result.Tgl_Plan,
				&result.Inisial_Cab,
				&result.No_Rekening,
				&result.Nama_Nasabah,
				&result.Kolektibilitas,
				&result.Saldo_Nominatif,
				&result.Cara_Penanganan,
				&result.Rencana_Penyelesaian,
				&result.Nominal_Pembayaran,
				&result.Ket_Rinci,
				&result.Id_Pic,
				&result.Inisial_Pic,
				&result.Nama,
				&result.Jabatan,
				&result.No_Hp,
			)
		}

		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailNasabahPlanHari(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	if err := validateParamsId(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "ID", Value: data.Id},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_AddNasabah_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var results []entity.ResponseDetailNasabahPlanHari
	for rows.Next() {
		var result entity.ResponseDetailNasabahPlanHari
		err = rows.Scan(
			&result.Id,
			&result.Unit,
			&result.Kode_Unit,
			&result.Inisial_Cab,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Saldo_Nominatif,
			&result.Tgl_Angsuran,
			&result.Cara_Penanganan,
			&result.Rencana_penyelesaian,
			&result.Input_pokok,
			&result.Nominal_lunas_bertahap,
			&result.Ket_rinci,
			&result.Tgl_Jb,
			&result.Id_pic,
			&result.Inisial_Pic,
			&result.Nama,
			&result.Jabatan,
			&result.No_Hp,
			&result.Inisial_Assigment,
			&result.Nama_Assigment,
			&result.Posisi_Assigment,
			&result.Hp_Assigment,
			&result.No_Rekening_r,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailNasabahWoPlanHari(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	if err := validation.ValidateNoRekening(*data.No_Rekening); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "NO_REKENING", Value: data.No_Rekening},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_AddNasabahWo_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var results []entity.ResponseDetailNasabahWoPlanHari
	for rows.Next() {
		var result entity.ResponseDetailNasabahWoPlanHari
		err = rows.Scan(
			&result.No_Rekening,
			&result.Kode_Unit,
			&result.Unit,
			&result.Inisial_Cab,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Cara_Penanganan,
			&result.Rencana_penyelesaian,
			&result.Nama_Pic,
			&result.No_Hp,
			&result.Jabatan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbResumePlanWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	// IdUnit :=  "'"+ Kode +"'"
	// fromDate :=  "'"+ data.From_Date +"'"
	// toDate :=  "'"+ data.To_Date +"'"

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string
	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "ResumePlanHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbResumePlanWo
	for rows.Next() {
		var result entity.ResponseCbResumePlanWo
		err = rows.Scan(
			&result.Unit,
			&result.Os,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Po_Sesuai_Os,
			&result.Os_Po_Sesuai_Os,
			&result.Noa_Po_Dibawah_Os,
			&result.Os_Po_Dibawah_OS,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListDetailPlanNpl(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string
	var search *string

	wilayah = nilIfEmpty(data.Wilayah)
	search = nilIfEmpty(data.Search)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "WILAYAH", Value: wilayah},
	}
	rows, err := service.ExecuteSP(db, "RMD_PlanHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailPlanNpl
	for rows.Next() {
		var result entity.ResponseDetailPlanNpl
		err = rows.Scan(
			&result.Id,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Saldo_Nominatif,
			&result.Os_Penyelesaian,
			&result.Rencana_Penyelesaian,
			&result.Wilayah_Rmd,
			&result.Nama,
			&result.Bln_Plan,
			&result.Thn_Plan,
			&result.S_Realisasi,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbDetailPlanNpl(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string
	var search *string

	cabang = nilIfEmpty(data.Inisial_Cab)
	search = nilIfEmpty(data.Search)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}
	rows, err := service.ExecuteSP(db, "PlanHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbDetailPlanNpl
	for rows.Next() {
		var result entity.ResponseCbDetailPlanNpl
		err = rows.Scan(
			&result.Id,
			&result.No_Rekening,
			&result.Kode_Unit,
			&result.Unit,
			&result.Inisial_Cab,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Saldo_Nominatif,
			&result.Tgl_Angsuran,
			&result.Tgl_Jb,
			&result.Bln_Plan,
			&result.Thn_Plan,
			&result.Cara_Penanganan,
			&result.Input_pokok,
			&result.Nominal_lunas_bertahap,
			&result.Ket_rinci,
			&result.Id_pic,
			&result.Tgl_angs_jb,
			&result.Os_penyelesaian,
			&result.Rencana_penyelesaian,
			&result.S_realisasi,
			&result.Status,
			&result.S_penyelesaian,
			&result.S_penyelesaian_hari_ini,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
func ListDetailPlanWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string
	var search *string

	wilayah = nilIfEmpty(data.Wilayah)
	search = nilIfEmpty(data.Search)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_PlanHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailPlanWo
	for rows.Next() {
		var result entity.ResponseDetailPlanWo
		err = rows.Scan(
			&result.Id,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Saldo_Nominatif,
			&result.Nominal_Pembayaran,
			&result.Rencana_Penyelesaian,
			&result.Wilayah_Rmd,
			&result.Bln_Plan,
			&result.Thn_Plan,
			&result.S_Realisasi,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbDetailPlanWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string
	var search *string

	cabang = nilIfEmpty(data.Inisial_Cab)
	search = nilIfEmpty(data.Search)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "PlanHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbDetailPlanWo
	for rows.Next() {
		var result entity.ResponseCbDetailPlanWo
		err = rows.Scan(
			&result.Id,
			&result.No_Rekening,
			&result.Kode_Unit,
			&result.Unit,
			&result.Inisial_Cab,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Tgl_Jb,
			&result.Bln_Plan,
			&result.Thn_Plan,
			&result.Cara_Penanganan,
			&result.Rencana_penyelesaian,
			&result.Nominal_Pembayaran,
			&result.Ket_Rinci,
			&result.Id_pic,
			&result.S_Realisasi,
			&result.Status,
			&result.S_penyelesaian,
			&result.S_penyelesaian_hari_ini,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
func ListResumeRealisasiCol(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string
	wilayah = nilIfEmpty(data.Wilayah)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeRealisasiHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseResumeRealisasiCol
	for rows.Next() {
		var result entity.ResponseResumeRealisasiCol
		err = rows.Scan(
			&result.Cabang,
			&result.Noa_Baseon,
			&result.Os_Baseon,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Input_Pokok,
			&result.Os_Input_Pokok,
			&result.Noa_Po,
			&result.Os_Po,
			&result.Noa_Luhap,
			&result.Os_Luhap,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_3r,
			&result.Os_3r,
			&result.Noa_No_Plan,
			&result.Os_No_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbResumeRealisasiCol(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string

	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "ResumeRealisasiHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbResumeRealisasiCol
	for rows.Next() {
		var result entity.ResponseCbResumeRealisasiCol
		err = rows.Scan(
			&result.Unit,
			&result.Noa_Baseon,
			&result.Os_Baseon,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Input_Pokok,
			&result.Os_Input_Pokok,
			&result.Noa_Po,
			&result.Os_Po,
			&result.Noa_Luhap,
			&result.Os_Luhap,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_3r,
			&result.Os_3r,
			&result.Noa_No_Plan,
			&result.Os_No_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListResumeRealisasiWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)
	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeRealisasiHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseResumeRealisasiWo
	for rows.Next() {
		var result entity.ResponseResumeRealisasiWo
		err = rows.Scan(
			&result.Cabang,
			&result.Noa_Baseon,
			&result.Os_Baseon,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Po_Sesuai_Os,
			&result.Os_Po_Sesuai_Os,
			&result.Noa_Po_Dibawah_Os,
			&result.Os_Po_Dibawah_Os,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_No_Plan,
			&result.Os_No_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbResumeRealisasiWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var cabang *string

	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "ResumeRealisasiHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbResumeRealisasiWo
	for rows.Next() {
		var result entity.ResponseCbResumeRealisasiWo
		err = rows.Scan(
			&result.Unit,
			&result.Noa_Baseon,
			&result.Os_Baseon,
			&result.Noa_Coll,
			&result.Os_Coll,
			&result.Noa_Po_Sesuai_Os,
			&result.Os_Po_Sesuai_Os,
			&result.Noa_Po_Dibawah_Os,
			&result.Os_Po_Dibawah_Os,
			&result.Noa_Lelang,
			&result.Os_Lelang,
			&result.Noa_No_Plan,
			&result.Os_No_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListDetailRealisasiNpl(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailRealisasi
	for rows.Next() {
		var result entity.ResponseDetailRealisasi
		err = rows.Scan(
			// &result.Id,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Saldo_Nominatif,
			&result.Os_Plan,
			&result.Os_Selesai,
			&result.Os_Realisasi,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbDetailRealisasiNpl(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var search *string
	var cabang *string

	search = nilIfEmpty(data.Search)
	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "RealisasiHari_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbDetailRealisasi
	for rows.Next() {
		var result entity.ResponseCbDetailRealisasi
		err = rows.Scan(
			&result.Id,
			&result.Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Saldo_Nominatif,
			&result.Os_Plan,
			&result.Os_Selesai,
			&result.Os_Realisasi,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListDetailRealisasiWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParams(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	// var search *string
	var wilayah *string

	// search = nilIfEmpty(data.Search)
	wilayah = nilIfEmpty(data.Wilayah)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		// {Name: "SEARCH", Value: search},
		{Name: "WILAYAH", Value: wilayah},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailRealisasi
	for rows.Next() {
		var result entity.ResponseDetailRealisasi
		err = rows.Scan(
			// &result.Id,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Saldo_Nominatif,
			&result.Os_Plan,
			&result.Os_Selesai,
			&result.Os_Realisasi,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListCbDetailRealisasiWo(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsCab(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var search *string
	var cabang *string

	search = nilIfEmpty(data.Search)
	cabang = nilIfEmpty(data.Inisial_Cab)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "SEARCH", Value: search},
		{Name: "INISIAL_CAB", Value: cabang},
	}

	rows, err := service.ExecuteSP(db, "RealisasiHariWo_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCbDetailRealisasi
	for rows.Next() {
		var result entity.ResponseCbDetailRealisasi
		err = rows.Scan(
			&result.Id,
			&result.Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Saldo_Nominatif,
			&result.Os_Plan,
			&result.Os_Selesai,
			&result.Os_Realisasi,
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailRealisasiNplById(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)

	if err := validateParamsId(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "ID", Value: data.Id},
	}

	rows, err := service.ExecuteSP(db, "RealisasiHari_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailRealisasiNplById
	for rows.Next() {
		var result entity.ResponseDetailRealisasiNplById
		err = rows.Scan(
			&result.Id,
			&result.Tgl_Plan,
			&result.Unit,
			&result.Kode_Unit,
			&result.Inisial_Cab,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Saldo_Nominatif,
			&result.Tgl_Angsuran,
			&result.Cara_Penanganan,
			&result.Rencana_Penyelesaian,
			&result.Input_Pokok,
			&result.Nominal_Lurus_Bertahap,
			&result.Ket_Rinci,
			&result.Id_Pic,
			&result.Os_Penyelesaian,
			&result.Inisial_Pic,
			&result.Nama,
			&result.Jabatan,
			&result.No_Hp,
			&result.Status_Penyelesaian,
			&result.Hasil_Penyelesaian,
			&result.Ket_Penyelesaian,
			&result.Id_Pic_Selesai,
			&result.Inisial_Pic_Selesai,
			&result.Nama_Pic_Selesai,
			&result.Jabatan_Selesai,
			&result.No_Hp_Selesai,
			&result.Tgl_Jb,
			&result.Transaksi,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func DetailRealisasiWoById(c *gin.Context) {
	var data entity.ParamsTable
	c.Bind(&data)
	if err := validateParamsId(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "ID", Value: data.Id},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiHariWo_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailRealisasiWoById
	for rows.Next() {
		var result entity.ResponseDetailRealisasiWoById
		err = rows.Scan(
			&result.Id,
			&result.Tgl_Plan,
			&result.Unit,
			&result.Kode_Unit,
			&result.Inisial_Cab,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Cara_Penanganan,
			&result.Rencana_Penyelesaian,
			&result.Nominal_Pembayaran,
			&result.Nominal_Pembayaran_Selesai,
			&result.Ket_Rinci,
			&result.Id_Pic,
			&result.Inisial_Pic,
			&result.Nama,
			&result.Jabatan,
			&result.No_Hp,
			&result.Status_Penyelesaian,
			&result.Hasil_Penyelesaian,
			&result.Ket_Penyelesaian,
			&result.Id_Pic_Selesai,
			&result.Inisial_Pic_Selesai,
			&result.Nama_Pic_Selesai,
			&result.Jabatan_Selesai,
			&result.No_Hp_Selesai,
			&result.Tgl_Jb,
			&result.Transaksi,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanResume(c *gin.Context) {
	var data entity.ParamsPapelineBulanResume
	c.Bind(&data)

	if err := validateParamsBulan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayah *string
	var wilayahCabang *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeBulanBucket61", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponsePapelineBulanResume
	for rows.Next() {
		var result entity.ResponsePapelineBulanResume
		err = rows.Scan(
			&result.Inisial_Cabang,
			&result.Nama_Cabang,
			&result.Keterangan,
			&result.Noa_Coll,
			&result.Coll,
			&result.Noa_3r_Baru,
			&result.R3_Baru,
			&result.Noa_3r,
			&result.R3,
			&result.Noa_Po_Sekaligus,
			&result.Po_Sekaligus,
			&result.Noa_po_bertahap,
			&result.Po_bertahap,
			&result.Noa_lelang,
			&result.Lelang,
			&result.Noa_back_to_0,
			&result.Back_to_0,
			&result.Noa_back_1_30,
			&result.Back_1_30,
			&result.Noa_back_31_60,
			&result.Back_31_60,
			&result.Noa_rembes_npl,
			&result.Rembes_Npl,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanResumeNpl(c *gin.Context) {
	var data entity.ParamsPapelineBulanResume
	c.Bind(&data)

	if err := validateParamsBulan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeBulanBucketNPL", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponsePapelineBulanResumeNpl
	for rows.Next() {
		var result entity.ResponsePapelineBulanResumeNpl
		err = rows.Scan(
			&result.Inisial_Cabang,
			&result.Nama_Cabang,
			&result.Noa_Coll,
			&result.Coll,
			&result.Noa_Input_Pokok,
			&result.Input_Pokok,
			&result.Noa_3r_Baru,
			&result.R3_Baru,
			&result.Noa_3r,
			&result.R3,
			&result.Noa_Po_Sekaligus,
			&result.Po_Sekaligus,
			&result.Noa_po_bertahap,
			&result.Po_bertahap,
			&result.Noa_lelang,
			&result.Lelang,
			&result.Noa_back_to_0,
			&result.Back_to_0,
			&result.Noa_back_1_30,
			&result.Back_1_30,
			&result.Noa_back_31_60,
			&result.Back_31_60,
			&result.Noa_Stay_Npl,
			&result.Stay_Npl,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanResumeTotal(c *gin.Context) {
	var data entity.ParamsPapelineBulanResume
	c.Bind(&data)

	if err := validateParamsBulan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeBulanBucketTotal", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponsePapelineBulanResumeTotal
	for rows.Next() {
		var result entity.ResponsePapelineBulanResumeTotal
		err = rows.Scan(
			&result.Inisial_Cabang,
			&result.Nama_Cabang,
			&result.Noa_Coll_Npl,
			&result.Coll_Npl,
			&result.Noa_Input_Pokok_Npl,
			&result.Input_Pokok_Npl,
			&result.Os_Input_Pokok_Npl,
			&result.Noa_3r_Npl_Baru,
			&result.R3_Npl_Baru,
			&result.Noa_3r_Npl,
			&result.R3_Npl,
			&result.Noa_Po_Sekaligus_Npl,
			&result.Po_Sekaligus_Npl,
			&result.Noa_Po_Bertahap_Npl,
			&result.Po_Bertahap_Npl,
			&result.Os_Po_Bertahap_Npl,
			&result.Noa_Stay_Npl_Npl,
			&result.Stay_Npl_Npl,
			&result.Noa_Coll_61_90,
			&result.Coll_61_90,
			&result.Noa_3r_61_90_Baru,
			&result.R3_61_90_Baru,
			&result.Noa_R3_61_90,
			&result.R3_61_90,
			&result.Noa_Po_Sekaligus_61_90,
			&result.Po_Sekaligus_61_90,
			&result.Noa_Po_Bertahap_61_90,
			&result.Po_Bertahap_61_90,
			&result.Os_Po_Bertahap_61_90,
			&result.Noa_Rembes_Npl_61_90,
			&result.Rembes_Npl_61_90,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanResumeWo(c *gin.Context) {
	var data entity.ParamsPapelineBulanResume
	c.Bind(&data)

	if err := validateParamsBulan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "RMD_ResumeBulanBucketWO", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponsePapelineBulanResumeWo
	for rows.Next() {
		var result entity.ResponsePapelineBulanResumeWo
		err = rows.Scan(
			&result.Inisial_Cabang,
			&result.Nama_Cabang,
			&result.Noa_Coll,
			&result.Coll,
			&result.Noa_Po,
			&result.Po,
			&result.Noa_Lelang,
			&result.Lelang,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
func ListBulanDetailPlanNpl(c *gin.Context) {
	var data entity.ParamsDetailPlan
	c.Bind(&data)

	if err := validateDetailPlan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string
	var kode_unit *string
	var kode_bucket *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)
	kode_unit = nilIfEmpty(data.Kode_Unit)
	kode_bucket = nilIfEmpty(data.Kode_Bucket)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "KODE_UNIT", Value: kode_unit},
		{Name: "KODE_BUCKET", Value: kode_bucket},
	}

	rows, err := service.ExecuteSP(db, "RMD_DetailPlanBulanNPL", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseBulanDetailPlanNpl
	for rows.Next() {
		var result entity.ResponseBulanDetailPlanNpl
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Angs_Ke,
			&result.Saldo_Nominatif,
			&result.Tgl_Realisasi,
			&result.Tgl_Jatuh_Tempo,
			&result.Jml_Pinjaman,
			&result.Rencana_Penyelesaian,
			&result.Nama_Pic,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanDetailPlanNplById(c *gin.Context) {
	var data entity.ParamsDetailPlanBulanById
	c.Bind(&data)

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "KODE_PLAN", Value: data.Kode_Plan},
	}

	rows, err := service.ExecuteSP(db, "RMD_DetailPlanBulanNPL_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseBulanDetailPlanNplById
	for rows.Next() {
		var result entity.ResponseBulanDetailPlanNplById
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Tgl_Realisasi,
			&result.Tgl_Jatuh_Tempo,
			&result.Jml_Pinjaman,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Ft,
			&result.Jenis_Usaha,
			&result.Kondisi_Usaha,
			&result.Sp_1,
			&result.Sp_2,
			&result.Sp_3,
			&result.Obyek_Agunan,
			&result.Dokumen_Agunan,
			&result.Pengikatan,
			&result.Jenis_Pengikatan,
			&result.Status_Dokumen,
			&result.Nilai_Saat_Ini,
			&result.Mudah_Jual,
			&result.Keberadaan_Agunan,
			&result.Sengketa,
			&result.S_Diputuskan_Kpp,
			&result.Status_Lelang,
			&result.Lelang_Ke,
			&result.Permasalahan,
			&result.Cara_Penanganan,
			&result.Rencana_Penyelesaian,
			&result.Tgl_Jb,
			&result.Ket_Rinci,
			&result.Inisial_Pic,
			&result.Nama_Pic,
			&result.Jabatan_Pic,
			&result.No_Hp_Pic,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanDetailPlanWo(c *gin.Context) {
	var data entity.ParamsDetailPlan
	c.Bind(&data)

	if err := validateDetailPlan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string
	var kode_unit *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)
	kode_unit = nilIfEmpty(data.Kode_Unit)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "KODE_UNIT", Value: kode_unit},
	}

	rows, err := service.ExecuteSP(db, "RMD_DetailPlanBulanWO", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseBulanDetailPlanWo
	for rows.Next() {
		var result entity.ResponseBulanDetailPlanWo
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.Cabang,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tgl_Wo,
			&result.Saldo_Nominatif,
			&result.Nominal_Pembayaran,
			&result.Rencana_Penyelesaian,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanDetailPlanWoById(c *gin.Context) {
	var data entity.ParamsDetailPlanBulanById
	c.Bind(&data)

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "KODE_PLAN", Value: data.Kode_Plan},
	}

	rows, err := service.ExecuteSP(db, "RMD_DetailPlanBulanWO_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseBulanDetailPlanWoById
	for rows.Next() {
		var result entity.ResponseBulanDetailPlanWoById
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tgl_Realisasi,
			&result.Tgl_Jatuh_Tempo,
			&result.Jml_Pinjaman,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Sp_1,
			&result.Sp_2,
			&result.Sp_3,
			&result.Obyek_Agunan,
			&result.Dokumen_Agunan,
			&result.Pengikatan,
			&result.Jenis_Pengikatan,
			&result.Status_Dokumen,
			&result.Nilai_Saat_Ini,
			&result.Mudah_Jual,
			&result.Keberadaan_Agunan,
			&result.Sengketa,
			&result.S_Diputuskan_Kpp,
			&result.Status_Lelang,
			&result.Lelang_Ke,
			&result.Permasalahan,
			&result.Cara_Penanganan,
			&result.Rencana_Penyelesaian,
			&result.Tgl_Jb,
			&result.Ket_Rinci,
			&result.Inisial_Pic,
			&result.Nama_Pic,
			&result.Jabatan_Pic,
			&result.No_Hp_Pic,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListBulanCekPlan(c *gin.Context) {
	var data entity.ParamsCekPlan
	c.Bind(&data)

	if err := validateCekPlan(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "WIL_RMD", Value: wilayah},
	}
	rows, err := service.ExecuteSP(db, "RMD_PlanCabangBulan", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCekPlan
	for rows.Next() {
		var result entity.ResponseCekPlan
		err = rows.Scan(
			&result.Inisial_Cab,
			&result.Nama,
			&result.Id_Approve,
			&result.Id_Approve_Pbu,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListKolektibilitasRecovery(c *gin.Context) {
	var data entity.ParamsRecovery
	c.Bind(&data)

	if err := validateRecovery(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string
	var kode_unit *string
	var kode_bucket *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)
	kode_unit = nilIfEmpty(data.Kode_Unit)
	kode_bucket = nilIfEmpty(data.Kode_Bucket)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "KODE_UNIT", Value: kode_unit},
		{Name: "KODE_BUCKET", Value: kode_bucket},
	}

	rows, err := service.ExecuteSP(db, "RMD_RecoveryKol", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseRecoveryKol
	for rows.Next() {
		var result entity.ResponseRecoveryKol
		err = rows.Scan(
			&result.Cabang,
			&result.Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Ft,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Kategori_Recovery,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListWoRecovery(c *gin.Context) {
	var data entity.ParamsRecovery
	c.Bind(&data)

	if err := validateRecovery(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var wilayah *string
	var kode_unit *string

	wilayah = nilIfEmpty(data.Wilayah)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)
	kode_unit = nilIfEmpty(data.Kode_Unit)

	params := []entity.Param{

		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "KODE_UNIT", Value: kode_unit},
	}

	rows, err := service.ExecuteSP(db, "RMD_RecoveryWo", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseRecoveryWo
	for rows.Next() {
		var result entity.ResponseRecoveryWo
		err = rows.Scan(
			&result.Cabang,
			&result.Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Jml_Wo,
			&result.Kolektibilitas,
			&result.Saldo_Nominatif,
			&result.Kategori_Recovery,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func KirimPlanHari(c *gin.Context) {
	var data entity.ParamsKirim
	c.Bind(&data)

	if err := validateKirim(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	params := []entity.Param{
		{Name: "tipe", Value: data.Tipe},
		{Name: "inisial_cab_approve", Value: data.Inisial_Cab},
		{Name: "status", Value: data.Status},
		{Name: "user_approve", Value: data.User_Approve},
		{Name: "user_level", Value: data.User_Level},
		{Name: "komen", Value: data.Komen},
	}

	rows, err := service.ExecuteSP(db, "PlanHari_Kirim", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseKirimPlan
	for rows.Next() {
		var result entity.ResponseKirimPlan
		err = rows.Scan(
			&result.Status,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func SubmitPlanHari(c *gin.Context) {
	var requestData entity.ParamsSubmit

	err := c.BindJSON(&requestData)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error parsing JSON data", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error starting transaction", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Function to rollback the transaction and handle errors
	rollbackTx := func() {
		if err := tx.Rollback(); err != nil {
			res := helper.BuildErrorResponse(400, "Error rolling back transaction", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	for _, item := range requestData.Data {
		var noRek *string
		var noRekWo *string

		if item.NoRekening == "0" {
			noRek = nil
		} else {
			noRek = &item.NoRekening
		}
		if item.NoRekeningWo == "0" {
			noRekWo = nil
		} else {
			noRekWo = &item.NoRekeningWo
		}

		params := []entity.Param{
			{Name: "no_rekening", Value: noRek},
			{Name: "no_rekening_wo", Value: noRekWo},
			{Name: "nama_nasabah", Value: item.NamaNasabah},
			{Name: "kode_unit", Value: item.KodeUnit},
			{Name: "inisial_cab", Value: item.InisialCab},
			{Name: "kolektibilitas", Value: item.Kolektibilitas},
			{Name: "ft", Value: item.Ft},
			{Name: "saldo_nominatif", Value: item.SaldoNominatif},
			{Name: "nominal_pembayaran", Value: item.NominalPembayaran},
			{Name: "os_penyelesaian", Value: item.OsPenyelesaian},
			{Name: "input_pokok", Value: item.InputPokok},
			{Name: "nominal_lunas_bertahap", Value: item.NominalLunas},
			{Name: "cara_penanganan", Value: item.CaraPenanganan},
			{Name: "rencana_penyelesaian", Value: item.RencanaPenanganan},
			{Name: "tgl_angsuran", Value: item.TglAngsuran},
			{Name: "ket_rinci", Value: item.KetRinci},
			{Name: "id_pic", Value: item.IdPic},
		}

		_, err := service.ExecuteSP(db, "PlanHari_AddNasabahKol", params)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Gagal Input Data Nasabah", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			rollbackTx()
			return
		}
	}
	if err := tx.Commit(); err != nil {
		res := helper.BuildErrorResponse(400, "Error committing transaction:", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(200, "All inserts executed successfully.", requestData.Data)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBListPipelineBulanKol3(c *gin.Context) {
	var data entity.ParamsCbBulan
	c.Bind(&data)

	if err := validateBulanKol(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var search *string

	search = nilIfEmpty(data.Search)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{

		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "SEARCH", Value: search},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanCol3", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulan
	for rows.Next() {
		var result entity.CBResponseBulan
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Angs_Ke,
			&result.Saldo_Nominatif,
			&result.Rencana_Penyelesaian,
			&result.Id_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBListPipelineBulanKol4(c *gin.Context) {
	var data entity.ParamsCbBulan
	c.Bind(&data)

	if err := validateBulanKol(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var search *string
	search = nilIfEmpty(data.Search)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{

		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "SEARCH", Value: search},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanCol4", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulan
	for rows.Next() {
		var result entity.CBResponseBulan
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Angs_Ke,
			&result.Saldo_Nominatif,
			&result.Rencana_Penyelesaian,
			&result.Id_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBListPipelineBulanKol5(c *gin.Context) {
	var data entity.ParamsCbBulan
	c.Bind(&data)

	if err := validateBulanKol(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var search *string

	search = nilIfEmpty(data.Search)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{

		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "SEARCH", Value: search},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanCol5", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulan
	for rows.Next() {
		var result entity.CBResponseBulan
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Angs_Ke,
			&result.Saldo_Nominatif,
			&result.Rencana_Penyelesaian,
			&result.Id_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBListPipelineBulanWo(c *gin.Context) {
	var data entity.ParamsCbBulan
	c.Bind(&data)

	if err := validateBulanKol(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string
	var search *string

	search = nilIfEmpty(data.Search)
	wilayahCabang = nilIfEmpty(data.Wilayah_Cabang)

	params := []entity.Param{

		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "INISIAL_CAB", Value: wilayahCabang},
		{Name: "SEARCH", Value: search},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanWO", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulan
	for rows.Next() {
		var result entity.CBResponseBulan
		err = rows.Scan(
			&result.Id,
			&result.Nama_Unit,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Kolektibilitas,
			&result.Ft,
			&result.Angs_Ke,
			&result.Saldo_Nominatif,
			&result.Rencana_Penyelesaian,
			&result.Id_Plan,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBTotalResume(c *gin.Context) {
	var data entity.CBResumeTotal
	c.Bind(&data)

	if err := validateResume(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var inisialCab *string
	var kol *string

	if data.Inisial_Cab == "0" {
		inisialCab = nil
	} else {
		inisialCab = &data.Inisial_Cab
	}
	if data.Kol == "0" {
		kol = nil
	} else {
		kol = &data.Kol
	}

	params := []entity.Param{

		{Name: "KOL", Value: kol},
		{Name: "INISIAL_CAB", Value: inisialCab},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanResumeTotal", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulanResume
	for rows.Next() {
		var result entity.CBResponseBulanResume
		err = rows.Scan(
			&result.Noa,
			&result.Os,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func CBTableResume(c *gin.Context) {
	var data entity.CBResumeTotal
	c.Bind(&data)

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var inisialCab *string
	var kol *string

	if *&data.Kol == "0" {
		inisialCab = nil
	} else {
		inisialCab = &data.Inisial_Cab
	}
	if *&data.Kol == "0" {
		kol = nil
	} else {
		kol = &data.Kol
	}

	params := []entity.Param{

		{Name: "KOL", Value: kol},
		{Name: "INISIAL_CAB", Value: inisialCab},
	}

	rows, err := service.ExecuteSP(db, "CB_PipelineBulanResumeTable", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.CBResponseBulanResumeTable
	for rows.Next() {
		var result entity.CBResponseBulanResumeTable
		err = rows.Scan(
			&result.Noa_back_to_61_90,
			&result.Os_back_to_61_90,
			&result.Noa_back_to_31_60,
			&result.Os_back_to_31_60,
			&result.Noa_back_to_1_30,
			&result.Os_back_to_1_30,
			&result.Noa_back_to_current,
			&result.Os_back_to_current,
			&result.Noa_3r,
			&result.Os_3r,
			&result.Noa_lunas_sekaligus,
			&result.Os_lunas_sekaligus,
			&result.Noa_lunas_bertahap,
			&result.Os_lunas_bertahap,
			&result.Nominal_lunas_bertahap,
			&result.Noa_lelang_terjual,
			&result.Os_lelang_terjual,
			&result.Noa_stay_npl,
			&result.Os_stay_npl,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func ListHistoriPlanHari(c *gin.Context) {

	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateParamsHistory(data); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahCabang *string

	wilayahCabang = nilIfEmpty(&data.Wilayah_Cabang)

	params := []entity.Param{
		{Name: "INISIAL_CAB", Value: wilayahCabang},
	}

	rows, err := service.ExecuteSP(db, "History_Approval_Hari", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseHistoriPlanHari
	for rows.Next() {
		var result entity.ResponseHistoriPlanHari
		err = rows.Scan(
			&result.Nama,
			&result.Jabatan,
			&result.Komen,
			&result.Status,
			&result.Tgl_kirim,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func SubmitPeneyelesaian(c *gin.Context) {
	var requestData entity.ParamsPenyelesaian

	err := c.BindJSON(&requestData)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error parsing JSON data", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error starting transaction", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Function to rollback the transaction and handle errors
	rollbackTx := func() {
		if err := tx.Rollback(); err != nil {
			res := helper.BuildErrorResponse(400, "Error rolling back transaction", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}
	for _, item := range requestData.Data {
		var no_rek *string
		var no_rek_wo *string

		if *&item.NoRekening == "0" {
			no_rek = nil
		} else {
			no_rek = &item.NoRekening
		}
		if *&item.NoRekeningWo == "0" {
			no_rek_wo = nil
		} else {
			no_rek_wo = &item.NoRekeningWo
		}

		params := []entity.Param{
			{Name: "no_rekening", Value: no_rek},
			{Name: "no_rekening_wo", Value: no_rek_wo},
			{Name: "nama_nasabah", Value: item.NamaNasabah},
			{Name: "kode_unit", Value: item.KodeUnit},
			{Name: "inisial_cab", Value: item.InisialCab},
			{Name: "cara_penanganan", Value: item.CaraPenanganan},
			{Name: "rencana_penyelesaian", Value: item.RencanaPenanganan},
			{Name: "ket_rinci_plan", Value: item.KetPlan},
			{Name: "status_penyelesaian", Value: item.Status},
			{Name: "hasil_penyelesaian", Value: item.Hasil},
			{Name: "ket_rinci", Value: item.Ket},
			{Name: "nominal_pembayaran", Value: item.NominalPembayaran},
			{Name: "os_penyelesaian", Value: item.Os},
			{Name: "tgl_plan", Value: item.TglPlan},
			{Name: "tgl_jb", Value: item.TglJb},
			{Name: "id_pic", Value: item.IdPic},
		}

		_, err := service.ExecuteSP(db, "CB_Submit_Penyelesaian", params)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Gagal Input Data Penyelesaian", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			rollbackTx()
			return
		}
	}
	if err := tx.Commit(); err != nil {
		res := helper.BuildErrorResponse(400, "Error committing transaction:", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(200, "All inserts executed successfully.", requestData.Data)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func validateParams(data entity.ParamsTable) error {

	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(*data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateParamsCab(data entity.ParamsTable) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Inisial_Cab, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Inisial_Cab, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateParamsBulan(data entity.ParamsPapelineBulanResume) error {

	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(*data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}
func validateCekPlan(data entity.ParamsCekPlan) error {
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(*data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateRecovery(data entity.ParamsRecovery) error {

	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(*data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateDetailPlan(data entity.ParamsDetailPlan) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}
func validateBulanKol(data entity.ParamsCbBulan) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	// Validate PAGE
	if err := validation.IntegerInRange(*data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(*data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}
func validateParamsHistory(data entity.ParamsTransaksi) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}

	return nil
}
func validateParamsStatus(data entity.ParamsStatus) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Wilayah_Cabang, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Wilayah_Cabang, 4, "INISIAL CABANG"); err != nil {
		return err
	}

	return nil
}
func validateParamsId(data entity.ParamsTable) error {

	// Validate ID contains only integer
	if err := validation.JustInteger(*data.Id, "ID"); err != nil {
		return err
	}
	return nil
}

func validateKirim(data entity.ParamsKirim) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(*data.Inisial_Cab, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(*data.Inisial_Cab, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	if err := validation.StringMaxLength(*data.Tipe, 4, "Type"); err != nil {
		return err
	}
	if err := validation.ValidateNumberAlpha(*data.User_Approve); err != nil {
		return err
	}
	return nil
}
func validateResume(data entity.CBResumeTotal) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(data.Inisial_Cab, "INISIAL CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(data.Inisial_Cab, 4, "INISIAL CABANG"); err != nil {
		return err
	}
	return nil
}

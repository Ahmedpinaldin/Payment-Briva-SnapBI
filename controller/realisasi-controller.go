package controller

import (
	"fmt"
	"net/http"
	"time"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"
	"example.com/rest-api-recoll-mobile/validation"
	"github.com/gin-gonic/gin"
)

func nilIfEmptys(s *string) *string {
	if s != nil && *s == "" || s != nil && *s == "0" || s != nil && *s == "ALL" {
		return nil
	}
	return s
}
func ResumeTransaksiWidget(c *gin.Context) {
	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateWidget(data); err != nil {
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
	var dateAwal *string
	var dateAkhir *string
	var kodeUnit *string

	if *&data.Wilayah_Cabang == "0" {
		wilayahCabang = nil
	} else {
		wilayahCabang = &data.Wilayah_Cabang
	}
	if *&data.Wilayah == "0" {
		wilayah = nil
	} else {
		wilayah = &data.Wilayah
	}
	if *&data.Date_Awal == "0" {
		dateAwal = nil
	} else {
		dateAwal = &data.Date_Awal
	}
	if *&data.Date_Akhir == "0" {
		dateAkhir = nil
	} else {
		dateAkhir = &data.Date_Akhir
	}
	if *&data.UnitKode == "0" {
		kodeUnit = nil
	} else {
		kodeUnit = &data.UnitKode
	}

	params := []entity.Param{
		{Name: "WILAYAH_CABANG", Value: wilayahCabang},
		{Name: "WILAYAH", Value: wilayah},
		{Name: "UNIT_KODE", Value: kodeUnit},
		{Name: "DATE_AWAL", Value: dateAwal},
		{Name: "DATE_AKHIR", Value: dateAkhir},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiTransaksiResume", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseTransaksiResumeWidget
	for rows.Next() {
		var result entity.ResponseTransaksiResumeWidget
		err = rows.Scan(
			&result.Total_Noa,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
			&result.Total_Transaksi,
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

func getMonthDates() (string, string) {
	now := time.Now()
	dateAwal := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	dateAkhir := dateAwal.AddDate(0, 1, 0).Add(-time.Nanosecond)

	return dateAwal.Format("2006-1-2"), dateAkhir.Format("2006-1-2")
}

func ResumeTransaksiWidgetCab(c *gin.Context) {
	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateWidget(data); err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Validation Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()

	wilayahCabang, dateAwal, dateAkhir, kodeUnit := handleDatesAndParams(data)

	where := buildWhereClause(kodeUnit, wilayahCabang)

	query := buildQuery(&where, dateAwal, dateAkhir)

	rows, err := db.Query(query)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var results []entity.ResponseTransaksiResumeWidget
	for rows.Next() {
		var result entity.ResponseTransaksiResumeWidget
		err = rows.Scan(
			&result.Total_Noa,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
			&result.Total_Transaksi,
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

func handleDatesAndParams(data entity.ParamsTransaksi) (*string, *string, *string, *string) {
	awalDate, akhirDate := getMonthDates()

	wilayahCabang := handleParam("0", data.Wilayah_Cabang)
	dateAwal := handleParam(awalDate, data.Date_Awal)
	dateAkhir := handleParam(akhirDate, data.Date_Akhir)
	kodeUnit := handleParam("0", data.UnitKode)

	return wilayahCabang, dateAwal, dateAkhir, kodeUnit
}

func handleParam(defaultValue, paramValue string) *string {
	if paramValue == "0" {
		return &defaultValue
	}
	return &paramValue
}

func buildWhereClause(kodeUnit, wilayahCabang *string) string {
	var where string
	if kodeUnit != nil {
		where = fmt.Sprintf("KODE_CAB = '%s' AND ", *kodeUnit)
	}
	if wilayahCabang != nil {
		where = fmt.Sprintf("INISIAL_CAB_INDUK = '%s' AND ", *wilayahCabang)
	}
	return where
}

func buildQuery(where, dateAwal, dateAkhir *string) string {
	query := fmt.Sprintf(`SELECT
		COUNT(DISTINCT NO_REKENING) AS TOTAL_NOA,
		SUM(TRANSAKSI_POKOK) AS TRANSAKSI_POKOK,
		SUM(TRANSAKSI_BUNGA) AS TRANSAKSI_BUNGA,
		SUM(TRANSAKSI_DENDA) AS TRANSAKSI_DENDA,
		SUM(TRANSAKSI_POKOK + TRANSAKSI_BUNGA + TRANSAKSI_DENDA) AS TOTAL_TRANSAKSI
	FROM
		TblGetTransaksi_Fix
	WHERE
		%s
		TGLTRANS BETWEEN '%s' AND '%s'`, *where, *dateAwal, *dateAkhir)
	return query
}

func ListResumeTransaksi(c *gin.Context) {
	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateRealisasi(data); err != nil {
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
	var dateAwal *string
	var dateAkhir *string
	var kodeUnit *string

	if *&data.Wilayah_Cabang == "0" {
		wilayahCabang = nil
	} else {
		wilayahCabang = &data.Wilayah_Cabang
	}
	if *&data.Wilayah == "0" {
		wilayah = nil
	} else {
		wilayah = &data.Wilayah
	}
	if *&data.Date_Awal == "0" {
		dateAwal = nil
	} else {
		dateAwal = &data.Date_Awal
	}
	if *&data.Date_Akhir == "0" {
		dateAkhir = nil
	} else {
		dateAkhir = &data.Date_Akhir
	}
	if *&data.UnitKode == "0" {
		kodeUnit = nil
	} else {
		kodeUnit = &data.UnitKode
	}
	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH_CABANG", Value: wilayahCabang},
		{Name: "WILAYAH", Value: wilayah},
		{Name: "UNIT_KODE", Value: kodeUnit},
		{Name: "DATE_AWAL", Value: dateAwal},
		{Name: "DATE_AKHIR", Value: dateAkhir},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiTransaksiResume_Table", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseTransaksiTable
	for rows.Next() {
		var result entity.ResponseTransaksiTable
		err = rows.Scan(
			&result.Cabang,
			&result.Unit,
			&result.Wil,
			&result.Total_Noa,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
			&result.Total_Transaksi,
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

func buildWhereClauseTable(kodeUnit, wilayahCabang *string) string {
	var where string
	if kodeUnit != nil {
		where = fmt.Sprintf("T.KODE_CAB = '%s' AND ", *kodeUnit)
	}
	if wilayahCabang != nil {
		where = fmt.Sprintf("T.INISIAL_CAB_INDUK = '%s' AND ", *wilayahCabang)
	}
	return where
}

func buildQueryTable(where, dateAwal, dateAkhir *string) string {
	query := fmt.Sprintf(`SELECT
		C.nama AS CABANG,
		U.nama AS UNIT,
		COUNT(DISTINCT NO_REKENING) AS TOTAL_NOA,
		SUM(T.TRANSAKSI_POKOK) AS TRANSAKSI_POKOK,
		SUM(T.TRANSAKSI_BUNGA) AS TRANSAKSI_BUNGA,
		SUM(T.TRANSAKSI_DENDA) AS TRANSAKSI_DENDA,
		SUM(T.TRANSAKSI_POKOK + T.TRANSAKSI_BUNGA + T.TRANSAKSI_DENDA) AS TOTAL_TRANSAKSI
	FROM
		TblGetTransaksi_Fix AS T
	LEFT JOIN cabang AS C ON T.INISIAL_CAB_INDUK=C.inisial_cab
	LEFT JOIN unit AS U ON T.KODE_CAB=U.kode_unit
	WHERE
		%s
		TGLTRANS BETWEEN '%s' AND '%s'		
	GROUP BY C.nama, U.nama
	ORDER BY TOTAL_TRANSAKSI DESC`, *where, *dateAwal, *dateAkhir)
	return query
}

func ListResumeTransaksiCab(c *gin.Context) {
	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateWidget(data); err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Validation Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(http.StatusBadRequest, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()

	wilayahCabang, dateAwal, dateAkhir, kodeUnit := handleDatesAndParams(data)

	where := buildWhereClauseTable(kodeUnit, wilayahCabang)

	query := buildQueryTable(&where, dateAwal, dateAkhir)

	rows, err := db.Query(query)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseTransaksiTable
	for rows.Next() {
		var result entity.ResponseTransaksiTable
		err = rows.Scan(
			&result.Cabang,
			&result.Unit,
			&result.Total_Noa,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
			&result.Total_Transaksi,
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

func ListDetailTransaksi(c *gin.Context) {

	var data entity.ParamsTransaksi
	c.Bind(&data)

	if err := validateRealisasi(data); err != nil {
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
	var dateAwal *string
	var dateAkhir *string
	var kodeUnit *string

	wilayahCabang = nilIfEmptys(&data.Wilayah_Cabang)
	wilayah = nilIfEmptys(&data.Wilayah)
	dateAwal = nilIfEmptys(&data.Date_Awal)
	dateAkhir = nilIfEmptys(&data.Date_Akhir)
	kodeUnit = nilIfEmptys(&data.UnitKode)

	params := []entity.Param{
		{Name: "PAGE", Value: data.Page},
		{Name: "LIMIT", Value: data.Limit},
		{Name: "WILAYAH_CABANG", Value: wilayahCabang},
		{Name: "WILAYAH", Value: wilayah},
		{Name: "UNIT_KODE", Value: kodeUnit},
		{Name: "DATE_AWAL", Value: dateAwal},
		{Name: "DATE_AKHIR", Value: dateAkhir},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiTransaksiDetailTransaksi", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailTransaksi
	for rows.Next() {
		var result entity.ResponseDetailTransaksi
		err = rows.Scan(
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Cabang,
			&result.Unit,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
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

func DetailTransaksiById(c *gin.Context) {
	var data entity.ParamsTransaksi
	c.Bind(&data)

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()
	params := []entity.Param{
		{Name: "KODE", Value: data.Kode},
	}

	rows, err := service.ExecuteSP(db, "RMD_RealisasiTransaksiDetailTransaksi_Detail", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseDetailTransaksiById
	for rows.Next() {
		var result entity.ResponseDetailTransaksiById
		err = rows.Scan(
			&result.Kode_Unit,
			&result.TglTrans,
			&result.No_Rekening,
			&result.Nama_Nasabah,
			&result.Kode_Trans,
			&result.Transaksi_Pokok,
			&result.Transaksi_Bunga,
			&result.Transaksi_Denda,
			&result.Unit,
			&result.Id_Nasabah,
			&result.Tipe_Kredit,
			&result.Ft,
			&result.Angs_Ke,
			&result.Tgl_Angsuran,
			&result.Tgl_Realisasi,
			&result.Tgl_Jth_Tempo,
			&result.Kol,
			&result.Angs_Total,
			&result.Jml_Pinjaman,
			&result.Os,
			&result.Rek_Dca,
			&result.Saldo_Dca,
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

func validateRealisasi(data entity.ParamsTransaksi) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(data.Wilayah_Cabang, "WILAYAH CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(data.Wilayah_Cabang, 4, "WILAYAH CABANG"); err != nil {
		return err
	}
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate PAGE
	if err := validation.IntegerInRange(data.Page, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(data.Limit, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateWidget(data entity.ParamsTransaksi) error {

	// Validate INISIAL CABANG contains only alphabetic characters
	if err := validation.StringIsAlpha(data.Wilayah_Cabang, "WILAYAH CABANG"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(data.Wilayah_Cabang, 4, "WILAYAH CABANG"); err != nil {
		return err
	}
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(data.Wilayah, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	return nil
}

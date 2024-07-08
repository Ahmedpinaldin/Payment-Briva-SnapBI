package controller

import (
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"
	"github.com/gin-gonic/gin"
)

func ListMonitoringDCA(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	Type := c.Params.ByName("type")
	Search := c.Params.ByName("search")
	Row := c.Params.ByName("row")
	Page := c.Params.ByName("page")
	Sort := c.Params.ByName("sort")
	
	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	params := []entity.Param {
		{Name: "posisi", Value: Kode},
		{Name: "type", Value: Type},
		{Name: "search", Value: Search},
		{Name: "pageNo", Value: Page},
		{Name: "pageSize", Value: Row},
		{Name: "sortOrder", Value: Sort},
	}

	rows, err := service.ExecuteSP(db, "sp_getmondca", params); 
	if err != nil {
		res := helper.BuildErrorResponse(400,"Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer rows.Close()
	
	var resp entity.ResponseMonitoringDca
	for rows.Next() {
		var result entity.ResponseMonitoringDca
		err = rows.Scan(
			&result.JMLH,
			&result.Data,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400,"Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		resp.Data = result.Data
		resp.JMLH =result.JMLH
		// results = result
	}
	
	if err := rows.Err(); err != nil {
		res := helper.BuildErrorResponse(400,"Error creating", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(200, "Success", resp)
	c.AbortWithStatusJSON(http.StatusOK, res)

}


func ListLMPD(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	var data entity.ParamListLmpd
    c.Bind(&data)

	IdUnit :=  "'"+ Kode +"'"
	fromDate :=  "'"+ data.From_Date +"'"
	toDate :=  "'"+ data.To_Date +"'"

	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := 
	`SELECT
		c.no_memo,
		c.tgl_memo,
		a.nama_nasabah,
		a.no_rekening, 
		a.angsuran_total,
		a.angs_ke,
		a.saldo_dca,
		c.janji_pemenuhan,
		c.debet_dca,
		d.tgl_trans,
		d.setoran
	FROM
		nominatif_realtime a
	LEFT JOIN( SELECT nomer_rek, MAX(no) AS id FROM tbl_mppd 
	WHERE `+`kode_unit = `+IdUnit+ ` and try_convert(date,tgl_memo) BETWEEN CONVERT(date, `+fromDate+`) and CONVERT(date, `+toDate+`)
	GROUP BY nomer_rek )b ON a.no_rekening = b.nomer_rek
	LEFT JOIN tbl_mppd c ON c.nomer_rek = a.no_rekening AND b.id = c.no
	LEFT JOIN v_lmpd d on d.no_rek_dca = c.rek_dca
	WHERE `+`a.kode_unit = `+ IdUnit +` and try_convert(date,tgl_memo) BETWEEN CONVERT(date, `+fromDate+`) and CONVERT(date, `+toDate+`)`

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.DataLMPD
	
	for resp.Next() {
		var result entity.DataLMPD
		err = resp.Scan(
			&result.No_Memo,
			&result.Tgl_Memo,
			&result.Nama_Nasabah,
			&result.No_Rekening,
			&result.Angsuran_Total,
			&result.Angs_Ke,
			&result.Saldo_Dca,
			&result.Janji_Pertemuan,
			&result.Debet_DCA,
			&result.Tgl_Trans,
			&result.Setoran,
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

func ListMPPD(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	var data entity.ParamListMppd
    c.Bind(&data)
	
	IdUnit :=  "'"+ Kode +"'"
	noMemo :=  "'"+ data.No_Memo +"'"

	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := 
	`SELECT
		a.saldo_dca,
		a.debet_dca,
		b.nama_nasabah,
		a.alasan_debet,
		b.tipe_kredit,
		a.no_memo,
		a.tgl_memo,
		a.nomer_rek,
		a.janji_pemenuhan,
		a.angs_ke,
		a.angsuran,
		a.no as id
	FROM
		tbl_mppd a
	join nominatif_realtime b on a.nomer_rek = b.no_rekening
	WHERE
		a.no_memo = `+noMemo+`
	AND a.kode_unit = `+IdUnit

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.DataMPPD
	
	for resp.Next() {
		var result entity.DataMPPD
		err = resp.Scan(
			&result.Saldo_Dca,
			&result.Debet_DCA,
			&result.Nama_Nasabah,
			&result.Alasan_Debet,
			&result.Tipe_Kredit,
			&result.No_Memo,
			&result.Tgl_Memo,
			&result.Nomer_Rek,
			&result.Janji_Pemenuhan,
			&result.Angs_Ke,
			&result.Angsuran,
			&result.Id,
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


func InsertMPPD(c *gin.Context) {
	// form := &entity.InserMPPD{
	// 	Angsuran: c.Request.FormValue("angsuran"),
	// 	Debet_DCA: c.Request.FormValue("debet_dca"),
	// 	Sisa_DCA: c.Request.FormValue("sisa_dca"),
	// 	Kode_Unit: c.Request.FormValue("kode_unit"),
	// 	Kode_Cabang: c.Request.FormValue("kode_cabang"),
	// 	Bulan: c.Request.FormValue("bulan"),
	// 	Tahun: c.Request.FormValue("tahun"),
	// 	DateEntry_Time: c.Request.FormValue("date_time_entry"),
	// 	Date_Entry: c.Request.FormValue("date_entry"),
	// 	Saldo_DCA: c.Request.FormValue("saldo_dca"),
	// }
	// _, err := govalidator.ValidateStruct(form)
	// if err != nil {
	// 	res := helper.BuildErrorResponse(400,"Data Required", err.Error(), helper.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 	return
	// }

	// db, err := config.NewConnection()
	// if err != nil { 
	// 	res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 	return
	// }
	// defer db.Close()
	
	// stmt, err := db.Prepare("INSERT INTO tbl_mppd (no_rekening, tgl_deskcall, ket, stats_telp) VALUES (@no_rekening, @tgl_deskcall, @keterangan, @status_telp)")
    // if err != nil {	
	// 	res := helper.BuildErrorResponse(400,"Failed Prepare the SQL statement", err.Error(), helper.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 	return
    // }
    // defer stmt.Close()

    // // Execute the SQL statement with parameters
    // _, err = stmt.Exec(
	// 	sql.Named("no_rekening", form.No_Rekening), 
	// 	sql.Named("tgl_deskcall", form.Tgl_DeskCall), 
	// 	sql.Named("keterangan", form.Keterangan), 
	// 	sql.Named("status_telp", form.Status_Telp),
	// )
	
    // if err != nil {
	// 	res := helper.BuildErrorResponse(400,"Failed  Execute the SQL statement", err.Error(), helper.EmptyObj{})
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 	return
    // }
	// res := helper.BuildResponse(200, "Success", "Data Insert Success")
	// c.AbortWithStatusJSON(http.StatusOK, res)

}
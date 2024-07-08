package controller

import (
	"database/sql"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ListDeskCall(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	var data entity.ParamListLmpd
    c.Bind(&data)

	IdUnit :=  "'"+ Kode +"'"
	fromDate :=  "'"+ data.From_Date +"'"
	toDate :=  "'"+ data.To_Date +"'"

	// fmt.Println(IdUnit+"ASD") 
	// fmt.Println(fromDate) 
	// fmt.Println(toDate) 

	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := `
	SELECT 
		b.kode_unit, 
		b.inisial_cab, 
		CONVERT(DATE,a.tgl_deskcall) as tgl_deskcall, 
		CONVERT(DATE,c.tgl_jb) as tgl_jb, 
		a.stats_telp,b.no_hp, 
		a.ket, 
		b.nama_nasabah, 
		a.no_rekening, 
		b.tgl_jatuh_tempo, 
		b.tgl_angsuran, 
		b.angs_ke, 
		b.tipe_kredit, 
		b.angsuran_total, 
		d.nama_karyawan, 
		b.tgl_realisasi 
	FROM 
		history_deskcall a 
	LEFT JOIN 
		nominatif_realtime b ON a.no_rekening = b.no_rekening 
	LEFT JOIN 
		v_lastjb c ON a.no_rekening = c.no_rekening 
	LEFT JOIN 
		t_mapping_nasabah d ON a.no_rekening = d.no_rekening and d.status = 1 
	WHERE
		b.kode_unit = `+IdUnit+`AND b.nama_nasabah is not null AND a.tgl_deskcall BETWEEN`+fromDate+`AND`+toDate

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.ResponseListDeskCall
	
	for resp.Next() {
		var result entity.ResponseListDeskCall
		err = resp.Scan(
			&result.Kode_Unit,
			&result.Inisial_Cab,
			&result.Tgl_DeskCall,
			&result.Tgl_Jb,
			&result.Status_Telpon,
			&result.No_Hp,
			&result.Keterangan,
			&result.Nama_Nasabah,
			&result.No_Rekening,
			&result.Tgl_Jatuh_Tempo,
			&result.Tgl_Angsuran,
			&result.Angsuran_Ke,
			&result.Tipe_Kredit,
			&result.Angsuran_Total,
			&result.Nama_Karyawan,
			&result.Tgl_Realisasi,
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

func DetailDeskCall(c *gin.Context) {
	var data entity.ParamDeskCall
    c.Bind(&data)
	noRek :=  "'"+ data.No_Rekening +"'"

	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := `
		SELECT 
			a.nama_nasabah, 
			a.no_rekening, 
			a.jml_pinjaman, 
			a.angsuran_total, 
			a.angs_ke, 
			a.no_hp, 
			a.tipe_kredit, 
			CONVERT(VARCHAR(10),
			a.tgl_realisasi,105) AS tgl_realisasi, 
			b.nama_karyawan FROM nominatif_realtime a 
		LEFT JOIN 
			t_mapping_nasabah b ON a.no_rekening = b.no_rekening and b.status = 1 
		WHERE 
			b.no_rekening =` + noRek

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.ResponseDetailDeskCall
	
	for resp.Next() {
		var result entity.ResponseDetailDeskCall
		err = resp.Scan(
			&result.Nama_Nasabah,
			&result.No_Rekening,
			&result.Jml_Pimjaman,
			&result.Angsuran_Total,
			&result.Angsuran_Ke,
			&result.No_Hp,
			&result.Tipe_Kredit,
			&result.Tgl_Realisasi,
			&result.Nama_Karyawan,
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

func RiwayatDeskCall(c *gin.Context) {
	var data entity.ParamDeskCall
    c.Bind(&data)
	
	noRek :=  "'"+ data.No_Rekening +"'"
	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := "SELECT tgl_deskcall,  ket FROM  history_deskcall  WHERE  no_rekening = "+noRek+"order by tgl_deskcall desc"
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.ResponseRiwayatDeskCall
	
	for resp.Next() {
		var result entity.ResponseRiwayatDeskCall
		err = resp.Scan(
			&result.Tgl_DeskCall,
			&result.Keterangan,
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


func ListNasabahDeskCall(c *gin.Context) {
	kode := c.Params.ByName("kode")
	kodeUnit :=  "'"+ kode +"'"
	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	query := 
	`SELECT
		nb.NAMA_NASABAH,
		nb.NO_REKENING,
		nb.ft as dpd_clos,
		nb.angsuran_total as angsuran,
		nb.saldo_dca as dca,
		nb.jml_pinjaman as plafond,
		nb.saldo_nominatif as OUTSTANDING,
		nb.angs_ke as angsuran_ke,
		nb.no_hp,
		CONVERT(varchar,Convert(datetime,nb.tgl_angsuran),103) as temp_TGL_JTH_TEMPO,
		(select top 1 nama from 
			(SELECT no_rekening,nomor_induk from t_mapping_nasabah 
						where status = '1' and no_rekening = nb.NO_REKENING) a
			join [10.61.4.150].[HRIS].dbo.master_karyawan b on a.nomor_induk = b.nomor_induk and b.status_active = 'ACTIVE' ) as nama_aom,
		pn.cara_penanganan as plan_penanganan,
		ph.tgl_jb as tgl_jb,
		nb.tgl_angsuran,
		nb.tgl_jatuh_tempo as TGL_JTH_TEMPO
	FROM
		nominatif_realtime nb
	left JOIN 
		plan_bulan pn on nb.no_rekening = pn.no_rekening and pn.bln_plan = MONTH(GETDATE()) and pn.thn_plan = YEAR(GETDATE())
	LEFT JOIN 
		(
			SELECT max(tgl_jb) as tgl_jb,no_rekening from plan_hari where bln_plan = MONTH(GETDATE()) and thn_plan = YEAR(GETDATE())
			and kode_unit = `+kodeUnit+`
			GROUP BY no_rekening 
		) ph on ph.no_rekening = nb.no_rekening
	WHERE 
		nb.kode_unit = `+kodeUnit+`and CONVERT(date,case when nb.ft = 0 and CONVERT(date, nb.tgl_angsuran) < CONVERT(date, GETDATE()) then DATEADD(month, 1, nb.tgl_angsuran)
		else nb.tgl_angsuran end) between CONVERT(date, GETDATE()) AND CONVERT(date, DATEADD(day,+7, GETDATE()))`

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.ResponseListNasabahDeskCall
	
	for resp.Next() {
		var result entity.ResponseListNasabahDeskCall
		err = resp.Scan(
			&result.Nama_Nasabah,
			&result.No_Rekening,
			&result.Dpd_Clos,
			&result.Angsuran,
			&result.Dca,
			&result.Plafond,
			&result.OutStanding,
			&result.Angsuran_Ke,
			&result.No_Hp,
			&result.Temp_Tgl_Jatuh_Tempo,
			&result.Nama_Aom,
			&result.Plan_Penanganan,
			&result.Tgl_Jb,
			&result.Tgl_Angsuran,
			&result.Tgl_Jatuh_Tempo,
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

func InsertDeskCall(c *gin.Context) {
	form := &entity.InsertDeskCall{
		No_Rekening: c.Request.FormValue("no_rekening"),
		Status_Telp: c.Request.FormValue("status_telp"),
		Keterangan: c.Request.FormValue("keterangan"),
		Tgl_DeskCall: c.Request.FormValue("tgl_deskcall"),
	}
	_, err := govalidator.ValidateStruct(form)
	if err != nil {
		res := helper.BuildErrorResponse(400,"Data Required", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	db, err := config.NewConnection()
	if err != nil { 
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer db.Close()
	
	stmt, err := db.Prepare("INSERT INTO history_deskcall (no_rekening, tgl_deskcall, ket, stats_telp) VALUES (@no_rekening, @tgl_deskcall, @keterangan, @status_telp)")
    if err != nil {	
		res := helper.BuildErrorResponse(400,"Failed Prepare the SQL statement", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
    }
    defer stmt.Close()

    // Execute the SQL statement with parameters
    _, err = stmt.Exec(
		sql.Named("no_rekening", form.No_Rekening), 
		sql.Named("tgl_deskcall", form.Tgl_DeskCall), 
		sql.Named("keterangan", form.Keterangan), 
		sql.Named("status_telp", form.Status_Telp),
	)
	
    if err != nil {
		res := helper.BuildErrorResponse(400,"Failed  Execute the SQL statement", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
    }
	res := helper.BuildResponse(200, "Success", "Data Insert Success")
	c.AbortWithStatusJSON(http.StatusOK, res)

}



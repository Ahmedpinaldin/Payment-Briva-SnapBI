package controller

import (
	"database/sql"
	"fmt"

	"net/http"
	"strconv"

	"time"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/config/constvar"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/entity/response"
	"github.com/gin-gonic/gin"
)

// -----------------menampilkan kol 1-5 planbulan-----------------\\
func GetPlanBulan(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	kd_unit := c.Query("kode")

	limitInt, _ := strconv.ParseUint(limit, 10, 64)
	pageInt, _ := strconv.ParseUint(page, 10, 64)

	//hitung offset
	offset := (pageInt - 1) * limitInt

	kolektibilitas := c.Query("kol")
	switch kolektibilitas {
	case "1":
		kolektibilitas = constvar.KOL_1
	case "2":
		kolektibilitas = constvar.KOL_2
	case "3":
		kolektibilitas = constvar.KOL_3
	case "4":
		kolektibilitas = constvar.KOL_4
	case "5":
		kolektibilitas = constvar.KOL_5
	}

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// execute SQL statement
	query := `
		SELECT b.id, b.kode_unit, b.no_rekening, b.nama_nasabah, b.tipe_kredit, b.kolektibilitas, b.ft, b.angs_ke, b.saldo_nominatif, a.rencana_penyelesaian
		FROM dbo.plan_bulan a
		JOIN dbo.nominatif_baseon b ON a.no_rekening = b.no_rekening
		WHERE b.kolektibilitas = @kolektibilitas AND b.kode_unit = @kode_unit
		ORDER BY b.no_rekening
		OFFSET @offset ROWS FETCH NEXT @limit ROWS ONLY
	`

	rows, err := db.Query(query,
		sql.Named("kolektibilitas", kolektibilitas),
		sql.Named("kode_unit", kd_unit),
		sql.Named("offset", offset),
		sql.Named("limit", limitInt),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	//proses hasil query
	var planBulan []response.BulanResponse
	for rows.Next() {
		var plan response.BulanResponse
		// var tglJb sql.NullTime
		if err := rows.Scan(&plan.Id, &plan.Kode_Unit, &plan.No_Rekening, &plan.Nama_Nasabah, &plan.Tipe_Kredit, &plan.Kolektibilitas, &plan.Ft, &plan.Angs_Ke, &plan.Saldo_Nominatif, &plan.Rencana_Penyelesaian); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		planBulan = append(planBulan, plan)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  pageInt,
		"limit": limitInt,
		"data":  planBulan,
	})

}

// -----------------menampilkan recovery wo-----------------\\
func GetRecoveryWoById(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	kd_unit := c.Query("kode")

	limitInt, _ := strconv.ParseUint(limit, 10, 64)
	pageInt, _ := strconv.ParseUint(page, 10, 64)

	//hitung offset
	offset := (pageInt - 1) * limitInt

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	//execute SQL statement
	query := `
		SELECT N.id, N.kode_unit, N.no_rekening, N.nama_nasabah, N.tgl_wo, N.jml_wo, N.saldo_nominatif, P.rencana_penyelesaian
		FROM dbo.nominatif_wo N
		LEFT JOIN dbo.plan_bulan P ON N.no_rekening = P.no_rekening
		WHERE N.kode_unit = @kode_unit
		ORDER BY N.no_rekening
		OFFSET @offset ROWS FETCH NEXT @limit ROWS ONLY
	`
	rows, err := db.Query(query,
		sql.Named("kode_unit", kd_unit),
		sql.Named("offset", offset),
		sql.Named("limit", limitInt),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	//proses hasil query
	var PlanbulanWoByid []response.PlanBulanWo
	for rows.Next() {
		var plan response.PlanBulanWo
		if err := rows.Scan(&plan.Id, &plan.Kode_Unit, &plan.No_Rekening, &plan.Nama_Nasabah, &plan.Tgl_Wo, &plan.Jml_wo, &plan.Saldo_Nominatif, &plan.Rencana_Penyelesaian); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		PlanbulanWoByid = append(PlanbulanWoByid, plan)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  pageInt,
		"limit": limitInt,
		"data":  PlanbulanWoByid,
	})

}

// -----------------detail planbulan-----------------\\
func GetPlanBulanById(c *gin.Context) {
	id := c.Param("id")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
        SELECT 
            N.no_rekening, N.kode_unit, N.inisial_cab, N.nama_unit, N.nama_nasabah, N.tipe_kredit, 
            N.tgl_realisasi, N.tgl_jatuh_tempo, CONVERT(VARCHAR(6), N.tgl_jatuh_tempo, 112) AS bln_jatuh_tempo, 
            N.tgl_angsuran AS tgl_angs, N.jml_pinjaman, N.kolektibilitas, N.saldo_nominatif, N.ft, 
            P.jenis_usaha, P.kondisi_usaha, P.sp_1, P.sp_2, P.sp_3, P.obyek_agunan, P.dokumen_agunan, 
            P.pengikatan, P.jenis_pengikatan, P.status_dokumen, P.nilai_saat_ini, P.mudah_jual, 
            P.keberadaan_agunan, P.sengketa, P.s_diputuskan_kpp, P.status_lelang, P.lelang_ke, 
            P.permasalahan, R.kategori_recovery, R.ket_recovery
        FROM nominatif_baseon AS N
        LEFT JOIN plan_bulan AS P ON N.no_rekening = P.no_rekening
        LEFT JOIN recovery AS R ON P.no_rekening = R.no_rekening AND MONTH(GETDATE()) = MONTH(R.bln) AND YEAR(GETDATE()) = YEAR(R.bln)
        WHERE N.id = @id
        ORDER BY P.id DESC
        OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY
    `

	rows, err := db.Query(query, sql.Named("id", id))
	fmt.Println(rows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var result response.PlanBulanById
	if rows.Next() {
		// var result response.PlanBulanById
		err := rows.Scan(
			&result.No_Rekening, &result.Kode_Unit, &result.Inisial_Cab, &result.Nama_Unit,
			&result.Nama_Nasabah, &result.Tipe_Kredit, &result.Tgl_Realisasi, &result.Tgl_Jatuh_Tempo,
			&result.Bln_Jatuh_Tempo, &result.Tgl_Angs, &result.Jml_Pinjaman, &result.Kolektibilitas,
			&result.Saldo_Nominatif, &result.Ft, &result.Jenis_Usaha, &result.Kondisi_Usaha,
			&result.Sp_1, &result.Sp_2, &result.Sp_3, &result.Obyek_Agunan, &result.Dokumen_Agunan,
			&result.Pengikatan, &result.Jenis_Pengikatan, &result.Status_Dokumen, &result.Nilai_Saat_Ini,
			&result.Mudah_Jual, &result.Keberadaan_Agunan, &result.Sengketa, &result.S_Diputuskan_Kpp,
			&result.Status_Lelang, &result.Lelang_Ke, &result.Permasalahan, &result.Kategori_Recovery,
			&result.Ket_Recovery,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// -----------------insert planbulan (save)-----------------\\
func CreatePlanBulan(c *gin.Context) {
	//mengambil data yang diterima dari request
	var plan entity.PlanBulan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	//koneksi ke db
	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	defer db.Close()

	//memulai transaksi
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}

	//execute sql statement
	_, err = tx.Exec("INSERT INTO dbo.plan_bulan(tgl_plan, bln_plan, thn_plan, no_rekening, nama_nasabah, kode_unit, inisial_cab, jenis_usaha, kondisi_usaha, sp_1, sp_2, sp_3, obyek_agunan, dokumen_agunan, pengikatan, jenis_pengikatan, status_dokumen, nilai_saat_ini, mudah_jual, keberadaan_agunan, sengketa, s_diputuskan_kpp, status_lelang, lelang_ke, permasalahan, cara_penanganan, rencana_penyelesaian, input_pokok, nominal_lunas_bertahap, tgl_jb, ket_rinci, id_pic) VALUES (@tgl_plan, @bln_plan, @thn_plan, @no_rekening, @nama_nasabah, @kode_unit, @inisial_cab, @jenis_usaha, @kondisi_usaha, @sp_1, @sp_2, @sp_3, @obyek_agunan, @dokumen_agunan, @pengikatan, @jenis_pengikatan, @status_dokumen, @nilai_saat_ini, @mudah_jual, @keberadaan_agunan, @sengketa, @s_diputuskan_kpp, @status_lelang, @lelang_ke, @permasalahan, @cara_penanganan, @rencana_penyelesaian, @input_pokok, @nominal_lunas_bertahap, @tgl_jb, @ket_rinci, @id_pic)",
		sql.Named("tgl_plan", plan.Tgl_Plan),
		sql.Named("bln_plan", plan.Bln_Plan),
		sql.Named("thn_plan", plan.Thn_Plan),
		sql.Named("no_rekening", plan.No_Rekening),
		sql.Named("nama_nasabah", plan.Nama_Nasabah),
		sql.Named("kode_unit", plan.Kode_Unit),
		sql.Named("inisial_cab", plan.Inisial_Cab),
		sql.Named("jenis_usaha", plan.Jenis_Usaha),
		sql.Named("kondisi_usaha", plan.Kondisi_Usaha),
		sql.Named("sp_1", plan.Sp_1),
		sql.Named("sp_2", plan.Sp_2),
		sql.Named("sp_3", plan.Sp_3),
		sql.Named("obyek_agunan", plan.Obyek_Agunan),
		sql.Named("dokumen_agunan", plan.Dokumen_Agunan),
		sql.Named("pengikatan", plan.Pengikatan),
		sql.Named("jenis_pengikatan", plan.Jenis_Pengikatan),
		sql.Named("status_dokumen", plan.Status_Dokumen),
		sql.Named("nilai_saat_ini", plan.Nilai_Saat_Ini),
		sql.Named("mudah_jual", plan.Mudah_Jual),
		sql.Named("keberadaan_agunan", plan.Keberadaan_Agunan),
		sql.Named("sengketa", plan.Sengketa),
		sql.Named("s_diputuskan_kpp", plan.S_Diputuskan_Kpp),
		sql.Named("status_lelang", plan.Status_Lelang),
		sql.Named("lelang_ke", plan.Lelang_Ke),
		sql.Named("permasalahan", plan.Permasalahan),
		sql.Named("cara_penanganan", plan.Cara_Penanganan),
		sql.Named("rencana_penyelesaian", plan.Rencana_Penyelesaian),
		sql.Named("input_pokok", plan.Input_Pokok),
		sql.Named("nominal_lunas_bertahap", plan.Nominal_Lunas_Bertahap),
		sql.Named("tgl_jb", plan.Tgl_Jb),
		sql.Named("ket_rinci", plan.Ket_Rinci),
		sql.Named("id_pic", plan.Id_Pic),
	)

	if err != nil {
		//Rollback transaksi jika terjadi kesalahan
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error4": err.Error()})
		return
	}

	//commit transaksi jika tidak terjadi kesalahan
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error5": err.Error()})
		return
	}

	// Jika komit berhasil, masukkan data recovery
	dataRecovery := map[string]interface{}{
		"tgl":               time.Now().Format("2006-01-02"),
		"bln":               time.Now().Format("01"),
		"thn":               time.Now().Format("2006"),
		"no_rekening":       plan.No_Rekening,
		"nama_nasabah":      plan.Nama_Nasabah,
		"kode_unit":         plan.Kode_Unit,
		"inisial_cab":       plan.Inisial_Cab,
		"kategori_recovery": plan.Kategori_Recovery,
		"ket_recovery":      plan.Ket_Recovery,
	}

	_, err = db.Exec(`
    INSERT INTO recovery (tgl, bln, thn, no_rekening, nama_nasabah, kode_unit, inisial_cab, kategori_recovery, ket_recovery)
    VALUES (@tgl, @bln, @thn, @no_rekening, @nama_nasabah, @kode_unit, @inisial_cab, @kategori_recovery, @ket_recovery)
`,
		sql.Named("tgl", dataRecovery["tgl"]),
		sql.Named("bln", dataRecovery["bln"]),
		sql.Named("thn", dataRecovery["thn"]),
		sql.Named("no_rekening", dataRecovery["no_rekening"]),
		sql.Named("nama_nasabah", dataRecovery["nama_nasabah"]),
		sql.Named("kode_unit", dataRecovery["kode_unit"]),
		sql.Named("inisial_cab", dataRecovery["inisial_cab"]),
		sql.Named("kategori_recovery", dataRecovery["kategori_recovery"]),
		sql.Named("ket_recovery", dataRecovery["ket_recovery"]),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "berhasil membuat PlanBulan"})

}

// -----------------insert planbulan wo (save)-----------------\\
func CreatePlanBulanWo(c *gin.Context) {
	var plan entity.PlanBulanWo
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	//memulai transaksi
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//execute sql statement
	_, err = tx.Exec("INSERT INTO dbo.plan_bulan_wo(tgl_plan, bln_plan, thn_plan, no_rekening, kode_unit, inisial_cab, lunas_blm_lunas, nilai_jamkrindo, sp_1, sp_2, sp_3, jaminan_ada_tidak, obyek_agunan, dokumen_agunan, pengikatan, jenis_pengikatan, status_dokumen, nilai_saat_ini, mudah_jual, keberadaan_agunan, sengketa, s_diputuskan_kpp, status_lelang, lelang_ke, permasalahan, cara_penanganan, rencana_penyelesaian, nominal_pembayaran, tgl_jb, ket_rinci, id_pic) VALUES (@tgl_plan, @bln_plan, @thn_plan, @no_rekening, @kode_unit, @inisial_cab, @lunas_blm_lunas, @nilai_jamkrindo, @sp_1, @sp_2, @sp_3, @jaminan_ada_tidak, @obyek_agunan, @dokumen_agunan, @pengikatan, @jenis_pengikatan, @status_dokumen, @nilai_saat_ini, @mudah_jual, @keberadaan_agunan, @sengketa, @s_diputuskan_kpp, @status_lelang, @lelang_ke, @permasalahan, @cara_penanganan, @rencana_penyelesaian, @nominal_pembayaran, @tgl_jb, @ket_rinci, @id_pic)",
		sql.Named("tgl_plan", plan.Tgl_Plan),
		sql.Named("bln_plan", plan.Bln_Plan),
		sql.Named("thn_plan", plan.Thn_Plan),
		sql.Named("no_rekening", plan.No_Rekening),
		sql.Named("kode_unit", plan.Kode_Unit),
		sql.Named("inisial_cab", plan.Inisial_Cab),
		sql.Named("lunas_blm_lunas", plan.Lunas_Blm_Lunas),
		sql.Named("nilai_jamkrindo", plan.Nilai_Jamkrindo),
		sql.Named("sp_1", plan.Sp_1),
		sql.Named("sp_2", plan.Sp_2),
		sql.Named("sp_3", plan.Sp_3),
		sql.Named("jaminan_ada_tidak", plan.Jaminan_Ada_Tidak),
		sql.Named("obyek_agunan", plan.Obyek_Agunan),
		sql.Named("dokumen_agunan", plan.Dokumen_Agunan),
		sql.Named("pengikatan", plan.Pengikatan),
		sql.Named("jenis_pengikatan", plan.Jenis_Pengikatan),
		sql.Named("status_dokumen", plan.Status_Dokumen),
		sql.Named("nilai_saat_ini", plan.Nilai_Saat_Ini),
		sql.Named("mudah_jual", plan.Mudah_Jual),
		sql.Named("keberadaan_agunan", plan.Keberadaan_Agunan),
		sql.Named("sengketa", plan.Sengketa),
		sql.Named("s_diputuskan_kpp", plan.S_Diputuskan_Kpp),
		sql.Named("status_lelang", plan.Status_Lelang),
		sql.Named("lelang_ke", plan.Lelang_Ke),
		sql.Named("permasalahan", plan.Permasalahan),
		sql.Named("cara_penanganan", plan.Cara_Penanganan),
		sql.Named("rencana_penyelesaian", plan.Rencana_Penyelesaian),
		sql.Named("tgl_jb", plan.Tgl_Jb),
		sql.Named("ket_rinci", plan.Ket_Rinci),
		sql.Named("nominal_pembayaran", plan.Nominal_Pembayaran),
		sql.Named("id_pic", plan.Id_Pic),
	)

	if err != nil {
		//Rollback transaksi jika terjadi kesalahan
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error4": err.Error()})
		return
	}

	//commit transaksi jika tidak terjadi kesalahan
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error5": err.Error()})
		return
	}

	// Jika komit berhasil, masukkan data recovery
	dataRecovery := map[string]interface{}{
		"tgl":               time.Now().Format("2006-01-02"),
		"bln":               time.Now().Format("01"),
		"thn":               time.Now().Format("2006"),
		"no_rekening":       plan.No_Rekening,
		"nama_nasabah":      plan.Nama_Nasabah,
		"kode_unit":         plan.Kode_Unit,
		"inisial_cab":       plan.Inisial_Cab,
		"kategori_recovery": plan.Kategori_Recovery,
		"ket_recovery":      plan.Ket_Recovery,
	}

	_, err = db.Exec(`
    INSERT INTO recovery_wo (tgl, bln, thn, no_rekening, nama_nasabah, kode_unit, inisial_cab, kategori_recovery, ket_recovery)
    VALUES (@tgl, @bln, @thn, @no_rekening, @nama_nasabah, @kode_unit, @inisial_cab, @kategori_recovery, @ket_recovery)
`,
		sql.Named("tgl", dataRecovery["tgl"]),
		sql.Named("bln", dataRecovery["bln"]),
		sql.Named("thn", dataRecovery["thn"]),
		sql.Named("no_rekening", dataRecovery["no_rekening"]),
		sql.Named("nama_nasabah", dataRecovery["nama_nasabah"]),
		sql.Named("kode_unit", dataRecovery["kode_unit"]),
		sql.Named("inisial_cab", dataRecovery["inisial_cab"]),
		sql.Named("kategori_recovery", dataRecovery["kategori_recovery"]),
		sql.Named("ket_recovery", dataRecovery["ket_recovery"]),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "berhasil membuat PlanBulanWO"})

}

func UpdatePlanBulan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plan_bulan ID"})
		return
	}

	// Ambil data yang diterima dari request
	var plan entity.PlanBulan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query SQL update
	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Persiapkan nilai tgl_jb jika tidak kosong
	var tglJB sql.NullTime
	if plan.Tgl_Jb != "" {
		tglJB.Time, err = time.Parse("2006-01-02", plan.Tgl_Jb) // Ubah format tanggal sesuai kebutuhan
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tglJB.Valid = true
	}

	// Execute SQL statement
	_, err = db.Exec(`
		UPDATE dbo.plan_bulan 
		SET 
			tgl_plan = @tgl_plan, 
			bln_plan = @bln_plan, 
			thn_plan = @thn_plan, 
			no_rekening = @no_rekening, 
			nama_nasabah = @nama_nasabah, 
			kode_unit = @kode_unit, 
			inisial_cab = @inisial_cab, 
			jenis_usaha = @jenis_usaha, 
			kondisi_usaha = @kondisi_usaha, 
			sp_1 = @sp_1, 
			sp_2 = @sp_2, 
			sp_3 = @sp_3, 
			obyek_agunan = @obyek_agunan, 
			dokumen_agunan = @dokumen_agunan, 
			pengikatan = @pengikatan, 
			jenis_pengikatan = @jenis_pengikatan, 
			status_dokumen = @status_dokumen, 
			nilai_saat_ini = @nilai_saat_ini, 
			mudah_jual = @mudah_jual, 
			keberadaan_agunan = @keberadaan_agunan, 
			sengketa = @sengketa, 
			s_diputuskan_kpp = @s_diputuskan_kpp, 
			status_lelang = @status_lelang, 
			lelang_ke = @lelang_ke, 
			permasalahan = @permasalahan, 
			cara_penanganan = @cara_penanganan, 
			rencana_penyelesaian = @rencana_penyelesaian, 
			input_pokok = @input_pokok, 
			nominal_lunas_bertahap = @nominal_lunas_bertahap, 
			tgl_jb = @tgl_jb, 
			ket_rinci = @ket_rinci, 
			id_pic = @id_pic 
		WHERE id = @id`,
		sql.Named("tgl_plan", plan.Tgl_Plan),
		sql.Named("bln_plan", plan.Bln_Plan),
		sql.Named("thn_plan", plan.Thn_Plan),
		sql.Named("no_rekening", plan.No_Rekening),
		sql.Named("nama_nasabah", plan.Nama_Nasabah),
		sql.Named("kode_unit", plan.Kode_Unit),
		sql.Named("inisial_cab", plan.Inisial_Cab),
		sql.Named("jenis_usaha", plan.Jenis_Usaha),
		sql.Named("kondisi_usaha", plan.Kondisi_Usaha),
		sql.Named("sp_1", plan.Sp_1),
		sql.Named("sp_2", plan.Sp_2),
		sql.Named("sp_3", plan.Sp_3),
		sql.Named("obyek_agunan", plan.Obyek_Agunan),
		sql.Named("dokumen_agunan", plan.Dokumen_Agunan),
		sql.Named("pengikatan", plan.Pengikatan),
		sql.Named("jenis_pengikatan", plan.Jenis_Pengikatan),
		sql.Named("status_dokumen", plan.Status_Dokumen),
		sql.Named("nilai_saat_ini", plan.Nilai_Saat_Ini),
		sql.Named("mudah_jual", plan.Mudah_Jual),
		sql.Named("keberadaan_agunan", plan.Keberadaan_Agunan),
		sql.Named("sengketa", plan.Sengketa),
		sql.Named("s_diputuskan_kpp", plan.S_Diputuskan_Kpp),
		sql.Named("status_lelang", plan.Status_Lelang),
		sql.Named("lelang_ke", plan.Lelang_Ke),
		sql.Named("permasalahan", plan.Permasalahan),
		sql.Named("cara_penanganan", plan.Cara_Penanganan),
		sql.Named("rencana_penyelesaian", plan.Rencana_Penyelesaian),
		sql.Named("input_pokok", plan.Input_Pokok),
		sql.Named("nominal_lunas_bertahap", plan.Nominal_Lunas_Bertahap),
		sql.Named("tgl_jb", plan.Tgl_Jb),
		sql.Named("ket_rinci", plan.Ket_Rinci),
		sql.Named("id_pic", plan.Id_Pic),
		sql.Named("id", id),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PlanBulan berhasil diupdate"})
}

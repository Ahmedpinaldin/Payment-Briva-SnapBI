package controller

import (
	"database/sql"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/config/constvar"
	"example.com/rest-api-recoll-mobile/entity/response"
	"github.com/gin-gonic/gin"
)

// Resume Baseon
func ResumeBaseonKol(c *gin.Context) {
	kodeUnit := c.Query("kode")

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

	//excute SQL
	query := `
		SELECT
			COUNT(id) AS noa_baseon,
			SUM(saldo_nominatif) AS os_baseon
		FROM nominatif_baseon
		WHERE kode_unit = @kode_unit AND kolektibilitas = @kolektibilitas
	`
	rows, err := db.Query(query,
		sql.Named("kode_unit", kodeUnit),
		sql.Named("kolektibilitas", kolektibilitas),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	//proses hasil query
	var noaBaseon int
	var osBaseon sql.NullFloat64
	if rows.Next() {
		err := rows.Scan(&noaBaseon, &osBaseon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//Mengembalikan data dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"noa_baseon": noaBaseon,
		"os_baseon":  osBaseon.Float64,
	})
}

func ResumePlanKol1(c *gin.Context) {
	kodeUnit := c.Query("kode") //tabel planbulan -> kode_unit

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Execute SQL
	query := `
        SELECT
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 1-30' THEN P.id END) AS noa_ttp_pd_bucket_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 1-30' THEN N.saldo_nominatif END, 0)) AS os_ttp_pd_bucket_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id END) AS noa_back_to_current,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif END, 0)) AS os_back_to_current,
            COUNT(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN P.id END) AS noa_3r,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN N.saldo_nominatif END, 0)) AS os_3r,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id END) AS noa_pelunasan_sekaligus,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif END, 0)) AS os_pelunasan_sekaligus,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN N.saldo_nominatif END, 0)) AS os_lelang_terjual,
            COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke Bucket  31-60' THEN P.id END) AS noa_rembes_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Rembes ke Bucket  31-60' THEN N.saldo_nominatif END, 0)) AS os_rembes_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke Bucket  1-30' THEN P.id END) AS noa_rembes_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Rembes ke Bucket  1-30' THEN N.saldo_nominatif END, 0)) AS os_rembes_1_30
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit=@kode_unit AND bln_plan=MONTH(GETDATE()) AND thn_plan=YEAR(GETDATE()) AND N.kolektibilitas ='L'
    `
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var ResumeKol1 []response.ResumePlanKol1
	if rows.Next() {
		var kol1 response.ResumePlanKol1
		err := rows.Scan(
			&kol1.NoaTtpPdBucket130, &kol1.OsTtpPdBucket130,
			&kol1.NoaBackToCurrent, &kol1.OsBackToCurrent,
			&kol1.Noa3r, &kol1.Os3r,
			&kol1.NoaPelunasanSekaligus, &kol1.OsPelunasanSekaligus,
			&kol1.NoaLelangTerjual, &kol1.OsLelangTerjual,
			&kol1.NoaRembes3160, &kol1.OsRembes3160,
			&kol1.NoaRembes130, &kol1.OsRembes130,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ResumeKol1 = append(ResumeKol1, kol1)
	}

	// Return data in JSON format
	c.JSON(http.StatusOK, ResumeKol1)
}

// Resume planbulan kol 2
func ResumePlanKol2(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
        SELECT 
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 61-90' THEN P.id END) AS noa_ttp_pd_bucket_61_90,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 61-90' THEN N.saldo_nominatif END, 0)) AS os_ttp_pd_bucket_61_90,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 1-30' THEN P.id END) AS noa_ttp_pd_bucket_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 1-30' THEN N.saldo_nominatif END, 0)) AS os_ttp_pd_bucket_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 31-60' THEN P.id END) AS noa_ttp_pd_bucket_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Tetap pd Bucket 31-60' THEN N.saldo_nominatif END, 0)) AS os_ttp_pd_bucket_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN P.id END) AS noa_back_to_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN N.saldo_nominatif END, 0)) AS os_back_to_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN P.id END) AS noa_back_to_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN N.saldo_nominatif END, 0)) AS os_back_to_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id END) AS noa_back_to_current,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif END, 0)) AS os_back_to_current,
            COUNT(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN P.id END) AS noa_3r,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN N.saldo_nominatif END, 0)) AS os_3r,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id END) AS noa_pelunasan_sekaligus,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif END, 0)) AS os_pelunasan_sekaligus,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN N.saldo_nominatif END, 0)) AS os_pelunasan_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.nominal_lunas_bertahap END, 0)) AS os_nominal_lunas_bertahap,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.id END) AS noa_pelunasan_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN N.saldo_nominatif END, 0)) AS os_lelang_terjual,
            COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke 61-90' THEN P.id END) AS noa_rembes_61_90,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Rembes ke 61-90' THEN N.saldo_nominatif END, 0)) AS os_rembes_61_90,
            COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN P.id END) AS noa_rembes_npl,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN N.saldo_nominatif END, 0)) AS os_rembes_npl
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening = N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE()) AND N.kolektibilitas = 'PK'
    `
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var resumeKol2 []response.ResumePlanKol2
	if rows.Next() {
		var kol2 response.ResumePlanKol2
		err := rows.Scan(
			&kol2.NoaTtpPdBucket6190, &kol2.OsTtpPdBucket6190,
			&kol2.NoaTtpPdBucket130, &kol2.OsTtpPdBucket130,
			&kol2.NoaTtpPdBucket3160, &kol2.OsTtpPdBucket3160,
			&kol2.NoaBackTo3160, &kol2.OsBackTo3160,
			&kol2.NoaBackTo130, &kol2.OsBackTo130,
			&kol2.NoaBackToCurrent, &kol2.OsBackToCurrent,
			&kol2.Noa3r, &kol2.Os3r,
			&kol2.NoaPelunasanSekaligus, &kol2.OsPelunasanSekaligus,
			&kol2.NoaLelangTerjual,
			&kol2.NoaPelunasanBertahap, &kol2.OsPelunasanBertahap, &kol2.OsNominalLunasBertahap,
			&kol2.OsLelangTerjual,
			&kol2.NoaRembes6190, &kol2.OsRembes6190,
			&kol2.NoaRembesNPL, &kol2.OsRembesNPL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeKol2 = append(resumeKol2, kol2)
	}

	c.JSON(http.StatusOK, resumeKol2)
}

// resume planbulan kol 3
func ResumePlanKol3(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
        SELECT
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN P.id END) AS noa_back_to_61_90,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN N.saldo_nominatif END, 0)) AS os_back_to_61_90,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN P.id END) AS noa_back_to_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN N.saldo_nominatif END, 0)) AS os_back_to_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN P.id END) AS noa_back_to_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN N.saldo_nominatif END, 0)) AS os_back_to_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id END) AS noa_back_to_current,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif END, 0)) AS os_back_to_current,
            COUNT(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN P.id END) AS noa_3r,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN N.saldo_nominatif END, 0)) AS os_3r,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id END) AS noa_lunas_sekaligus,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif END, 0)) AS os_lunas_sekaligus,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.id END) AS noa_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN N.saldo_nominatif END, 0)) AS os_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.nominal_lunas_bertahap END, 0)) AS nominal_lunas_bertahap,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN N.saldo_nominatif END, 0)) AS os_lelang_terjual,
            COUNT(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN P.id END) AS noa_stay_npl,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN N.saldo_nominatif END, 0)) AS os_stay_npl
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE()) AND N.kolektibilitas = 'KL'
	`

	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var resumeKol3 []response.ResumePlanKol34
	if rows.Next() {
		var kol3 response.ResumePlanKol34
		err := rows.Scan(
			&kol3.NoaBackTo6190, &kol3.OsBackTo6190,
			&kol3.NoaBackTo3160, &kol3.OsBackTo3160,
			&kol3.NoaBackTo130, &kol3.OsBackTo130,
			&kol3.NoaBackToCurrent, &kol3.OsBackToCurrent,
			&kol3.Noa3r, &kol3.Os3r,
			&kol3.NoaLunasSekaligus, &kol3.OsLunasSekaligus,
			&kol3.NoaLunasBertahap, &kol3.OsLunasBertahap,
			&kol3.NominalLunasBertahap,
			&kol3.NoaLelangTerjual, &kol3.OsLelangTerjual,
			&kol3.NoaStayNPL, &kol3.OsStayNPL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeKol3 = append(resumeKol3, kol3)
	}
	c.JSON(http.StatusOK, resumeKol3)
}

// resume planbulan kol 4
func ResumePlanKol4(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
        SELECT
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN P.id END) AS noa_back_to_61_90,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN N.saldo_nominatif END, 0)) AS os_back_to_61_90,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN P.id END) AS noa_back_to_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN N.saldo_nominatif END, 0)) AS os_back_to_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN P.id END) AS noa_back_to_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN N.saldo_nominatif END, 0)) AS os_back_to_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id END) AS noa_back_to_current,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif END, 0)) AS os_back_to_current,
            COUNT(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN P.id END) AS noa_3r,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN N.saldo_nominatif END, 0)) AS os_3r,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id END) AS noa_lunas_sekaligus,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif END, 0)) AS os_lunas_sekaligus,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.id END) AS noa_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN N.saldo_nominatif END, 0)) AS os_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.nominal_lunas_bertahap END, 0)) AS nominal_lunas_bertahap,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN N.saldo_nominatif END, 0)) AS os_lelang_terjual,
            COUNT(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN P.id END) AS noa_stay_npl,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN N.saldo_nominatif END, 0)) AS os_stay_npl
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE()) AND N.kolektibilitas = 'D'
	`
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var resumeKol4 []response.ResumePlanKol34
	if rows.Next() {
		var kol4 response.ResumePlanKol34
		err := rows.Scan(
			&kol4.NoaBackTo6190, &kol4.OsBackTo6190,
			&kol4.NoaBackTo3160, &kol4.OsBackTo3160,
			&kol4.NoaBackTo130, &kol4.OsBackTo130,
			&kol4.NoaBackToCurrent, &kol4.OsBackToCurrent,
			&kol4.Noa3r, &kol4.Os3r,
			&kol4.NoaLunasSekaligus, &kol4.OsLunasSekaligus,
			&kol4.NoaLunasBertahap, &kol4.OsLunasBertahap,
			&kol4.NominalLunasBertahap,
			&kol4.NoaLelangTerjual, &kol4.OsLelangTerjual,
			&kol4.NoaStayNPL, &kol4.OsStayNPL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeKol4 = append(resumeKol4, kol4)
	}
	c.JSON(http.StatusOK, resumeKol4)
}

// resume planbulan kol 5
func ResumePlanKol5(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
        SELECT
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN P.id END) AS noa_back_to_61_90,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 61-90' THEN N.saldo_nominatif END, 0)) AS os_back_to_61_90,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN P.id END) AS noa_back_to_31_60,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 31-60' THEN N.saldo_nominatif END, 0)) AS os_back_to_31_60,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN P.id END) AS noa_back_to_1_30,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to 1-30' THEN N.saldo_nominatif END, 0)) AS os_back_to_1_30,
            COUNT(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN P.id END) AS noa_back_to_current,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Penagihan - Back to Current' THEN N.saldo_nominatif END, 0)) AS os_back_to_current,
			COUNT(CASE WHEN P.rencana_penyelesaian='Input Pokok' THEN P.id END) AS noa_input_pokok,
        	SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Input Pokok' THEN P.input_pokok END, 0)) AS os_input_pokok,
            COUNT(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN P.id END) AS noa_3r,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Restruktur (3-R)' THEN N.saldo_nominatif END, 0)) AS os_3r,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN P.id END) AS noa_lunas_sekaligus,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sekaligus' THEN N.saldo_nominatif END, 0)) AS os_lunas_sekaligus,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.id END) AS noa_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN N.saldo_nominatif END, 0)) AS os_lunas_bertahap,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' THEN P.nominal_lunas_bertahap END, 0)) AS nominal_lunas_bertahap,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN N.saldo_nominatif END, 0)) AS os_lelang_terjual,
            COUNT(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN P.id END) AS noa_stay_npl,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='STAY NPL' THEN N.saldo_nominatif END, 0)) AS os_stay_npl
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE()) AND N.kolektibilitas = 'M'
	`
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var resumeKol5 []response.ResumePlanKol5
	if rows.Next() {
		var kol5 response.ResumePlanKol5
		err := rows.Scan(
			&kol5.NoaBackTo6190, &kol5.OsBackTo6190,
			&kol5.NoaBackTo3160, &kol5.OsBackTo3160,
			&kol5.NoaBackTo130, &kol5.OsBackTo130,
			&kol5.NoaBackToCurrent, &kol5.OsBackToCurrent,
			&kol5.NoaInputPokok, &kol5.OsInputPokok,
			&kol5.Noa3r, &kol5.Os3r,
			&kol5.NoaLunasSekaligus, &kol5.OsLunasSekaligus,
			&kol5.NoaLunasBertahap, &kol5.OsLunasBertahap,
			&kol5.NominalLunasBertahap,
			&kol5.NoaLelangTerjual, &kol5.OsLelangTerjual,
			&kol5.NoaStayNPL, &kol5.OsStayNPL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeKol5 = append(resumeKol5, kol5)
	}
	c.JSON(http.StatusOK, resumeKol5)
}

func ResumeBaseonWo(c *gin.Context) {
	kodeUnit := c.Query("kode")
	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	//excute SQL
	query := `
		SELECT
			COUNT(id) AS noa_baseon,
			SUM(saldo_nominatif) AS os_baseon
		FROM nominatif_wo
		WHERE kode_unit = @kode_unit 
	`
	rows, err := db.Query(query,
		sql.Named("kode_unit", kodeUnit),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	//proses hasil query
	var noaBaseon int
	var osBaseon sql.NullFloat64
	if rows.Next() {
		err := rows.Scan(&noaBaseon, &osBaseon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//Mengembalikan data dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"noa_baseon": noaBaseon,
		"os_baseon":  osBaseon.Float64,
	})

}

func ResumePlanWo(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
		SELECT
            COUNT(CASE WHEN P.rencana_penyelesaian='Pickup Angsuran' THEN P.id END) AS noa_pickup_angsuran,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pickup Angsuran' THEN P.nominal_pembayaran END, 0)) AS os_pickup_angsuran,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Sesuai OS' THEN P.id END) AS noa_pelunasan_sesuai_os,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Sesuai OS' THEN P.nominal_pembayaran END, 0)) AS os_pelunasan_sesuai_os,
            COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Dibawah OS' THEN P.id END) AS noa_pelunasan_dibawah_os,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Dibawah OS' THEN P.nominal_pembayaran END, 0)) AS os_pelunasan_dibawah_os,
            COUNT(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.id END) AS noa_lelang_terjual,
            SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Lelang (terjual bulan ini)' THEN P.nominal_pembayaran END, 0)) AS os_lelang_terjual
        FROM plan_bulan_wo AS P
        LEFT JOIN nominatif_wo AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE())
	`
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	var resumeWo []response.ResumeWo
	if rows.Next() {
		var wo response.ResumeWo
		err := rows.Scan(
			&wo.NoaPickupAngsuran, &wo.OsPickupAngsuran,
			&wo.NoaPelunasanSesuaiOS, &wo.OsPelunasanSesuaiOS,
			&wo.NoaPelunasanDibawahOS, &wo.OsPelunasanDibawahOS,
			&wo.NoaLelangTerjual, &wo.OsLelangTerjual,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeWo = append(resumeWo, wo)
	}
	c.JSON(http.StatusOK, resumeWo)
}

func ResumeRembesNpl(c *gin.Context) {
	kodeUnit := c.Query("kode")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
		SELECT
		COUNT(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN P.id END) AS noa_rembes_npl,
        SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Rembes ke NPL' THEN N.saldo_nominatif END, 0)) AS os_rembes_npl,
        COUNT(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' AND N.ft >= 61 AND N.kolektibilitas='PK' THEN P.id END) AS noa_lunas_bertahap,
        SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' AND N.ft >= 61 AND N.kolektibilitas='PK' THEN N.saldo_nominatif END, 0)) AS os_lunas_bertahap,
        SUM(ISNULL(CASE WHEN P.rencana_penyelesaian='Pelunasan Bertahap' AND N.ft >= 61 AND N.kolektibilitas='PK' THEN P.nominal_lunas_bertahap END, 0)) AS nominal_lunas_bertahap
        FROM plan_bulan AS P
        LEFT JOIN nominatif_baseon AS N ON P.no_rekening=N.no_rekening
        WHERE P.kode_unit = @kode_unit AND bln_plan = MONTH(GETDATE()) AND thn_plan = YEAR(GETDATE())
	`
	rows, err := db.Query(query, sql.Named("kode_unit", kodeUnit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	var resumeRembesNpl []response.ResumeRembesNpl
	if rows.Next() {
		var rembes response.ResumeRembesNpl
		err := rows.Scan(
			&rembes.NoaRembesNpl, &rembes.OsRembesNpl,
			&rembes.NoaLunasBertahap, &rembes.OsLunasBertahap,
			&rembes.NominalLunasBertahap,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resumeRembesNpl = append(resumeRembesNpl, rembes)
	}
	c.JSON(http.StatusOK, resumeRembesNpl)

}

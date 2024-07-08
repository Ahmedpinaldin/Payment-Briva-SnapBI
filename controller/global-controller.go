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
)

func ListCabang(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	if err := validateCabang(Kode); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := "SELECT C.id_cabang, C.inisial_cab, C.nama, C.alamat FROM cabang AS C LEFT JOIN wilayah_rmd AS W ON C.inisial_cab = W.inisial_cab"

	if Kode != "0" {
		query += " WHERE W.wilayah_rmd = " + IdKode + " OR W.wilayah_rmd IS NULL"
	}
	query += " ORDER BY C.nama ASC"

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseCabang

	for resp.Next() {
		var result entity.ResponseCabang
		err = resp.Scan(
			&result.Id_Cabang,
			&result.Inisial_Cab,
			&result.Nama,
			&result.Alamat,
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

func ListUnitByCabang(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	if err := validateKodeCabang(Kode); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT kode_unit, nama_unit, inisial_unit FROM t_mapping_unit WHERE kode_cabang = " + IdKode + " ORDER BY nama_unit ASC;"
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseUnit

	for resp.Next() {
		var result entity.ResponseUnit
		err = resp.Scan(
			&result.Kode_Unit,
			&result.Nama_Unit,
			&result.Inisial_Unit,
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

func ListWilayah(c *gin.Context) {

	// Get Cabang
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT DISTINCT wilayah_rmd, keterangan FROM wilayah_rmd"

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseWilayah

	for resp.Next() {
		var result entity.ResponseWilayah
		err = resp.Scan(
			&result.Wilayah_Rmd,
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

func ListCaraPenanganan(c *gin.Context) {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// SELECT cara_penanganan FROM st_plan_hari_wo WHERE cara_penanganan !='UPDATE-AGUNAN' AND  cara_penanganan IS NOT NULL
	query := "SELECT cara_penanganan FROM st_plan WHERE cara_penanganan !='REMBES' AND  cara_penanganan IS NOT NULL"

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseCaraPenanganan

	for resp.Next() {
		var result entity.ResponseCaraPenanganan
		err = resp.Scan(
			&result.Cara_Penanganan,
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

func ListCaraPenangananWo(c *gin.Context) {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//
	query := "SELECT cara_penanganan FROM st_plan_hari_wo WHERE cara_penanganan !='UPDATE-AGUNAN' AND  cara_penanganan IS NOT NULL "

	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseCaraPenanganan

	for resp.Next() {
		var result entity.ResponseCaraPenanganan
		err = resp.Scan(
			&result.Cara_Penanganan,
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

func ListRencanaPenyelesaian(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT rencana_penyelesaian FROM st_rencana_penyelesaian_kol_5 WHERE key_rencana_penyelesaian =" + IdKode
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseRencananPenyelesaian

	for resp.Next() {
		var result entity.ResponseRencananPenyelesaian
		err = resp.Scan(
			&result.Rencana_Penyelesaian,
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

func ListRencanaPenyelesaianWo(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT rencana_penyelesaian FROM st_rencana_penyelesaian_hari_wo WHERE key_rencana_penyelesaian =" + IdKode
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseRencananPenyelesaian

	for resp.Next() {
		var result entity.ResponseRencananPenyelesaian
		err = resp.Scan(
			&result.Rencana_Penyelesaian,
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
func ListPicPenanganan(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	if err := validateKodeCabang(Kode); err != nil {
		res := helper.BuildErrorResponse(400, "Validation Error", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT id_pic, inisial_pic, nama FROM pic_penanganan WHERE status = '1' AND inisial_cab = " + IdKode
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponsePic

	for resp.Next() {
		var result entity.ResponsePic
		err = resp.Scan(
			&result.Id_Pic,
			&result.Inisial_Pic,
			&result.Nama,
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

func DetailPic(c *gin.Context) {

	// Get Cabang
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"

	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT P.id_pic, P.inisial_pic, P.nama, J.jabatan, P.no_hp  FROM pic_penanganan AS P LEFT JOIN pic_jabatan AS J ON P.id_jabatan=J.id_jabatan WHERE P.id_pic =  " + IdKode
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseDetailPic

	for resp.Next() {
		var result entity.ResponseDetailPic
		err = resp.Scan(
			&result.Id_Pic,
			&result.Inisial_Pic,
			&result.Nama,
			&result.Jabatan,
			&result.No_Hp,
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

func ListCabangApproval(c *gin.Context) {

	var data entity.ParamsApprovalCabang
	c.Bind(&data)

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()

	var wilayahRmd *string

	if data.Wil_Rmd == "0" {
		wilayahRmd = nil
	} else {
		wilayahRmd = &data.Wil_Rmd
	}
	params := []entity.Param{
		{Name: "WIL_RMD", Value: wilayahRmd},
	}

	rows, err := service.ExecuteSP(db, "List_Cabang_Approval", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var results []entity.ResponseCabangApproval
	for rows.Next() {
		var result entity.ResponseCabangApproval
		err = rows.Scan(
			&result.Inisial_Cab,
			&result.Nama,
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

func ListPenyelesaian(c *gin.Context) {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT * FROM st_status_penyelesaian"
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponsePenyelesaian

	for resp.Next() {
		var result entity.ResponsePenyelesaian
		err = resp.Scan(
			&result.Id_Status,
			&result.Status_Penyelesaian,
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

func ListHasilPenyelesaian(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	IdKode := "'" + Kode + "'"
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := "SELECT * FROM st_hasil_penyelesaian WHERE id_status = " + IdKode
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
	}
	defer resp.Close()

	var results []entity.ResponseHasilPenyelesaian

	for resp.Next() {
		var result entity.ResponseHasilPenyelesaian
		err = resp.Scan(
			&result.Id_Hasil,
			&result.Id_Status,
			&result.Hasil_Penyelesaian,
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

func validateCabang(kode string) error {
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate Kode
	if err := validation.StringInSet(kode, allowedValues, "Kode"); err != nil {
		return err
	}

	return nil
}

func validateKodeCabang(kode string) error {
	if err := validation.StringIsAlpha(kode, "KODE"); err != nil {
		return err
	}

	if err := validation.StringMaxLength(kode, 4, "KODE"); err != nil {
		return err
	}
	return nil
}

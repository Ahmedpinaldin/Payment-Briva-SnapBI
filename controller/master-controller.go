package controller

import (
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"
	"github.com/gin-gonic/gin"
)
 
func ListNasabah(c *gin.Context) {
	Kode := c.Params.ByName("kode")
	var data entity.ParamsNasabah
    c.Bind(&data)

	// IdUnit :=  "'"+ Kode +"'"
	// fromDate :=  "'"+ data.From_Date +"'"
	// toDate :=  "'"+ data.To_Date +"'"

	db, err := config.NewConnection()
	if err != nil {
		res := helper.BuildErrorResponse(400,"Failed Connectin to Database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	defer db.Close()
	
	var wilayah *string
	if (*data.WIL_RMD == "0"){
		wilayah = nil
	}else{
		wilayah = data.WIL_RMD 
	}
	
	params := []entity.Param{
		{Name: "PAGE", Value: data.PAGE},
		{Name: "LIMIT", Value:data.LIMIT},
		{Name: "TYPE", Value: data.TYPE},
		{Name: "GROUP_ROLE", Value: data.GROUP_ROLE},
		{Name: "ROLE", Value: data.ROLE},
		{Name: "WIL_RMD", Value: wilayah},
		{Name: "INISIAL_CAB",Value:data.INISIAL_CAB},
		{Name: "KODE_UNIT", Value:Kode},
		{Name: "BUCKET", Value: data.BUCKET},
		{Name: "SEARCH", Value:data.SEARCH},
	}
	println(params)

	rows, err := service.ExecuteSP(db, "M_Nasabah", params)
	if err != nil {
		res := helper.BuildErrorResponse(400,"Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer rows.Close()

	var results []entity.ResponseNasabah
	for rows.Next() {
		var result entity.ResponseNasabah
		err = rows.Scan(
			&result.ID,
			&result.Nama_Nasabah,
			&result.Cabang,
			&result.No_Rekening,
			&result.Kolektibilitas,
			&result.Nama_Nasabah,
			&result.Tipe_Kredit,
			&result.Ft,
			&result.Angs_Ke,
			&result.Jml_Pinjaman,
			&result.Saldo_Nominatif,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400,"Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		results = append(results, result)
	}
	res := helper.BuildResponse(200, "Success", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
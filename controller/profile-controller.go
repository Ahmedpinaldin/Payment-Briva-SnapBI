package controller

import (
	"log"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"github.com/gin-gonic/gin"
) 

func Profile(c *gin.Context) {   

	// Get Cabang
	IdQuery := c.Params.ByName("id") 
	Id :=  "'"+ IdQuery +"'"
	
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT unit.nama AS nama_unit, unit.kode_unit AS kode_unit, unit.inisial AS inisial_unit, unit.alamat AS alamat_unit, cabang.inisial_cab AS inisial_cabang, cabang.nama AS nama_cabang, cabang.alamat AS alamat_cabang FROM unit LEFT JOIN cabang ON unit.inisial_cab = cabang.inisial_cab  WHERE unit.kode_unit = " + Id
	resp, err := db.Query(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    }
    defer resp.Close() 

	var results []entity.ResponseProfile
	
	for resp.Next() {
		var result entity.ResponseProfile
		err = resp.Scan(
			&result.Nama_Unit,
			&result.Kode_Unit,
			&result.Inisial,
			&result.Alamat_Unit,
			&result.Inisial_Cabang,
			&result.Nama_Cabang,
			&result.Alamat_Cabang,
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
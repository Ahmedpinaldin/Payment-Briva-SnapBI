package controller

import (
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"
	"example.com/rest-api-recoll-mobile/validation"
	"github.com/gin-gonic/gin"
)

func ListPenurunanNFL(c *gin.Context) {
	var data entity.ParamsPenururnanNPL
	c.Bind(&data)

	if err := validateDashboardInputParams(data); err != nil {
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
	var wilayahPtr *string
	var page *string
	var limit *string

	if data.WILAYAH == "0" {
		wilayahPtr = nil
	} else {
		wilayahPtr = &data.WILAYAH
	}
	if data.PAGE == "0" {
		page = nil
	} else {
		page = &data.PAGE
	}
	if data.LIMIT == "0" {
		limit = nil
	} else {
		limit = &data.LIMIT
	}
	defer db.Close()
	params := []entity.Param{
		{Name: "PAGE", Value: page},
		{Name: "LIMIT", Value: limit},
		{Name: "WILAYAH", Value: wilayahPtr},
	}

	rows, err := service.ExecuteSP(db, "Dash_RMD_ProgressNPL_List", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer rows.Close()

	var results []entity.ResponsePenurunanNPL
	for rows.Next() {
		var result entity.ResponsePenurunanNPL
		err = rows.Scan(
			&result.Wil,
			&result.Cabang,
			&result.Inisial_Cab,
			&result.Baki_Debet,
			&result.Baseon_Npl,
			&result.Npl_Saat_Ini,
			&result.Os_Target,
			&result.RMBS,
			&result.Sisa_Potensi,
			&result.Col_Npl,
			&result.Col_Btb,
			&result.Col_Btc,
			&result.Phase_Out,
			&result.Restruktur,
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

func GrandTotal(c *gin.Context) {
	var data entity.ParamsPenururnanNPL
	c.Bind(&data)

	if err := validateGrandTotal(data); err != nil {
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

	var wilayahPtr *string

	if data.WILAYAH == "0" {
		wilayahPtr = nil
	} else {
		wilayahPtr = &data.WILAYAH
	}

	params := []entity.Param{
		{Name: "WILAYAH", Value: wilayahPtr},
	}

	// Use parameterized queries/prepared statements
	rows, err := service.ExecuteSP(db, "Dash_RMD_ProgressNPL_Total", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error executing stored procedure", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer rows.Close()

	var results []entity.ResponseGrandTotal
	for rows.Next() {
		var result entity.ResponseGrandTotal
		err = rows.Scan(
			&result.Baki_debet,
			&result.Baki_debet_baseon,
			&result.Baseon_npl,
			&result.Os_Target,
			&result.Npl_saat_ini,
			&result.Rmbsn_sudah_flow,
			&result.Sisa_potensi_npl,
			&result.Rencana_penyelesaian,
			&result.Col_npl,
			&result.Col_btb_btc,
			&result.Phase_out,
			&result.Restruktur,
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
func validateDashboardInputParams(data entity.ParamsPenururnanNPL) error {
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(data.WILAYAH, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	// Validate PAGE
	if err := validation.IntegerInRange(data.PAGE, 0, 100, "PAGE"); err != nil {
		return err
	}

	// Validate LIMIT
	if err := validation.IntegerInRange(data.LIMIT, 0, 100, "LIMIT"); err != nil {
		return err
	}

	return nil
}

func validateGrandTotal(data entity.ParamsPenururnanNPL) error {
	allowedValues := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true}

	// Validate WILAYAH
	if err := validation.StringInSet(data.WILAYAH, allowedValues, "WILAYAH"); err != nil {
		return err
	}

	return nil
}

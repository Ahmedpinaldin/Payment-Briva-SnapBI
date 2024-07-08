package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity/response"

	// "example.com/rest-api-recoll-mobile/controller"
	"example.com/rest-api-recoll-mobile/entity"
	"github.com/gin-gonic/gin"
)

// PIC PENANGANAN
func ListPicPenangananCab(c *gin.Context) {
	insl_cab := c.Query("inisialcab") //filter

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id_pic, inisial_pic, nama FROM pic_penanganan WHERE inisial_cab='%s' AND status=1", insl_cab)

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// var results []entity.ListPenanganan
	var result []entity.ListPenangananCab
	for rows.Next() {
		// var id_pic int
		// var inisial_pic, nama *string
		var results entity.ListPenangananCab
		err := rows.Scan(&results.Id_Pic, &results.Inisial_Pic, &results.Nama)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result = append(result, results)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetPicPenangananId(c *gin.Context) {
	id := c.Param("id")

	db, err := config.NewConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	query := `
		SELECT P.id_pic, P.inisial_pic, P.nama, J.jabatan, P.no_hp
		FROM pic_penanganan AS P
		LEFT JOIN pic_jabatan AS J ON P.id_jabatan=J.id_jabatan
		WHERE P.id_pic = @id
	`

	rows, err := db.Query(query, sql.Named("id", id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []response.GetPenangananId
	for rows.Next() {
		var result response.GetPenangananId
		err := rows.Scan(&result.Id_Pic, &result.Inisial_Pic, &result.Nama, &result.Jabatan, &result.No_Hp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

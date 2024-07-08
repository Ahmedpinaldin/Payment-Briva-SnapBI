package controller

import (
	"database/sql"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/helper"
	"github.com/gin-gonic/gin"
)

func VersionCheck(c *gin.Context) {
	currentVersion := c.Param("version")

	db, err := config.NewConnection()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	latestVersion, err := getLatestVersion(db)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Failed to retrieve the latest version from the database", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deviceVersion := currentVersion

	if deviceVersion < latestVersion {
		c.JSON(http.StatusOK, gin.H{
			"message":        "Please download the update",
			"latestVersion":  latestVersion,
			"yourVersion":    currentVersion,
			"updateRequired": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":        "Your device is up to date",
			"latestVersion":  latestVersion,
			"yourVersion":    currentVersion,
			"updateRequired": false,
		})
	}
}

func getLatestVersion(db *sql.DB) (string, error) {
	query := "SELECT TOP 1 version FROM mobile_version ORDER BY date DESC"
	var version string

	err := db.QueryRow(query).Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}

package controller

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"example.com/rest-api-recoll-mobile/auth"
	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/entity"
	"example.com/rest-api-recoll-mobile/helper"
	"example.com/rest-api-recoll-mobile/service"

	// "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("event", "event"))
	return nil
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

const maxLoginAttempts = 5
const lockoutDuration = 15 * time.Minute

func Login(c *gin.Context) {
	form := &entity.LoginDTO{
		Username: c.Request.FormValue("username"),
		Password: c.Request.FormValue("password"),
	}

	if err := validate.Struct(form); err != nil {
		// Prepare error messages
		var errorMessages []string

		// Loop through validation errors
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.StructField())
			tag := err.Tag()

			if len(errorMessages) < 1 {
				switch tag {
				case "required":
					errorMessages = append(errorMessages, fmt.Sprintf("%s tidak boleh kosong", field))
				case "min":
					errorMessages = append(errorMessages, fmt.Sprintf("%s tidak boleh kurang dari %s karakter", field, err.Param()))
				case "max":
					errorMessages = append(errorMessages, fmt.Sprintf("%s tidak boleh lebih dari %s karakter", field, err.Param()))
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("%s is not valid", field))
				}
			}
		}

		// Create a response with custom error messages
		res := gin.H{
			"status":  http.StatusBadRequest,
			"message": "Login Gagal",
			"error":   errorMessages,
			"data":    gin.H{},
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	url := fmt.Sprintf("http://10.61.3.37/WebService/SSO_Mobile/login.php?user=%s&pass=%s&app_code=REMED", form.Username, form.Password)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("event", "event"))

	resp, err := client.Do(req)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error sending  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error reading  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var response entity.LoginResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error parsing response body as JSON:", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if isAccountLocked(form.Username) {
		res := helper.BuildErrorResponse(400, "Akun Terkunci Sementara", "Coba lagi dalam 15 menit", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	auths := response.Login[0].Response

	if auths != "TRUE" {
		err := increaseFailedAttempts(form.Username)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Login Gagal", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res := helper.BuildResponse(400, "Username / Password Salah", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data := response.Login[0].Data[0]
	params := []entity.Param{
		{Name: "USERNAME", Value: form.Username},
		{Name: "password", Value: form.Password},
		{Name: "kode_cabang", Value: data.KodeCabang},
		{Name: "kode_unit", Value: data.KodeUnit},
		{Name: "posisi_nama", Value: data.PosisiNama},
		// {Name: "foto", Value: data.Foto},
	}
	rows, err := service.ExecuteSP(db, "sp_checklogin", params)
	if err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  HTTP request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	defer rows.Close()

	var results []entity.Response
	for rows.Next() {
		var result entity.Response
		err = rows.Scan(
			&result.User_ID,
			&result.Username,
			&result.Password,
			&result.Status,
			&result.Id_Level,
			&result.Inisial_Cab,
			&result.Kode_Unit,
			&result.Add_by,
			&result.Date_add,
			&result.Time_Login,
			&result.Group_Level,
			&result.Nama_Unit,
			&result.Map_Unit,
			&result.Level,
			&result.Wilayah_RMD,
			&result.Wilayah,
			&result.Bucket,
			&result.Unit,
			&result.Cabang,
			&result.Alternate,
		)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		jwtWrapper := auth.JwtWrapper{
			SecretKey:       "verysecretkey",
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateToken(form.Username)
		if err != nil {
			res := helper.BuildErrorResponse(400, "Error creating Token", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		result.Foto = &data.Foto
		result.Name = &data.Nama
		result.Unit = &data.Unit
		result.Kode_Cabang = &data.KodeCabang
		result.Lokasi_Kerja = &data.LokasiKerja
		result.Nomor_Induk = &data.NomorInduk
		result.Token = &signedToken
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		res := helper.BuildErrorResponse(400, "Error creating  scan row", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deleteFailedAttempts(form.Username)
	res := helper.BuildResponse(200, "Login Berhasil", results)
	c.AbortWithStatusJSON(http.StatusOK, res)
	return
}

func isAccountLocked(username string) bool {
	// Perform a database query to check if the account is locked
	query := `
		SELECT is_locked, lockout_time, getdate() as dtnow
		FROM login_attempts
		WHERE username = @username
		ORDER BY attempt_time DESC
		OFFSET 0 ROWS
		FETCH NEXT 1 ROWS ONLY`

	var isLocked bool
	var lockoutTime time.Time
	var dtnow time.Time

	db, err := config.NewConnection()
	if err != nil {
		log.Println("Error creating a database connection:", err)
		return false
	}
	defer db.Close()

	// Query for the account lock status
	var dbErr error
	dbErr = db.QueryRow(query, sql.Named("username", username)).Scan(&isLocked, &lockoutTime, &dtnow)
	if dbErr != nil {
		if dbErr == sql.ErrNoRows {
			// No login attempts for the user; the account is not locked
			return false
		}
		// Log the database error
		log.Println("Error checking account lock status:", dbErr)
		return false
	}

	// Check if lockout_time is greater than 15 minutes=
	// fifteenMinutesAgo := lockoutTime.Add(15 * time.Minute)

	difference := dtnow.Sub(lockoutTime)
	// Check if the account is locked based on the lockout_time
	if difference.Minutes() < 0 {
		// The account is locked
		return true
	}

	if difference.Minutes() > 0 {
		// If lockout_time is older than 15 minutes, delete the row where username matches
		deleteQuery := "DELETE FROM login_attempts WHERE username = @username"
		_, deleteErr := db.Exec(deleteQuery, sql.Named("username", username))
		if deleteErr != nil {
			// Handle the delete error if necessary
			log.Println("Error deleting the row:", deleteErr)
			return false
		}
	}

	return false
}

func deleteFailedAttempts(username string) error {
	// Perform a database update to reset failed login attempts
	db, err := config.NewConnection()
	if err != nil {
		log.Println("Error creating a database connection:", err)
		return nil
	}
	deleteQuery := "DELETE FROM login_attempts WHERE username = @username"
	_, deleteErr := db.Exec(deleteQuery, sql.Named("username", username))
	if deleteErr != nil {
		// Handle the delete error if necessary
		log.Println("Error deleting the row:", deleteErr)
		return nil
	}
	return nil
}
func increaseFailedAttempts(username string) error {
	// Define the SQL query to UPSERT the login_attempts table
	query := `
		MERGE INTO login_attempts AS target
		USING (SELECT @username AS username) AS source
		ON target.username = source.username
		WHEN MATCHED THEN
			UPDATE SET
				failed_attempts = target.failed_attempts + 1,
				attempt_time = GETDATE(),
				is_locked = CASE WHEN target.failed_attempts + 1 >= 5 THEN 1 ELSE target.is_locked END,
				lockout_time = CASE WHEN target.failed_attempts + 1 >= 5 THEN DATEADD(MINUTE, 15, GETDATE()) ELSE target.lockout_time END
		WHEN NOT MATCHED THEN
			INSERT (username, failed_attempts, attempt_time, is_locked, lockout_time) VALUES (source.username, 1, GETDATE(), 0, NULL);`

	db, err := config.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, dbErr := db.Exec(query, sql.Named("username", username)) // Use sql.Named to provide the parameter value
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func logSuccessfulLogin(username string) error {
	// Perform a database insert to log a successful login attempt
	query := "INSERT INTO login_attempts (username, attempt_time, is_locked, failed_attempts) VALUES (?, NOW(), FALSE, 0)"

	db, err := config.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, dbErr := db.Exec(query, username)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

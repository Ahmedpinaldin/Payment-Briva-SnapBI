package service

import (
	"database/sql"
	"fmt"
	"strings"

	"example.com/rest-api-recoll-mobile/entity"
)

func ExecuteSP(db *sql.DB, spName string, params []entity.Param) (*sql.Rows, error) {
	fmt.Println(spName)
	fmt.Println(params)
	// create the SQL query string with named parameters
	query := fmt.Sprintf("EXEC %s ", spName)
	placeholders := make([]string, len(params))
	for i, p := range params {
		placeholders[i] = fmt.Sprintf("@%s", p.Name)
		fmt.Println(p.Name)
	}
	query += strings.Join(placeholders, ",")

	// create a prepared statement for the query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// execute the prepared statement with the parameters
	args := make([]interface{}, len(params))
	for i, p := range params {
		args[i] = sql.Named(p.Name, p.Value)
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

package Database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "bookstoredb"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "bookstoredb"
)

type DbInfo struct {
	DBConn *sql.DB
	Query  string
}

func ConnectDatabase() (*sql.DB, error) {
	// tz := dbM.GetTimeZone() // fetch from env

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	// close database

	// check db
	err = conn.Ping()
	if err != nil {
		defer conn.Close()
		return nil, err
	}
	fmt.Println("Connected!")
	return conn, err

}

// func (dbi *DbInfo) fetchall() {

// }

func (dbI *DbInfo) CreateTable(newTable string, model string) (bool, error) {

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS  %s %s", newTable, model)

	_, err := dbI.DBConn.Exec(query)
	if err != nil {
		return false, err
	}
	return true, err
}
func (dbI *DbInfo) InsertRow(tableName string, row map[string]interface{}) error {
	columns := make([]string, 0, len(row))
	values := make([]interface{}, 0, len(row))
	placeholders := make([]string, 0, len(row))
	fmt.Println(row)
	for column, value := range row {
		columns = append(columns, column)
		values = append(values, value)
		placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
	}

	query := "INSERT INTO " + tableName + " (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")"
	fmt.Println(query)
	_, err := dbI.DBConn.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}
func (dbI *DbInfo) FetchRows(values ...interface{}) ([]map[string]interface{}, error) {
	rows, err := dbI.DBConn.Query(dbI.Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	values_list := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range values_list {
		pointers[i] = &values_list[i]
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, column := range columns {
			val := values_list[i]
			if b, ok := val.([]byte); ok {
				row[column] = string(b)
			} else {
				row[column] = val
			}
		}

		results = append(results, row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
func (dbI *DbInfo) Query_close() {
	dbI.Query = ""
}

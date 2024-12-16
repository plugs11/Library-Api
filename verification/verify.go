package verification

import (
	"database/sql"

	"manan.tola/config"
)

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func CheckduplicateId(id string) bool {
	var count int
	query := "SELECT COUNT(*) FROM book WHERE id = ?"
	row := db.QueryRow(query, id)
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

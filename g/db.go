package g

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Open the db connection, panic if it fails.
func OpenDB(dbConfig DbConfig) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

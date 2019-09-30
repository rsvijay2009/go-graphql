package go_graphql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // Sql driver package
)

// DB establish database connection
var DB *sql.DB

// ConnectDB define the database connectivity
func ConnectDB() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err)
	}

	DB = db
}

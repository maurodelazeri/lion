package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// PostgresDB ...
var PostgresDB *sql.DB

func init() {
	InitEngine()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	var err error
	var (
		host     = os.Getenv("PSQL_HOST")
		port     = os.Getenv("PSQL_PORT")
		user     = os.Getenv("PSQL_USER")
		password = os.Getenv("PSQL_PASS")
		dbname   = os.Getenv("PSQL_DB")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	PostgresDB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Postgres - Problem with database connection ", err)
	}
}

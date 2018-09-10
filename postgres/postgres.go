package postgres

import (
	"log"
	"os"

	"github.com/oleiade/lane"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB ...
var (
	DB      *sqlx.DB
	DBQueue *lane.Queue
)

func init() {
	InitEngine()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	var err error
	DB, err = sqlx.Connect("postgres", "host="+os.Getenv("PSQL_HOST")+" user="+os.Getenv("PSQL_USER")+" password="+os.Getenv("PSQL_PASS")+" dbname="+os.Getenv("PSQL_DB")+" sslmode=disable")
	if err != nil {
		log.Fatal("Problem with database connection", err)
	}
}

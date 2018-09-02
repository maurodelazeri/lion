package postgres

import (
	"log"
	"os"

	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"

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
	DBQueue = lane.NewQueue()
	InitQueue()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	var err error
	DB, err = sqlx.Connect("postgres", "host="+os.Getenv("PSQL_HOST")+" user="+os.Getenv("PSQL_USER")+" password="+os.Getenv("PSQL_PASS")+" dbname="+os.Getenv("PSQL_DB")+" sslmode=disable")
	if err != nil {
		log.Fatal("Problem with database connection", err)
	}
}

// InitQueue to update the database, the operations are in a queue to guarantee the correct execution order
func InitQueue() {
	// Let's handle the clients asynchronously
	go func() {
		for {
			for DBQueue.Head() != nil {
				item := DBQueue.Dequeue()
				Worker(item)
			}
		}
	}()
}

// Worker execute ACID transations based on the received array of strings
func Worker(item interface{}) {
	switch t := item.(type) {
	case []string:
		tx := DB.MustBegin()
		for _, value := range t {
			tx.MustExec(value)
		}
		tx.Commit()
	default:
		logrus.Error("Data is not an array ", t)
	}
}

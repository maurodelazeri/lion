package postgres

import (
	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"
)

// DBQueue ...
var DBQueue *lane.Queue

func init() {
	DBQueue = lane.NewQueue()
	InitQueue()
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
		tx := PostgresDB.MustBegin()
		for _, value := range t {
			tx.MustExec(value)
		}
		tx.Commit()
	default:
		logrus.Error("Data is not an array ", t)
	}
}

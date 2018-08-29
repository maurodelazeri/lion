package postgres

import (
	"fmt"
	"os"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

// PostgresClient Session
var PostgresClient *xorm.Engine

// GetPostgressSession start redis
func GetPostgressSession() *xorm.Engine {
	if PostgresClient == nil {
		logrus.Info("Starting postgres session!")
		Initialize()
		return PostgresClient
	}
	return PostgresClient
}

// Initialize a postgres instance
func Initialize() {
	dbstring := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s  sslmode=disable",
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DB"), os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"))
	DB, err := xorm.NewEngine("postgres", dbstring)
	if err != nil {
		logrus.Error("postgres", err)
		os.Exit(1)
	}
	DB.SetMapper(core.GonicMapper{})
	PostgresClient = DB
}

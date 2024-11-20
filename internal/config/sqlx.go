package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewDatabase(config *viper.Viper, log *logrus.Logger) *sqlx.DB {
	driverName := "mysql"

	user := config.GetString("database.user")
	password := config.GetString("database.password")
	dbName, dbHost, dbPort := config.GetString("database.name"), config.GetString("database.dbHost"), config.GetInt("database.port")
	mydsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, password, dbHost, dbPort, dbName)

	db, err := sqlx.Connect(driverName, mydsn)
	if err != nil {
		log.Fatalf("Cannot connect into database : %v", err)
	}

	return db
}

package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetUpDatabase() *sql.DB {
	logrus.Info("Connecting to MySQL Database...")

	db, err := sql.Open("mysql", viper.GetString("DSN"))
	if err != nil {
		logrus.Panic(err)
	}

	db.SetMaxOpenConns(0)    //default: 0
	db.SetMaxIdleConns(2)    //default: 2
	db.SetConnMaxLifetime(0) //default: 0

	err = db.Ping()
	if err != nil {
		logrus.Fatalln("Could not connected to the database: ", err)
	}
	return db
}

package dals

import (
	"database/sql"
	"fmt"
	"log"
	"v1/constants"
)

func ConnectMysql() string {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		constants.DBUser,
		constants.DBPassword,
		constants.DBHost,
		constants.DBPort,
		constants.DBName,
	)

	return database
}

func OpenConnectionMysql() *sql.DB {
	db, err := sql.Open(constants.DBType, ConnectMysql())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect to database")

	return db
}

package repo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Password@123@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}


var dbConn = &DB{}

func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error){

	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)

	if err != nil {
		panic(err)
	}

	dbConn.SQL = d
	return dbConn, err
}
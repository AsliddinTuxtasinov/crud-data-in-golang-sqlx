package db_client

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	db, err := sqlx.Open("mysql", "root:@/testdb?parseTime=true") // dataSourceName => "root:@tcp(localhost:3306)/testdb"
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DBClient = db
}

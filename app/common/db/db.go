package db

import (
	config "app/common/config/db"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Setup() {
	var err error
	conf := config.DB

	DB, err = sql.Open("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Name)
	if err != nil {
		panic(err.Error())
	}
}

func Close() {
	DB.Close()
}

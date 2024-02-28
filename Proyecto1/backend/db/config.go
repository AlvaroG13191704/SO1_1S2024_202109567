package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

func GetDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/Proyecto1")
		if err != nil {
			return
		}
		err = db.Ping()
		if err == nil {
			fmt.Println("Successfully connected to the database")
		}
	})

	return db, err
}

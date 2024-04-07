package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	DB *sql.DB
}

var (
	once     sync.Once
	instance *DBClient
)

func ConnectToDB() *DBClient {
	once.Do(func() {
		db, err := sql.Open("mysql", "root:ga1318@tcp(35.184.115.84:3306)/Tarea4SO1")
		if err != nil {
			fmt.Println("Error connecting to the database: ", err)
			return
		}
		err = db.Ping()
		if err != nil {
			fmt.Println("Error pinging the database: ", err)
			return
		}
		fmt.Println("Successfully connected to the database")
		instance = &DBClient{DB: db}
	})
	return instance
}

// Insert data
func (client *DBClient) InsertData(name string, album string, year string, rank string) {
	_, err := client.DB.Exec("INSERT INTO `Data` (`Name`, `Album`, `Year`, `Rank`) VALUES (?, ?, ?, ?)", name, album, year, rank)
	if err != nil {
		fmt.Println("Error inserting data: ", err)
		return
	}
	fmt.Println("Data inserted successfully")
}

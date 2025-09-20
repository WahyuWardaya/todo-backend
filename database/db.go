package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/todolist_db" 
	
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi: %v", err)
	}

	// tes koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	fmt.Println("Berhasil koneksi ke database MySQL di localhost!")
}

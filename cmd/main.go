package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/lelandaure/ecom/cmd/api"
	"github.com/lelandaure/ecom/config"
	"github.com/lelandaure/ecom/db"
	"log"
)

func main() {
	storage, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(storage)

	server := api.NewAPIServer(":8080", storage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected.")
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	"ecole_drive/cmd/api"
	"ecole_drive/configs"
	"ecole_drive/db"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize MySQL storage: %v", err)
	}
	
	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start the API server: %v", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("DB: Successfully connected!")
}

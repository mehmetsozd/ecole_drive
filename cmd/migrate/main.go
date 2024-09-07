package main

import (
	"ecole_drive/db"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// .env dosyasından alınan bilgileri kullanarak MySQL konfigürasyonunu ayarla
	cfg := mysqlDriver.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_ADDRESS"),
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Veritabanı bağlantısı oluşturma
	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Migration işlemleri için driver oluşturma
	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migration işlemi için migrate nesnesi oluşturma
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Mevcut versiyonu kontrol etme
	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	// Komut satırı argümanlarına göre migration işlemi yapma
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

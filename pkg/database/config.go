package database

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_go-toolkit-secrets/pkg/secrets"
	"log"
	"os"
	"time"
)

type DB struct {
	Connection *sql.DB
}

func NewConnectionDB(db *mysql.Config) (*DB, error) {
	conn, err := sql.Open("mysql", db.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Connection: conn}, nil
}

func (Db *DB) Close() error {
	return Db.Connection.Close()
}

func GetDBConfig() (*mysql.Config, error) {
	client, err := secrets.NewClient()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_MYSQL_<CLUSTER>_<SCHEMA>_<SCHEMA>_ENDPOINT")
	dbUser, OK := client.GetSecret("DB_MYSQL_<CLUSTER>_<SCHEMA>_<SCHEMA>_<ROLE>_USER")

	if !OK {
		log.Println("error recovering Username")
	}

	dbPassword, OK := client.GetSecret("DB_MYSQL_<CLUSTER>_<SCHEMA>_<SCHEMA>_<ROLE>")

	if !OK {
		log.Println("error recovering password")
	}

	dbName := "dbname"

	if dbHost == "" || dbUser == "" || dbPassword == "" {
		return nil, errors.New("missing required environment configuration for the database")
	}

	return &mysql.Config{
		User:                 dbUser,
		Passwd:               dbPassword,
		Net:                  "tcp",
		Addr:                 dbHost,
		DBName:               dbName,
		Timeout:              100 * time.Millisecond,
		ReadTimeout:          100 * time.Millisecond,
		WriteTimeout:         100 * time.Millisecond,
		ParseTime:            true,
		AllowNativePasswords: true,
	}, nil
}

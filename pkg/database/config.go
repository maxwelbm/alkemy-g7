package database

import (
	"database/sql"
	"errors"
	"os"

	"github.com/go-sql-driver/mysql"
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
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbNet := os.Getenv("DB_NET")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbNet == "" {
		return nil, errors.New("missing required environment configuration for the database")
	}

	return &mysql.Config{
		User:      dbUser,
		Passwd:    dbPassword,
		Net:       dbNet,
		Addr:      dbHost,
		ParseTime: true,
		DBName:    dbName,
	}, nil
}

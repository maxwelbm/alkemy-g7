package database

import (
	"database/sql"
	"errors"
	"os"

	// "github.com/maxwelbm/alkemy-g7.git/internal/handler"

	"github.com/go-sql-driver/mysql"
)

type Db struct {
	Connection *sql.DB
}

func NewConnectionDb(db *mysql.Config) (*Db, error) {

	conn, err := sql.Open("mysql", db.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &Db{Connection: conn}, nil
}

func (Db *Db) Close() error {
	return Db.Connection.Close()
}

func GetDbConfig() (*mysql.Config, error) {

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

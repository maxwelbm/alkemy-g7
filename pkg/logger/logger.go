package logger

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"

	_ "github.com/go-sql-driver/mysql"
	logrus "github.com/sirupsen/logrus"
)

type Logger struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewLogger(db *sql.DB) *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{db: db, logger: logger}
}

func (l *Logger) Log(layer string, level, message string) {
	entry := model.LogEntry{
		Level:   level,
		Message: fmt.Sprintf("[%s] %s", layer, message),
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := l.db.Exec("INSERT INTO logs (level, message, time) VALUES (?, ?, ?)", entry.Level, entry.Message, entry.Time)
	if err != nil {
		log.Printf("failed to log to database: %v", err)
	}

	fmt.Printf("[%s] %s: %s\n", entry.Time, entry.Level, entry.Message)
}

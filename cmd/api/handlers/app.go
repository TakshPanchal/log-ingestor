package handlers

import (
	"database/sql"
	"log"
)

type Application struct {
	DB                 *sql.DB
	ErrLogger, ILogger *log.Logger
}

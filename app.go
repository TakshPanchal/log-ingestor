package main

import (
	"database/sql"
	"log"
)

type application struct {
	db     *sql.DB
	logger *log.Logger
}

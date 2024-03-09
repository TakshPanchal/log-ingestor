package models

import (
	"database/sql"
	"time"
)

type LogsModel struct {
	db *sql.DB
}

func (model LogsModel) Insert() {

}

type Log struct {
	id         int
	LogLevel   string    `json:"log_level"`
	ResourceId string    `json:"resource_id"`
	TraceId    string    `json:"trace_id"`
	SpanId     string    `json:"span_id"`
	Commit     string    `json:"commit"`
	CreatedAt  time.Time `json:"created_at"`
}

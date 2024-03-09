package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup
	port := ":3030"
	mux := http.NewServeMux()
	logger := log.New(os.Stdout, "", log.Ldate)
	db, err := setupDB()

	if err != nil {
		logger.Fatalln(err)
	}
	db.Ping()

	app := &application{
		db:     db,
		logger: logger,
	}

	mux.HandleFunc("/ingest", app.handleIngest)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	logger.Printf("Server is starting on %s port\n", port)
	err = srv.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}
}

func setupDB() (*sql.DB, error) {
	connStr := "postgres://sjubjzfl:1LoC9R0p0H1M3CWHEdP811E80DFk5c18@tiny.db.elephantsql.com/sjubjzfl?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

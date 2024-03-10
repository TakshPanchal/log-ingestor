package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"github.com/takshpanchal/log_ingestor/cmd/api/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup
	var port string
	flag.StringVar(&port, "port", ":8080", "port running server")
	flag.Parse()
	mux := http.NewServeMux()
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR: ", log.Ltime|log.Llongfile)
	db, err := setupDB()

	if err != nil {
		errLogger.Fatalln(err)
	}
	db.Ping()

	app := &handlers.Application{
		DB:        db,
		ILogger:   infoLogger,
		ErrLogger: errLogger,
	}

	mux.HandleFunc("/ingest", app.HandleIngest)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	infoLogger.Printf("Server is starting on %s port\n", port)
	err = srv.ListenAndServe()

	if err != nil {
		errLogger.Fatal(err)
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

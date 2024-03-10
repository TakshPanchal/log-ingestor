package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/takshpanchal/log_ingestor/cmd/api/handlers"
	"log"
	"net/http"
	"os"
)

var (
	// DB
	dbName, user, password, host, port string
)

func main() {
	// Setup
	var port string
	flag.StringVar(&port, "port", ":8080", "port running server")
	flag.Parse()
	loadEnv()
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

func loadEnv() {
	dbName = os.Getenv("DB_NAME")
	user = os.Getenv("USER_NAME")
	password = os.Getenv("DB_PASS")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
}

func setupDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s", dbName, user, password, host, port)
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

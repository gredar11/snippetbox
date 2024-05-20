package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"golangify.com/snippetbox/pkg/models/mysql"
)

func main() {

	// addr := flag.String("addr", ":4000", "IP address")
	// flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// connString := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Название MySQL источника данных")

	var connString string = "user=snippetbox password=snippetbox dbname=snippetbox sslmode=disable"

	db, err := openDB(connString)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	snippetModel := &mysql.SnippetModel{
		DB: db,
	}

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		snippetModel: snippetModel,
	}

	server := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запус веб-сервера на %s", connString)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	snippetModel *mysql.SnippetModel
}

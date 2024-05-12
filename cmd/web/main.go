package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golangify.com/snippetbox/pkg/models/mysql"
)

func main() {

	addr := flag.String("addr", ":4000", "IP address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connString := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Название MySQL источника данных")

	db, err := openDB(*connString)
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
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запус веб-сервера на %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
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

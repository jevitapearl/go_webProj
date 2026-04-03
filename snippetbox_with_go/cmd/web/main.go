package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // just runs init so it should be prefixed by _
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network address")
	flag.Parse() // After all flags to parse it

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	

	db, err := sql.Open("mysql", "web:Web@12345@/snippetbox?parseTime=true") // Returns pool of connections in mysql
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	// app := &application{
	// 	errorLog: errorLog
	// }

	

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), 	
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe() // TCP network address, servemux
	errorLog.Fatal(err)        // Calls os.exit(1) at the end
}

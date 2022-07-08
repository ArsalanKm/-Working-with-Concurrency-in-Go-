package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// connect to database
	db := initDB()
	db.Ping()
	// create sessions

	// create some channels

	// create waitgroup

	// set up the application config

	// sending email

	// listen for web connections
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}
func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgress not yet ready...")
		} else {
			log.Println("connected to database!")
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Print("Backing off for 1 seconds")
		time.Sleep(time.Second * 10)
		counts++
		continue
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

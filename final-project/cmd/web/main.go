package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

func main() {
	// connect to database
	db := initDB()
	// db.Ping()
	// create sessions
	session := initSession()
	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	// create some channels

	// create waitgroup
	wg := sync.WaitGroup{}

	// set up the application config

	app := Config{
		Session:  session,
		DB:       db,
		Wait:     &wg,
		InfoLog:  infoLog,
		ErrorLog: errLog,
	}
	// sending email

	// listen for signals

	go app.listenForShutDown()

	// listen for web connections
	app.serve()
}

func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	app.InfoLog.Println("Starting the Server..")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
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

func initSession() *scs.SessionManager {
	// set up sessions
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true
	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}

func (app *Config) listenForShutDown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutDown()
	os.Exit(0)
}

func (app *Config) shutDown() {
	app.InfoLog.Println("would run cleanup tasks")

	// block until waitgroup is empty
	app.Wait.Wait()
	app.InfoLog.Println("closing channels and shuting down application...")
}

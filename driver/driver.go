package driver

import (
	"database/sql"
	"log"

	"github.com/klahnen/spotifyService/config"
	"github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	conf := config.GetConfig()
	pgURL, err := pq.ParseURL(conf.PostgresURL)
	db, err := sql.Open("postgres", pgURL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

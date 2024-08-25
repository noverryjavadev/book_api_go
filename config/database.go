package config

import (
	"database/sql"
	"fmt"
	"fresh_api/helper"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
	schema   = "public"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbName, schema,
	)
	log.Printf("Terhubung ke database '%s' di %s:%s\n", dbName, host, port)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!!")
	return db
}

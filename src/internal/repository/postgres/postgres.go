package postgres

import (
	"fmt"
	"persserv/internal/config"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfgDB config.DB) (*sqlx.DB, error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfgDB.Host, cfgDB.Port, cfgDB.Username, cfgDB.DBName, cfgDB.Password, cfgDB.SSLMode)

	db, err := sqlx.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

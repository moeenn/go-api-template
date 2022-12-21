package database

import (
	"app/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sqlx.DB
}

/**
 *  establish a connection with the database
 *
 */
func (db *Database) Connect(config config.DatabaseConfig) error {
	connString := config.ConnectionString()
	conn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	db.Conn = conn
	return nil
}

/**
 *  close an open connection to the database
 *
 */
func (db *Database) Disconnect() {
	db.Conn.Close()
}

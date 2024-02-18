package database

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Username string
	Password string
	Db       string
	Endpoint string
}

func NewPostgresUnsecureConn(cfg Config) (*sql.DB, error) {
	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%v:%v@%v/%v?sslmode=disable",
			cfg.Username,
			cfg.Password,
			cfg.Endpoint,
			cfg.Db,
		),
	)
	if err != nil {
		return nil, err
	}
	// db conn is lazy. ping in order to verify it.
	err = verifyPostgresConn(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewPostgresSecureConn(cfg Config) (*sql.DB, error) {
	panic("NewPostgresSecureConn not implemented")
}

func verifyPostgresConn(conn *sql.DB) error {
	// db conn is lazy. ping in order to verify it.
	return conn.Ping()
}

package config

import (
	"os"
)

type Config struct {
	SQLiteConnString   string `json:"sqlite_conn_string"`
	PostgresConnString string `json:"postgres_conn_string"`
}

func LoadConfig() (*Config, error) {
	sqliteConn := os.Getenv("SQLITE_CONN_STRING")
	if sqliteConn == "" {
		sqliteConn = "sqlite.db" // Default value if the environment variable is not set
	}

	postgresConn := os.Getenv("POSTGRES_CONN_STRING")
	if postgresConn == "" {
		postgresConn = "postgres://user:password@localhost:5432/dbname?sslmode=disable" // Default value
	}

	return &Config{
		SQLiteConnString:   sqliteConn,
		PostgresConnString: postgresConn,
	}, nil
}
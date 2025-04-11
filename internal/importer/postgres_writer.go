package importer

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresWriter struct {
	db *sql.DB
}

func (pw *PostgresWriter) OpenConnection(connStr string) error {
	var err error
	pw.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open connection to PostgreSQL: %w", err)
	}
	return pw.db.Ping()
}

func (pw *PostgresWriter) WriteData(tableName string, data map[string]interface{}) error {
	columns := ""
	values := ""
	args := []interface{}{}

	i := 1
	for col, val := range data {
		if columns != "" {
			columns += ", "
			values += ", "
		}
		columns += col
		values += fmt.Sprintf("$%d", i)
		args = append(args, val)
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, values)
	_, err := pw.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to write data to PostgreSQL: %w", err)
	}
	return nil
}

func (pw *PostgresWriter) Close() error {
	if pw.db != nil {
		return pw.db.Close()
	}
	return nil
}
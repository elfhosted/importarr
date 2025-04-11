package importer

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteReader struct {
	db *sql.DB
}

func (sr *SQLiteReader) OpenConnection(connStr string) error {
	var err error
	sr.db, err = sql.Open("sqlite3", connStr)
	if err != nil {
		return fmt.Errorf("failed to open connection to SQLite: %w", err)
	}
	return nil
}

func (sr *SQLiteReader) Close() error {
	if sr.db != nil {
		return sr.db.Close()
	}
	return nil
}

func (sr *SQLiteReader) ReadData(tableName string) ([]map[string]interface{}, error) {
	rows, err := sr.db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}
		results = append(results, rowMap)
	}

	return results, nil
}

type PostgreSQLWriter struct {
	db *sql.DB
}

func (pw *PostgreSQLWriter) OpenConnection(connStr string) error {
	var err error
	pw.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open connection to PostgreSQL: %w", err)
	}
	return nil
}

func (pw *PostgreSQLWriter) Close() error {
	if pw.db != nil {
		return pw.db.Close()
	}
	return nil
}

func (pw *PostgreSQLWriter) WriteData(tableName string, data map[string]interface{}) error {
	// Implementation for writing data to PostgreSQL
	return nil
}

func main() {
	sqliteReader := &SQLiteReader{}
	err := sqliteReader.OpenConnection("your_connection_string") // Replace "your_connection_string" with the actual connection string
	if err != nil {
		log.Fatalf("Failed to open SQLite connection: %v", err)
	}
	defer sqliteReader.Close()

	data, err := sqliteReader.ReadData("table_name") // Replace "table_name" with the actual table name
	if err != nil {
		log.Fatalf("Failed to read data from SQLite: %v", err)
	}

	postgresWriter := &PostgreSQLWriter{}
	err = postgresWriter.OpenConnection("your_postgres_connection_string") // Replace "your_postgres_connection_string" with the actual connection string
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL connection: %v", err)
	}
	defer postgresWriter.Close()

	for _, row := range data {
		if err := postgresWriter.WriteData("table_name", row); err != nil { // Replace "table_name" with the actual table name
			log.Fatalf("Failed to write data to PostgreSQL: %v", err)
		}
	}
}
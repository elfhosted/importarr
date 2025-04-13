package importer

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresWriter struct {
	db *sql.DB
}

func (pw *PostgresWriter) OpenConnection(connStr string) error {
	var err error
	pw.db, err = sql.Open("postgres", connStr)
	if (err != nil) {
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
		// Wrap column names in double quotes to preserve case
		columns += fmt.Sprintf(`"%s"`, col)
		values += fmt.Sprintf("$%d", i)
		args = append(args, val)
		i++
	}

	// Wrap the table name in double quotes to preserve case
	query := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`, tableName, columns, values)
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

// GetTables retrieves the list of all tables in the PostgreSQL database
func (pw *PostgresWriter) GetTables() ([]string, error) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE';
	`

	rows, err := pw.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tables: %w", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %w", err)
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func (pw *PostgresWriter) UpdateSequence(tableName, columnName string) error {
	query := fmt.Sprintf(`
		SELECT setval(pg_get_serial_sequence('"%%s"', '%%s'), MAX("%%s") + 1)
		FROM "%%s"
	`, tableName, columnName, columnName, tableName)

	_, err := pw.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to update sequence for table %s: %w", tableName, err)
	}
	return nil
}
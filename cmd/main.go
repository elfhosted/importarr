package main

import (
    "log"

    "github.com/elfhosted/importarr/internal/config"
    "github.com/elfhosted/importarr/internal/importer"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize PostgreSQL writer
    postgresWriter := &importer.PostgresWriter{}
    if err := postgresWriter.OpenConnection(cfg.PostgresConnString); err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }
    defer postgresWriter.Close()

    // Get the list of tables from PostgreSQL
    tables, err := postgresWriter.GetTables()
    if err != nil {
        log.Fatalf("Failed to retrieve tables from PostgreSQL: %v", err)
    }

    // Initialize SQLite reader
    sqliteReader := &importer.SQLiteReader{}
    if err := sqliteReader.OpenConnection(cfg.SQLiteConnString); err != nil {
        log.Fatalf("Failed to connect to SQLite: %v", err)
    }
    defer sqliteReader.Close()

    // Process each table
    for _, table := range tables {
        log.Printf("Processing table: %s", table)

        // Read data from SQLite
        data, err := sqliteReader.ReadData(table)
        if err != nil {
            log.Printf("Failed to read data from table %s: %v", table, err)
            continue
        }

        // Write data to PostgreSQL
        for _, row := range data {
            if err := postgresWriter.WriteData(table, row); err != nil {
                log.Printf("Failed to write data to table %s in PostgreSQL: %v", table, err)
            }
        }

        // Update the sequence for the table
        if err := postgresWriter.UpdateSequence(table, "Id"); err != nil {
            log.Printf("Failed to update sequence for table %s: %v", table, err)
        }
    }

    log.Println("Data import completed successfully!")
}
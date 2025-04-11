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
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize SQLite reader
    sqliteReader := importer.SQLiteReader{}
    err = sqliteReader.OpenConnection(cfg.SQLiteConnString)
    if err != nil {
        log.Fatalf("Error connecting to SQLite: %v", err)
    }
    defer sqliteReader.CloseConnection()

    // Initialize PostgreSQL writer
    postgresWriter := importer.PostgresWriter{}
    err = postgresWriter.OpenConnection(cfg.PostgresConnString)
    if err != nil {
        log.Fatalf("Error connecting to PostgreSQL: %v", err)
    }
    defer postgresWriter.CloseConnection()

    // Read data from SQLite
    data, err := sqliteReader.ReadData()
    if err != nil {
        log.Fatalf("Error reading data from SQLite: %v", err)
    }

    // Write data to PostgreSQL
    err = postgresWriter.WriteData(data)
    if err != nil {
        log.Fatalf("Error writing data to PostgreSQL: %v", err)
    }

    log.Println("Data import completed successfully.")
}
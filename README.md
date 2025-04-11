# SQLite to PostgreSQL Importer

This utility is designed to import data from an SQLite database into a PostgreSQL database. It handles differences in schema, including case sensitivity in table names and variations in data types.

## Project Structure

```
importarr
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── importer
│   │   ├── sqlite_reader.go    # Handles reading data from SQLite
│   │   ├── postgres_writer.go   # Manages writing data to PostgreSQL
│   │   └── mapper.go           # Maps SQLite types to PostgreSQL types
│   └── config
│       └── config.go          # Configuration settings for the application
├── go.mod                      # Module definition
├── go.sum                      # Dependency checksums
└── README.md                   # Project documentation
```

## Installation

To install the necessary dependencies, run:

```
go mod tidy
```

## Configuration

Before running the utility, ensure that you have a configuration file set up with the necessary database connection strings for both SQLite and PostgreSQL. The configuration can be loaded from a file or environment variables.

## Usage

To run the application, execute the following command:

```
go run cmd/main.go
```

This will initiate the import process, reading data from the specified SQLite database and writing it to the PostgreSQL database.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
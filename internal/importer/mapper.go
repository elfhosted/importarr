package importer

type Mapper struct{}

func (m *Mapper) MapType(sqliteType string) string {
    switch sqliteType {
    case "INTEGER":
        return "INTEGER"
    case "TEXT":
        return "VARCHAR"
    case "REAL":
        return "DOUBLE PRECISION"
    case "BLOB":
        return "BYTEA"
    case "NULL":
        return "NULL"
    default:
        return "TEXT" // Default mapping for unknown types
    }
}

func (m *Mapper) MapTableName(sqliteTableName string) string {
    // Convert SQLite table name to PostgreSQL format (e.g., lower case)
    return lowerCase(sqliteTableName)
}

func lowerCase(s string) string {
    // Convert string to lower case
    return strings.ToLower(s)
}
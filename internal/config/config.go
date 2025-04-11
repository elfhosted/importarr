package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SQLiteConnectionString string `json:"sqlite_connection_string"`
	PostgresConnectionString string `json:"postgres_connection_string"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
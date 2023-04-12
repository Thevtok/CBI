package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Host            string `json:"host"`
		Port            int    `json:"port"`
		Name            string `json:"name"`
		User            string `json:"user"`
		Password        string `json:"password"`
		MaxOpenConns    int    `json:"max_open_conns"`
		MaxIdleConns    int    `json:"max_idle_conns"`
		MaxConnLifetime int    `json:"max_conn_lifetime"`
		URL             string `json:"-"`
	} `json:"db"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

func LoadConfig(configFile string) (Config, error) {
	var config Config

	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config.Database)
	if err != nil {
		return config, err
	}

	return config, nil
}

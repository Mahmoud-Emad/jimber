package internal

import (
	"io"
	"strings"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Database DatabaseConfiguration `toml:"database"`
	Server   ServerConfiguration   `toml:"server"`
}

type ServerConfiguration struct {
	Host         string `toml:"host"`
	Port         int64  `toml:"port"`
	JWTSecretKey string `toml:"jwtSecretKey"`
}

type DatabaseConfiguration struct {
	Host     string `toml:"host"`
	Port     int64  `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
}

// Read the config file.
func ReadConfigFromFile(path string) (Configuration, error) {
	config := Configuration{}
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return Configuration{}, cantLoadConfigFileError
	}
	err = config.validateConfiguration()
	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}

// Read the config from a string.
func ReadConfigFromString(content string) (Configuration, error) {
	return ReadConfigFromReader(strings.NewReader(content))
}

// Read the config from a reader.
func ReadConfigFromReader(r io.Reader) (Configuration, error) {
	config := Configuration{}
	_, err := toml.DecodeReader(r, &config)
	if err != nil {
		return Configuration{}, cantDecodeConfigError
	}
	err = config.validateConfiguration()
	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}

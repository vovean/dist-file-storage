package config

import "pkg/config"

type Config struct {
	Server struct {
		Port uint `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	Storage struct {
		Root string          `yaml:"root" env:"STORAGE_ROOT"`
		Size config.DataSize `yaml:"size" env:"STORAGE_SIZE"`
	} `yaml:"storage"`

	API struct {
		ServeChunkSize config.DataSize `yaml:"serve_chunk_size" env:"API_SERVECHUNKSIZE"`
	} `yaml:"api"`
}

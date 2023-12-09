package config

import "pkg/config"

type Config struct {
	Server struct {
		Port uint `yaml:"port"`
	} `yaml:"server"`

	Storage struct {
		Root string          `yaml:"root"`
		Size config.DataSize `yaml:"size"`
	} `yaml:"storage"`

	API struct {
		ServeChunkSize config.DataSize `yaml:"serve_chunk_size"`
	} `yaml:"api"`
}

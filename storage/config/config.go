package config

import "pkg/config"

type Config struct {
	Server struct {
		// На каком порту запустить сервер
		Port uint `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	Storage struct {
		// Корневая папка, в которой будут создаваться файлы и папки для хранения данных
		Root string `yaml:"root" env:"STORAGE_ROOT"`
		// Размер места, которое должно использоваться хранилищем
		Size config.DataSize `yaml:"size" env:"STORAGE_SIZE"`
	} `yaml:"storage"`

	API struct {
		// Размер батча, которым данные будут отдаваться при скачивании
		ServeChunkSize config.DataSize `yaml:"serve_chunk_size" env:"API_SERVECHUNKSIZE"`
	} `yaml:"api"`
}

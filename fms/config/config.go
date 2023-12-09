package config

type Config struct {
	Server struct {
		Port uint `yaml:"port"`
	} `yaml:"server"`

	FileParts uint `yaml:"file_parts"`

	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"db"`
}

package config

type Config struct {
	Server struct {
		// На каком порту запустить сервер
		Port uint `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	// На сколько частей делить файл
	FileParts uint `yaml:"file_parts" env:"FILEPARTS"`

	DB struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		User     string `yaml:"user" env:"DB_USER"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Dbname   string `yaml:"dbname" env:"DB_DBNAME"`
		SSLMode  string `yaml:"ssl_mode" env:"DB_SSLMODE"` // постгряшный sslmode, ставьте disable, если не уверены
	} `yaml:"db"`
}

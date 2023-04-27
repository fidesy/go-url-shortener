package config

type Config struct {
	Host     string         `yaml:"host"`
	Port     int64          `yaml:"port"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

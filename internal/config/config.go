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
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

var DefaultConfig = &Config{
	Host: "http://localhost",
	Port: 8000,
	Postgres: PostgresConfig{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		SSLMode:  "disable",
	},
}

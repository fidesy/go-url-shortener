package config

type Config struct {
	Host     string   `yaml:"host"`
	Port     string   `yaml:"port"`
	Database string   `yaml:"database"`
	Postgres Postgres `yaml:"postgres"`
	Mongo    Mongo    `yaml:"mongo"`
}

type Postgres struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Mongo struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

var Default = &Config{
	Host: "http://localhost",
	Port: "8000",
	Database: "postgres",
	Postgres: Postgres{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     "5432",
		DBName:   "postgres",
		SSLMode:  "disable",
	},
	Mongo: Mongo{
		Username: "mongo",
		Password: "mongo",
		Host:     "localhost",
		Port:     "27017",
	},
}

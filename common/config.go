package common

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	Server *serverModel `yaml:"server"`
	DB     *DbConfig    `yaml:"db"`
	Redis  *RedisConfig `yaml:"redis"`
}

type serverModel struct {
	Model                string        `yaml:"model"`
	Host                 string        `yaml:"host"`
	Port                 string        `yaml:"port"`
	EnableHttps          bool          `yaml:"enable_https"`
	CertFile             string        `yaml:"cert_file"`
	KeyFile              string        `yaml:"key_file"`
	JwtPublicKey         string        `yaml:"jwt_public_key"`
	JwtPrivateKey        string        `yaml:"jwt_private_key"`
	TokenExpireSecond    time.Duration `yaml:"token_expire_second"`
	SystemStaticFilePath string        `yaml:"system_static_file_path"`
	Banner               *Banner       `yaml:"banner"`
}

type Banner struct {
	Name    string `yaml:"name"`
	Loading bool   `yaml:"loading"`
}

type DbConfig struct {
	Host        string  `yaml:"host"`
	Port        string  `yaml:"port"`
	Username    string  `yaml:"username"`
	Password    string  `yaml:"password"`
	Database    string  `yaml:"database"`
	Charset     string  `yaml:"charset"`
	Category    string  `yaml:"category"`
	PrintSql    bool    `yaml:"sql"`
	TablePrefix string  `yaml:"table_prefix"`
	DbPool      *DbPool `yaml:"pool"`
	DbInit      *DbInit `yaml:"init"`
}

type DbPool struct {
	MaxIdleConns    int           `yaml:"maxIdle"`
	MaxOpenConns    int           `yaml:"maxOpen"`
	ConnMaxLifetime time.Duration `yaml:"maxLifetime"`
}

type DbInit struct {
	Name   string `yaml:"name"`
	Status bool   `yaml:"status"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) ReadConfig() *Config {
	file, err := os.ReadFile("config/config.local.yaml")
	if err != nil {
		log.Fatalln("Read config file error: ", err)
		return nil
	}
	if yaml.Unmarshal(file, config) != nil {
		log.Fatalln("Unmarshal config file error: ", err)
		return nil
	}
	return config
}

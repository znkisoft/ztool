package config

import (
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"strconv"
)

type Configurations struct {
	Redis  RedisConfig
	Server ServerConfig
}

type RedisConfig struct {
	Address  string `validate:"required"`
	Password string `validate:"required"`
	DB       int    `validate:"required"`
}

type DbConfig struct {
	Name     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required"`
}

type ServerConfig struct {
	Host string `validate:"required,ip"`
}

type OSSConfig struct {
	AccessKey string `validate:"required"`
	SecretKey string `validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func LoadRedisConfig() *RedisConfig {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("redis env variable is missing, %s", err)
	}
	var redisConfig = &RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
	validateConfig(redisConfig)
	return redisConfig
}

func LoadServerConfig() *ServerConfig {
	host := os.Getenv("ECS_SERVER_HOST")
	if host == "" {
		log.Fatalf("host info is missing.")
	}
	var serverConfig = &ServerConfig{
		Host: host,
	}
	validateConfig(serverConfig)
	return serverConfig
}

func LoadDbConfig() *DbConfig {
	var dbConfig = &DbConfig{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
	}
	validateConfig(dbConfig)
	return dbConfig
}

func LoadOSSConfig() *OSSConfig {
	var ossConfig = &OSSConfig{
		AccessKey: os.Getenv("OSS_ACCESS_KEY"),
		SecretKey: os.Getenv("OSS_SECRET_KEY"),
	}
	validateConfig(ossConfig)
	return ossConfig
}

func validateConfig(config interface{}) {
	err := validate.Struct(config)
	if err != nil {
		log.Fatalf(">>> Error: %s:\n", err.Error())
	}
}

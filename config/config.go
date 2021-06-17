package config

import (
	"log"
	"os"
	"strconv"
)

type Configurations struct {
	Redis  RedisConfig
	Server ServerConfig
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type DbConfig struct {
	Name     string
	User     string
	Password string
	Host     string
}

type ServerConfig struct {
	Host string
}

func LoadRedisConfig() *RedisConfig {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("redis env variable is missing, %s", err)
	}
	return &RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}

func LoadServerConfig() *ServerConfig {
	host := os.Getenv("ECS_SERVER_HOST")
	if host == "" {
		host = ""
	}
	return &ServerConfig{
		Host: host,
	}
}

func LoadDbConfig() *DbConfig {
	return &DbConfig{
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
	}
}

package config

import (
	"log"
	"os"
	"strconv"
)

// Configurations exported
type Configurations struct {
	Redis RedisConfig
}

// RedisConfig  exported
type RedisConfig struct {
	Address  string
	Password string
	DB       int
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

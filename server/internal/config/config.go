// Package config 负责从环境变量加载并校验服务配置。
package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// Config 保存服务端运行所需的配置项。
type Config struct {
	Port          string
	DatabaseURL   string
	JWTSecret     string
	AdminUsername string
	AdminPassword string
	JWTTTL        time.Duration
}

// Load 从环境变量读取配置并校验必填项与合法性。
func Load() (Config, error) {
	cfg := Config{
		Port:          os.Getenv("PORT"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		AdminUsername: os.Getenv("ADMIN_USERNAME"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		JWTTTL:        24 * time.Hour,
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	} else {
		p, err := strconv.ParseUint(cfg.Port, 10, 16)
		if err != nil || p == 0 {
			return Config{}, errors.New("invalid PORT: must be an integer between 1 and 65535")
		}
		cfg.Port = strconv.FormatUint(p, 10)
	}

	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("missing DATABASE_URL")
	}

	if cfg.JWTSecret == "" {
		return Config{}, errors.New("missing JWT_SECRET")
	}

	if len(cfg.JWTSecret) < 16 {
		return Config{}, errors.New("JWT_SECRET must be at least 16 bytes")
	}

	return cfg, nil
}

package config

import (
	"os"
)

type EnvConfig struct {
	JwtSecret []byte
}

func SetupConfig() *EnvConfig {
	env := EnvConfig{
		JwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
	return &env
}

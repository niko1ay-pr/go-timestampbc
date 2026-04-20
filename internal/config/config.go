package config

import "github.com/caarlos0/env"

type Config struct {
	HTTPPort string `env:"HTTP_PORT" envDefault:"8008"`
	HTTPHost string `env:"HTTP_HOST" envDefault:"localhost"`

	SQLitePath string `env:"SQLITE_PATH" envDefault:"./data/sqlitedb"`
	BoltPath   string `env:"BOLT_PATH" envDefault:"./data/bolt.db"`

	BlockchainDifficulty int    `env:"BLOCKCHAIN_DIFFICULTY" envDefault:"4"`
	MiningInterval       string `env:"MINING_INTERVAL" envDefault:"10s"`

	AppEnv   string `env:"APP_ENV" envDefault:"develop"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

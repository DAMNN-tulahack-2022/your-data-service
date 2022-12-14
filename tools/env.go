package tools

import (
	"github.com/caarlos0/env"
)

type GenericEnvAppConfig struct {
	DataProxyPort            string `env:"DATA_PROXY_PORT"`
	AuthFlowSault        string `env:"FLOW_SAULT"`
	AdminPermissionSault string `env:"PERMISSION_SAULT"`

	DbSoruceDSN     string `env:"DB_DSN"`
	MaxIdleConns    int    `env:"MAX_IDLE_CONNS"`
	MaxOpenConns    int    `env:"MAX_OPEN_CONNS"`
	MaxIdleTime     int    `env:"MAX_IDLE_TIME"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME"`

	CacheAddr                 string `env:"CACHE_ADDR"`
	CachePassword             string `env:"CACHE_PASSWORD"`
	CacheDbCursor             int    `env:"CACHE_DB_CURSOR"`
	CacheConnectionMaxRetries int    `env:"CACHE_CONN_MAX_RETRIES"`
}

func LoadDammnAppConfig(log *YourLogger) (*GenericEnvAppConfig, error) {
	cfg := new(GenericEnvAppConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

package config

import (
	"fmt"

	"github.com/ardanlabs/conf/v3"
)

type Config struct {
	Port             string `conf:"default:8081"`
	DatabaseUrl      string `conf:"default:localhost"`
	DatabasePort     int    `conf:"default:5432"`
	DatabaseUser     string `conf:"default:postgres"`
	DatabasePassword string `conf:"default:postgres"`
	DatabaseDb       string `conf:"default:subscriptions"`
}

func LoadConfig() (c Config, err error) {
	if _, err = conf.Parse("", &c); err != nil {
		return c, fmt.Errorf("failed to parse config: %w", err)
	}
	return
}

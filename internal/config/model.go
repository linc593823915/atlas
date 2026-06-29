package config

import (
	"github.com/linc593823915/atlas/internal/logger"
	"github.com/linc593823915/atlas/internal/version"
)

const (
	DefaultConfigPath = "configs/config.yaml"

	envAppName      = "ATLAS_APP_NAME"
	envAppVersion   = "ATLAS_APP_VERSION"
	envLoggerLevel  = "ATLAS_LOGGER_LEVEL"
	envLoggerFormat = "ATLAS_LOGGER_FORMAT"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	Logger LoggerConfig `yaml:"logger"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func defaultConfig() Config {
	return Config{
		App: AppConfig{
			Name:    "Atlas",
			Version: version.Version,
		},
		Logger: LoggerConfig{
			Level:  logger.LevelInfo,
			Format: logger.FormatText,
		},
	}
}

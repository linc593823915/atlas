package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/linc593823915/atlas/internal/logger"
	"gopkg.in/yaml.v3"
)

func loadConfig(path string, lookupEnv lookupEnvFunc) (Config, error) {
	config := defaultConfig()
	if err := loadConfigFile(path, &config); err != nil {
		return Config{}, err
	}
	applyEnvOverrides(&config, lookupEnv)
	if err := validateConfig(config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func loadConfigFile(path string, config *Config) error {
	resolvedPath, err := resolveConfigPath(path)
	if err != nil {
		return err
	}
	content, err := os.ReadFile(resolvedPath)
	if err != nil {
		return fmt.Errorf("read config file %s: %w", resolvedPath, err)
	}
	return unmarshalConfig(resolvedPath, content, config)
}

func unmarshalConfig(path string, content []byte, config *Config) error {
	if err := yaml.Unmarshal(content, config); err != nil {
		return fmt.Errorf("parse config file %s: %w", path, err)
	}
	return nil
}

func applyEnvOverrides(config *Config, lookupEnv lookupEnvFunc) {
	overrideString(&config.App.Name, lookupEnv, envAppName)
	overrideString(&config.App.Version, lookupEnv, envAppVersion)
	overrideString(&config.Logger.Level, lookupEnv, envLoggerLevel)
	overrideString(&config.Logger.Format, lookupEnv, envLoggerFormat)
}

func overrideString(target *string, lookupEnv lookupEnvFunc, key string) {
	if lookupEnv == nil {
		return
	}
	value, ok := lookupEnv(key)
	if ok && value != "" {
		*target = value
	}
}

func validateConfig(config Config) error {
	if err := validateAppConfig(config.App); err != nil {
		return err
	}
	if err := validateLoggerLevel(config.Logger.Level); err != nil {
		return err
	}
	return validateLoggerFormat(config.Logger.Format)
}

func validateAppConfig(config AppConfig) error {
	if config.Name == "" {
		return fmt.Errorf("invalid app name: empty")
	}
	if config.Version == "" {
		return fmt.Errorf("invalid app version: empty")
	}
	return nil
}

func validateLoggerLevel(level string) error {
	switch level {
	case logger.LevelDebug, logger.LevelInfo, logger.LevelWarn, logger.LevelError:
		return nil
	default:
		return fmt.Errorf("invalid logger level: %s", level)
	}
}

func validateLoggerFormat(format string) error {
	switch format {
	case logger.FormatText, logger.FormatJson:
		return nil
	default:
		return fmt.Errorf("invalid logger format: %s", format)
	}
}

func resolveConfigPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("resolve config path %s: %w", path, err)
	}
	return searchConfigPath(currentDir, path), nil
}

func searchConfigPath(dir, path string) string {
	for {
		candidate := filepath.Join(dir, path)
		if fileExists(candidate) {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return path
		}
		dir = parent
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

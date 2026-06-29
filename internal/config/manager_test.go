package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestManagerLoadUsesDefaultFileAndEnvPrecedence(t *testing.T) {
	path := writeConfigFile(t, strings.Join([]string{
		"app:",
		"  name: Atlas Local",
		"  version: 1.2.3",
		"logger:",
		"  level: warn",
		"  format: json",
		"",
	}, "\n"))
	manager := NewManager(
		WithConfigPath(path),
		WithLookupEnv(envMap(map[string]string{
			envAppVersion:  "9.9.9",
			envLoggerLevel: "error",
		})),
	)

	if err := manager.Load(); err != nil {
		t.Fatal(err)
	}

	assertAppConfig(t, manager.App(), "Atlas Local", "9.9.9")
	assertLoggerConfig(t, manager.Logger(), "error", "json")
}

func TestManagerLoadKeepsDefaultsForMissingFileValues(t *testing.T) {
	path := writeConfigFile(t, "logger:\n  level: debug\n")
	manager := NewManager(WithConfigPath(path), WithLookupEnv(envMap(nil)))

	if err := manager.Load(); err != nil {
		t.Fatal(err)
	}

	assertAppConfig(t, manager.App(), "Atlas", "0.0.1")
	assertLoggerConfig(t, manager.Logger(), "debug", "text")
}

func TestManagerLoadReturnsErrorWhenConfigFileMissing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.yaml")
	manager := NewManager(WithConfigPath(path), WithLookupEnv(envMap(nil)))

	if err := manager.Load(); err == nil {
		t.Fatal("expected config load error")
	}
}

func TestManagerLoadReturnsErrorForInvalidEnvOverride(t *testing.T) {
	path := writeConfigFile(t, "logger:\n  level: info\n  format: text\n")
	manager := NewManager(
		WithConfigPath(path),
		WithLookupEnv(envMap(map[string]string{envLoggerFormat: "xml"})),
	)

	if err := manager.Load(); err == nil {
		t.Fatal("expected config validation error")
	}
}

func TestManagerLoadReturnsErrorForEmptyAppName(t *testing.T) {
	path := writeConfigFile(t, "app:\n  name: \"\"\n")
	manager := NewManager(WithConfigPath(path), WithLookupEnv(envMap(nil)))

	if err := manager.Load(); err == nil {
		t.Fatal("expected app validation error")
	}
}

func TestDefaultConfigPathRemainsStable(t *testing.T) {
	if !strings.HasSuffix(DefaultConfigPath, "configs/config.yaml") {
		t.Fatalf("unexpected default config path: %s", DefaultConfigPath)
	}
}

func writeConfigFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func envMap(values map[string]string) lookupEnvFunc {
	return func(key string) (string, bool) {
		value, ok := values[key]
		return value, ok
	}
}

func assertLoggerConfig(t *testing.T, config LoggerConfig, level, format string) {
	t.Helper()
	if config.Level != level {
		t.Fatalf("level = %s, want %s", config.Level, level)
	}
	if config.Format != format {
		t.Fatalf("format = %s, want %s", config.Format, format)
	}
}

func assertAppConfig(t *testing.T, config AppConfig, name, version string) {
	t.Helper()
	if config.Name != name {
		t.Fatalf("name = %s, want %s", config.Name, name)
	}
	if config.Version != version {
		t.Fatalf("version = %s, want %s", config.Version, version)
	}
}

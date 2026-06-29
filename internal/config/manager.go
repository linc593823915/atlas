package config

import "os"

type Manager struct {
	path      string
	lookupEnv lookupEnvFunc
	config    Config
}

func NewManager(options ...Option) *Manager {
	manager := &Manager{
		path:      DefaultConfigPath,
		lookupEnv: os.LookupEnv,
		config:    defaultConfig(),
	}
	applyOptions(manager, options)
	return manager
}

func (m *Manager) Load() error {
	config, err := loadConfig(m.path, m.lookupEnv)
	if err != nil {
		return err
	}
	m.config = config
	return nil
}

func (m *Manager) Config() Config {
	return m.config
}

func (m *Manager) App() AppConfig {
	return m.config.App
}

func (m *Manager) Logger() LoggerConfig {
	return m.config.Logger
}

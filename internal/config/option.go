package config

type Option func(*Manager)

type lookupEnvFunc func(string) (string, bool)

func WithConfigPath(path string) Option {
	return func(manager *Manager) {
		if path != "" {
			manager.path = path
		}
	}
}

func WithLookupEnv(lookupEnv lookupEnvFunc) Option {
	return func(manager *Manager) {
		if lookupEnv != nil {
			manager.lookupEnv = lookupEnv
		}
	}
}

func applyOptions(manager *Manager, options []Option) {
	for _, option := range options {
		if option != nil {
			option(manager)
		}
	}
}

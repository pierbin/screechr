package config

type Config struct {
	ConnString string
	Driver     string
}

func NewConfig() *Config {
	cfg := &Config{
		ConnString: ":memory:",
		Driver:     "sqlite3",
	}

	return cfg
}

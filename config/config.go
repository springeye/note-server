package config

type AppConfig struct {
	Debug bool
}

var DefaultConfig = AppConfig{Debug: true}

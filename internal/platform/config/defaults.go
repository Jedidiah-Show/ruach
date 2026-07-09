package config

func Default() *Config {
	return &Config{
		App: AppConfig{
			Name:    "Ruach",
			Version: "dev",
			Env:     "development",
		},
		Logging: LoggingConfig{
			Level: "info",
			Format: "text",
		},
		AI: AIConfig{
			Provider: "openai",
			Model:    "gpt-5.5",
		},
	}
}

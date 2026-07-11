package config

type Config struct {
	App      AppConfig
	Logging  LoggingConfig
	AI       AIConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

type LoggingConfig struct {
	Level string
	Format string
}

type AIConfig struct {
	Provider string
	Model    string
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Name string
}

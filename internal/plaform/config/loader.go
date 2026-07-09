package config

func Load() (*Config, error) {
	cfg := Default()
	return cfg, nil
}

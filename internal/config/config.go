package config

type HTTPConfig struct {
	Port int
}

type PostgresConfig struct {
	DSN string
}

type Config struct {
	HTTPConfig     HTTPConfig
	PostgresConfig PostgresConfig
}

type Loader interface {
	Load() Config
}

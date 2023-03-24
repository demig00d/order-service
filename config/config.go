package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Log       `yaml:"logger"`
		PG        `yaml:"postgres"`
		JetStream `yaml:"jetstream"`
		CacheSize int `env-required:"true" yaml:"cache_size" env:"CACHE_SIZE"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	JetStream struct {
		StreamName              string `env-required:"true" yaml:"stream_name" env:"STREAM_NAME"`
		StreamSubjects          string `env-required:"true" yaml:"stream_subjects" env:"STREAM_SUBJECTS"`
		SubjectNameOrderCreated string `env-required:"true" yaml:"subject_name_order_created" env:"SUBJECT_NAME_ORDER_CREATED"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	return cfg, err
}

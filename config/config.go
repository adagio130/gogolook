package config

type Config struct {
	Server Server `mapstructure:"server" yaml:"server"`
}

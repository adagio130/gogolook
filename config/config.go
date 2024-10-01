package config

type Config struct {
	Server Server `mapstructure:"server" yaml:"server"`
	DB     DB     `mapstructure:"db" yaml:"db"`
}

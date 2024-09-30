package config

type Server struct {
	Port string `mapstructure:"port" yaml:"port" default:"8080"`
	Mode string `mapstructure:"mode" yaml:"mode" default:"debug"`
}

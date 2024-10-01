package config

type DB struct {
	Driver  string `mapstructure:"driver" yaml:"driver" default:"mysql"`
	Dsn     string `mapstructure:"dsn" yaml:"dsn" default:"file:test.db?cache=shared&mode=memory"`
	MaxOpen int    `mapstructure:"max_open" yaml:"max_open" default:"10"`
}

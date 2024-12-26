package config

type WalletConf struct {
	Port string `env:"PORT" yaml:"port"`
	DSN  string `env:"DSN" yaml:"dsn"`
}

package config

import (
	"context"
	_ "embed"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
	"log/slog"
	"sync"
)

var Object WalletConf
var once sync.Once

//go:embed default.yml
var defaultConf []byte

func init() {
	once.Do(func() {
		if err := yaml.Unmarshal(defaultConf, &Object); err != nil {
			slog.Error("Cannot load default wallet conf")
		}
		if err := envconfig.Process(context.Background(), &Object); err != nil {
			slog.Error("Cannot load env wallet conf")
		}
	})
}

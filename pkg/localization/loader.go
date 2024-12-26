package localization

import (
	_ "embed"
	"encoding/json"
	"log/slog"
	"sync"
)

//go:embed catalog.json
var catalogJson []byte

var once sync.Once

var object map[string]map[string]string

func init() {
	once.Do(func() {
		if err := json.Unmarshal(catalogJson, &object); err != nil {
			slog.Error("Error parsing catalog json: ", err)
		}
	})
}

func Get(lang string, key string) string {
	dict, ok := object[lang]
	if !ok {
		return key
	}
	val, ok := dict[key]
	if !ok {
		return key
	}
	return val
}

package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	DEFAULT_JOBHUNT_SERVICE_GRPC_ADDRESS     = "localhost:50051"
	DEFAULT_FETCH_RANDOM_SOURCE_COMPANY_SIZE = 2
)

// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var Koanf = koanf.New(".")

func Load() {
	// Load default values using the confmap provider.
	Koanf.Load(confmap.Provider(map[string]any{
		"jobhunt_service.grpc_address": DEFAULT_JOBHUNT_SERVICE_GRPC_ADDRESS,
		"workflows.populate_company.fetch_random_source_company_size": DEFAULT_FETCH_RANDOM_SOURCE_COMPANY_SIZE,
	}, "."), nil)

	// Load YAML config (TODO for deployment files).
	Koanf.Load(file.Provider("config.yml"), yaml.Parser())

	Koanf.Load(env.Provider(".", env.Opt{
		Prefix: "WORKFLOW__",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "WORKFLOW__")), "__", ".")

			// Transform the value into slices, if they contain spaces.
			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)
	log.Println("Loaded config variables.")
}

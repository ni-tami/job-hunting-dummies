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
	// Cockroach DB
	defaultCrdbDbName  = "job_hunt_service"
	defaultCrdbHost    = "localhost"
	defaultCrdbPort    = 26257
	defaultCrdbSslMode = "disable"

	// GQL
	defaultGQLPort = 8888

	// GRPC
	defaultGrpcPort = 50051
)

// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var Koanf = koanf.New(".")

func Load() {
	// Load default values using the confmap provider.
	Koanf.Load(confmap.Provider(map[string]any{
		"crdb.dbname":  defaultCrdbDbName,
		"crdb.host":    defaultCrdbHost,
		"crdb.port":    defaultCrdbPort,
		"crdb.sslmode": defaultCrdbSslMode,
		"gql.port":     defaultGQLPort,
		"grpc.port":    defaultGrpcPort,
	}, "."), nil)

	// Load YAML config (TODO for deployment files).
	Koanf.Load(file.Provider("config.yml"), yaml.Parser())

	Koanf.Load(env.Provider(".", env.Opt{
		Prefix: "SERVICE__",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "SERVICE__")), "__", ".")

			// Transform the value into slices, if they contain spaces.
			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)
	log.Println("Loaded config variables.")
}

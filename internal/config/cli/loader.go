package cli

import (
	"backend/internal/config"
	"flag"
	"log"
	"os"
)

type ConfigLoader struct{}

var isHelp = flag.Bool("help", false, "show this help message")
var pgDSN = flag.String("pgDSN", "", "set up DSN for postgres")
var httpPort = flag.Int("httpPort", 7880, "set up port for http handler")

func (ConfigLoader) Load() config.Config {
	flag.Parse()
	if *isHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *pgDSN == "" {
		log.Fatal("please, set Postgres DSN  (use flag -pgDSN)")
	}

	return config.Config{
		PostgresConfig: config.PostgresConfig{
			DSN: *pgDSN,
		},
		HTTPConfig: config.HTTPConfig{
			Port: *httpPort,
		},
	}
}

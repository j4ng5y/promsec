package promsec

import (
	"context"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"

	RootCLI = &cobra.Command{
		Use:     "promsec",
		Version: Version,
		Run:     func(ccmd *cobra.Command, args []string) {},
	}

	configFile string
)

func setLogLevel(l string) {
	switch strings.ToLower(l) {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Info().Msg("setting level to trace level")
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("setting level to debug level")
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("setting level to info level")
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Info().Msg("setting level to warn level")
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		log.Info().Msg("setting level to error level")
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		log.Info().Msg("setting level to fatal level")
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
		log.Info().Msg("setting level to panic level")
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msgf("%s is not a recognized log level, setting to info level", l)
	}
}

func rootFlags(ccmd *cobra.Command) {
	loglevel := ccmd.PersistentFlags().StringP("log-level", "v", "info", "The log level to use.")
	setLogLevel(*loglevel)

	ccmd.PersistentFlags().StringVarP(&configFile, "config-file", "f", "./config.yaml", "The config file to use to run the service.")
	if err := ccmd.MarkFlagFilename("config-file", "yaml", "yml", "json"); err != nil {
		log.Fatal().Err(err)
	}
}

func CLI() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	rootFlags(RootCLI)

	if err := RootCLI.Execute(); err != nil {
		log.Ctx(ctx).Err(err)
	}
}

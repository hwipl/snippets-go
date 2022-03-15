package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Trace().Msg("Hello, Trace")
	log.Debug().Msg("Hello, Debug")
	log.Info().Msg("Hello, Info")
	log.Warn().Msg("Hello, Warn")
	log.Error().Msg("Hello, Error")
	// log.Fatal().Msg("Bye, Fatal")
	// log.Panic().Msg("Bye, Panic")
	log.Info().Str("from", "zerolog").Str("to", "World").Msg("Hello")
	log.Error().Err(errors.New("something bad happened")).Msg("Oops")
	log.Fatal().Str("from", "zerolog").Str("to", "World").Msg("Bye")
}

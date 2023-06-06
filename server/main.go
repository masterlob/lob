package main

import (
	_ "embed"
	"os"
	"os/signal"
	"syscall"

	"github.com/masterlob/lob/server/infrastructure"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	_ = os.Setenv("TZ", "UTC")
	setupLogging()

	app := infrastructure.NewApplication()
	app.Start()
	<-shutdown
	_ = app.Stop()
}

func setupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Logger = log.With().Caller().Str("application", "lob-sample-app").Logger()
}

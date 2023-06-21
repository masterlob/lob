package infrastructure

import (
	"context"
	_ "embed"
	"errors"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	companiesapi "github.com/masterlob/lob/server/companies/api"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

const (
	serverReadHeaderTimeout = 5 * time.Second
	serverWriteTimeout      = 10 * time.Second
	serverReadTimeOut       = 120 * time.Second
)

type Application struct {
	*http.Server
}

type AppOptions struct {
	IsProduction bool
}

func NewApplication() *Application {
	return newApplication(AppOptions{IsProduction: true})
}

func newApplication(options AppOptions) *Application {
	router := httprouter.New()

	companiesapi.Register(router)

	return &Application{
		&http.Server{
			Addr:              getPort(options.IsProduction),
			ReadTimeout:       serverReadTimeOut,
			ReadHeaderTimeout: serverReadHeaderTimeout,
			WriteTimeout:      serverWriteTimeout,
			// FIXME: Setup CORS properly.
			// https://github.com/rs/cors/tree/master#parameters
			Handler: cors.Default().Handler(router),
		}}
}

func (a *Application) Start() {
	log.Info().Msg("Application starting.")

	a.startHTTPServer()
	log.Info().Msgf("Application started. Listening on port %s.", a.Addr)
}

func (a *Application) startHTTPServer() {
	go func() {
		if err := a.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info().Msgf("Server shutdown.")
			} else {
				log.Fatal().Msgf("Unexpected server error on server shutdown. Cause: %s\n", err.Error())
			}
		}
	}()
}

func (a *Application) Stop() error {
	log.Print("Application shutting down.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := a.Shutdown(ctx)

	log.Print("Application shutdown.")
	return err
}

func getPort(production bool) string {
	if production {
		return ":8080"
	}

	port, err := findRandomFreePort()
	if err != nil {
		log.Fatal().Msgf("Could not find a free port. Cause: %s\n", err.Error())
	}
	return ":" + strconv.Itoa(port)
}

func findRandomFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var listener *net.TCPListener
		if listener, err = net.ListenTCP("tcp", a); err == nil {
			defer listener.Close()
			return listener.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return 0, errors.New("could not find free port")
}

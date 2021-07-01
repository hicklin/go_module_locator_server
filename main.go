package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"moduleLocator/internal/views"
	"net/http"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	hostname := flag.String("hostname", "mycooldomain.com", "The URL of this server.")
	port := flag.Int("port", 80, "The server's hosting port.")
	repoLocation := flag.String("repoLocation", "", "The location where you repos can be found. Requires a final '/' if the location is a directory.")
	debug := flag.Bool("debug", false, "sets log level to debug.")
	human := flag.Bool("human", false, "sets log style to human friendly.")

	flag.Parse()

	// Set the logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if *human {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	views.Hostname = *hostname
	views.RepoLocation = *repoLocation

	log.Info().Str("hostname", *hostname).Int("port", *port).Msg("Starting Go module locator server.")

	http.HandleFunc("/", views.RepoPage)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
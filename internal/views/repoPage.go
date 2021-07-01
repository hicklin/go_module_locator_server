package views

import (
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

var (
	Hostname string
	RepoLocation string
)

type repoData struct {
	Hostname string
	ModuleName string
	RepoLocation string
}

// Serve an HTML page with the required meta tag that point go mod to the repo location.
func RepoPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pages/main/repoPage.html")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	data := repoData{
		Hostname: Hostname,
		ModuleName: r.URL.Path[1:],
		RepoLocation: RepoLocation,
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

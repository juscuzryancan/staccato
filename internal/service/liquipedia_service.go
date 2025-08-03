package service

import (
	"log"
	"net/http"
	"net/url"

	"github.com/juscuzryancan/staccato/internal/store"
)

type LiquipediaService struct {
	tournamentStore store.TournamentStore
	logger          *log.Logger
	client          *http.Client
}

const (
	baseURL = "https://liquipedia.net"
)

func NewLiquipediaService(tournamentStore store.TournamentStore, logger *log.Logger) *LiquipediaService {
	return &LiquipediaService{
		tournamentStore,
		logger,
		&http.Client{},
	}
}

type TournamentScraper interface{}

func (ls *LiquipediaService) GetTournament() (any, error) {
	req, err := http.NewRequest(
		"GET",
		"https://liquipedia.net/marvelrivals/api.php?action=query&format=json&titles=Cozyverse/Cash_Cup/9&prop=extracts|categories|pageprops|info|revisions|links|templates|images&explaintext=&exsectionformat=plain&rvprop=content&inprop=url&pllimit=max&tllimit=max&imlimit=max",
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "HumbleWinnings/0.1 (juscuzryancan@gmail.com)")

	data, err := ls.client.Do(req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ls *LiquipediaService) QueryLiquipedia() (any, error) {
	u, _ := url.Parse(baseURL)
	params := url.Values{}

	action := "query"
	format := "json"

	params.Add("format", format)

	// Valid Actions:
	params.Add("action", action)
	// params.Add("props", props.)

	data, err := ls.client.Get(u.String())
	if err != nil {
		return nil, err
	}

	return data, nil
}

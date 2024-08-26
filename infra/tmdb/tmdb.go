package tmdb

import (
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

type TMDB struct {
	Client *tmdb.TMDb
}

func NewTMDB() *TMDB {
	return &TMDB{}
}

func (tm *TMDB) Connect() {
	config := tmdb.Config{
		APIKey:   os.Getenv("TMDB_API_KEY"),
		Proxies:  nil,
		UseProxy: false,
	}

	tm.Client = tmdb.Init(config)
}

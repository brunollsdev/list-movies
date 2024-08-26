package facades

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/sirupsen/logrus"

	"github.com/brunollsdev/like-movies/application/dto"
	infraTMDB "github.com/brunollsdev/like-movies/infra/tmdb"
	"github.com/brunollsdev/like-movies/infra/utils"
)

type MoviesFacadeI interface {
	Search(name string) ([]dto.SearchResultDto, error)
	Description(id int) (*tmdb.Movie, error)
	FindById(id int) error
}

type MoviesFacade struct {
	tmdb *infraTMDB.TMDB
}

func NewMoviesFacade() *MoviesFacade {
	return &MoviesFacade{
		tmdb: infraTMDB.NewTMDB(),
	}
}

func (mf *MoviesFacade) FindById(id int) error {
	var opts = make(map[string]string)

	mf.tmdb.Connect()

	_, err := mf.tmdb.Client.GetMovieInfo(id, opts)

	if err != nil {
		logId := uuid.NewString()
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "facades",
			"method": "facades.movies_facade.find_by_id",
			"data":   err,
		}).Error(err.Error())

		return &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return nil
}

func (mf *MoviesFacade) Search(name string) ([]dto.SearchResultDto, error) {
	var opts = make(map[string]string)

	mf.tmdb.Connect()

	opts["language"] = "pt-BR"
	opts["include_adult"] = "false"
	opts["page_size"] = "1"

	result, err := mf.tmdb.Client.SearchMovie(name, opts)

	if err != nil {
		logId := uuid.NewString()
		utils.Log().WithFields(logrus.Fields{
			"id":     logId,
			"layer":  "facades",
			"method": "facades.movies_facade.search",
			"data":   err,
		}).Error(err.Error())

		return nil, &utils.LogDesc{
			LogId:     logId,
			EventDate: time.Now(),
		}
	}

	return parseSearchResult(result), nil
}

func (mf *MoviesFacade) Description(id int) (*tmdb.Movie, error) {
	var opts = make(map[string]string)

	mf.tmdb.Connect()

	movie, err := mf.tmdb.Client.GetMovieInfo(id, opts)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func parseSearchResult(result *tmdb.MovieSearchResults) []dto.SearchResultDto {
	var res []dto.SearchResultDto
	for _, item := range result.Results {
		res = append(res, dto.SearchResultDto{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Overview,
			ReleaseDate: item.ReleaseDate,
		})
	}

	return res
}

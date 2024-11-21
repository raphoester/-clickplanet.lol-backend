package clicks_controller

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
)

func New(
	tilesChecker domain.TilesChecker,
	countryChecker domain.CountryChecker,
	tilesStorage domain.TileStorage,
	answerer *httpserver.Answerer,
	reader httpserver.Reader,
) *Controller {
	return &Controller{
		answerer:       answerer,
		reader:         reader,
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		tilesStorage:   tilesStorage,
	}
}

type Controller struct {
	answerer *httpserver.Answerer
	reader   httpserver.Reader

	tilesChecker   domain.TilesChecker
	countryChecker domain.CountryChecker
	tilesStorage   domain.TileStorage
}

package clicks_controller

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
)

func New(
	tilesChecker game_map.TilesChecker,
	countryChecker game_map.CountryChecker,
	mapGetter game_map.Getter,
	tilesStorage game_map.TileStorage,
	answerer *httpserver.Answerer,
	reader httpserver.Reader,
) *Controller {
	return &Controller{
		answerer:       answerer,
		reader:         reader,
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
		tilesStorage:   tilesStorage,
	}
}

type Controller struct {
	answerer *httpserver.Answerer
	reader   httpserver.Reader
	
	tilesChecker   game_map.TilesChecker
	countryChecker game_map.CountryChecker
	mapGetter      game_map.Getter
	tilesStorage   game_map.TileStorage
}

package clicks_controller

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/click_handler_service"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
)

func New(
	clickHandlerService click_handler_service.IService,
	tilesChecker domain.TilesChecker,
	tilesStorage domain.TileStorage,
	answerer *httpserver.Answerer,
	reader httpserver.Reader,
) *Controller {
	return &Controller{
		answerer:            answerer,
		reader:              reader,
		clickHandlerService: clickHandlerService,
		tilesChecker:        tilesChecker,
		tilesStorage:        tilesStorage,
	}
}

type Controller struct {
	answerer *httpserver.Answerer
	reader   httpserver.Reader

	clickHandlerService click_handler_service.IService
	tilesChecker        domain.TilesChecker
	tilesStorage        domain.TileStorage
}

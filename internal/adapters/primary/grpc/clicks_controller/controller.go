package clicks_controller

import (
	"context"
	"errors"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
)

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
	mapGetter MapGetter,
) *Controller {
	return &Controller{
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
	}
}

type Controller struct {
	tilesChecker   TilesChecker
	countryChecker CountryChecker
	mapGetter      MapGetter
	*clicksv1.UnimplementedClicksServer
}

type TilesChecker interface {
	Check(tile string) bool
}

type CountryChecker interface {
	Check(country string) bool
}

type MapGetter interface {
	GetMap() *game_map.GameMap
}

func (c *Controller) HandleClick(
	ctx context.Context,
	req *clicksv1.ClickRequest,
) (*clicksv1.ClickResponse, error) {

	if !c.countryChecker.Check(req.GetCountryId()) {
		return nil, errors.New("invalid country")
	}

	if !c.tilesChecker.Check(req.GetTileId()) {
		return nil, errors.New("invalid tile")
	}

	return &clicksv1.ClickResponse{}, nil
}

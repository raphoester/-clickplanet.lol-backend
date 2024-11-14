package clicks_controller

import (
	"context"
	"errors"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
) *Controller {
	return &Controller{
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
	}
}

type Controller struct {
	tilesChecker   TilesChecker
	countryChecker CountryChecker
	*clicksv1.UnimplementedClicksServer
}

type TilesChecker interface {
	Check(tile string) bool
}

type CountryChecker interface {
	Check(country string) bool
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

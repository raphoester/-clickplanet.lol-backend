package click_handler_service

import (
	"context"
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
)

func New(
	tilesChecker domain.TilesChecker,
	tileStorage domain.TileStorage,
	countryChecker domain.CountryChecker,
) *Service {
	return &Service{
		tilesChecker:   tilesChecker,
		tileStorage:    tileStorage,
		countryChecker: countryChecker,
	}
}

type IService interface {
	HandleClick(ctx context.Context, tileId uint32, countryID string) error
}

type Service struct {
	tilesChecker   domain.TilesChecker
	tileStorage    domain.TileStorage
	countryChecker domain.CountryChecker
}

func (s *Service) HandleClick(ctx context.Context, tileId uint32, countryID string) error {
	if !s.countryChecker.CheckCountry(countryID) {
		return fmt.Errorf("invalid country code %q", countryID)
	}

	if !s.tilesChecker.CheckTile(tileId) {
		return fmt.Errorf("invalid tile id %q", tileId)
	}

	if err := s.tileStorage.Set(ctx, tileId, countryID); err != nil {
		return fmt.Errorf("failed to set tile: %w", err)
	}

	return nil
}

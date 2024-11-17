package game_map

import (
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"
)

type Getter interface {
	GetMap() *GameMap
}

type GameMapConfig struct {
	RowsCount              int
	HorizontalDensityIndex int
	RegionsCount           int
}

func (g GameMapConfig) Validate() error {
	if g.RowsCount < 1 {
		return fmt.Errorf("rows count must be at least 1, was %d", g.RowsCount)
	}

	if g.HorizontalDensityIndex < 1 {
		return fmt.Errorf("horizontal density index must be at least 1, was %d", g.HorizontalDensityIndex)
	}

	if g.RegionsCount < 1 {
		return fmt.Errorf("regions count must be at least 1, was %d", g.RegionsCount)
	}

	return nil
}

func Generate(config GameMapConfig) (*GameMap, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	tiles := generateTiles(
		config.RowsCount,
		config.HorizontalDensityIndex)
	epicenters := GenerateFibonacciSphere(config.RegionsCount)
	geodesicEpicenters := make([]coordinates.Geodesic, 0, len(epicenters))
	for _, epicenter := range epicenters {
		// points should already be normalized
		geodesicEpicenters = append(geodesicEpicenters, epicenter.ToGeodesic())
	}
	return &GameMap{
		Regions: assignTilesToRegions(geodesicEpicenters, tiles),
		Tiles:   tiles,
	}, nil
}

type GameMap struct {
	Regions []Region
	Tiles   []Tile
}

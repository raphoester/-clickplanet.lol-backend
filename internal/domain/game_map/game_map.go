package game_map

import "github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"

type GameMapConfig struct {
	RowsCount              int
	HorizontalDensityIndex int
	RegionsCount           int
}

func Generate(config GameMapConfig) *GameMap {
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
	}
}

type GameMap struct {
	Regions []Region
	Tiles   []Tile
}

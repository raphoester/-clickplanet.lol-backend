package game_map

import (
	"math"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"
)

type Region struct {
	epicenter coordinates.Geodesic
	tiles     []Tile
}

func (r *Region) Epicenter() coordinates.Geodesic {
	return r.epicenter
}

func (r *Region) Tiles() []Tile {
	return r.tiles
}

func assignTilesToRegions(
	epicenters []coordinates.Geodesic,
	tiles []Tile,
) []Region {
	regions := make([]Region, len(epicenters))
	for i, epicenter := range epicenters {
		regions[i] = Region{epicenter: epicenter}
	}

	for _, tile := range tiles {
		minDistance := math.Inf(1)
		minIndex := -1

		for i, region := range regions {
			distance := region.epicenter.HaversineDistanceTo(tile.Epicenter())
			if distance < minDistance {
				minDistance = distance
				minIndex = i
			}
		}

		regions[minIndex].tiles = append(regions[minIndex].tiles, tile)
	}

	return regions
}

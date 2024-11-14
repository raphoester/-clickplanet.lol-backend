package game_map

import (
	"fmt"
	"math"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"
)

type Tile struct {
	southWest coordinates.Geodesic
	southEast coordinates.Geodesic
	northEast coordinates.Geodesic
	northWest coordinates.Geodesic
	epicenter *coordinates.Geodesic
}

func NewTile(southWestLon, southWestLat, lonStep, latStep float64) *Tile {
	southWest := coordinates.MustNewGeodesic(southWestLon, southWestLat)
	southEast := coordinates.MustNewGeodesic(southWestLon+lonStep, southWestLat)
	northEast := coordinates.MustNewGeodesic(southWestLon+lonStep, southWestLat+latStep)
	northWest := coordinates.MustNewGeodesic(southWestLon, southWestLat+latStep)

	return &Tile{
		southWest: southWest,
		southEast: southEast,
		northEast: northEast,
		northWest: northWest,
		epicenter: nil,
	}
}

func (t *Tile) GetBoundaries() []coordinates.Geodesic {
	return []coordinates.Geodesic{t.southWest, t.southEast, t.northEast, t.northWest}
}

func (t *Tile) Epicenter() coordinates.Geodesic {
	if t.epicenter == nil {
		latitude := (t.southWest.Latitude() + t.southEast.Latitude() + t.northEast.Latitude() + t.northWest.Latitude()) / 4
		longitude := (t.southWest.Longitude() + t.southEast.Longitude() + t.northEast.Longitude() + t.northWest.Longitude()) / 4
		*t.epicenter = coordinates.MustNewGeodesic(longitude, latitude)
	}
	return *t.epicenter
}

func (t *Tile) ID() string {
	return fmt.Sprintf("%f;%f", t.southWest.Latitude(), t.southWest.Longitude())
}

func absDiff(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

const (
	startLat = -90
	endLat   = 90
	startLon = -180
	endLon   = 180
)

func generateTiles(rows int, density int) []Tile {
	latitudeStep := absDiff(startLat, endLat) / float64(rows)
	tiles := make([]Tile, 0)

	for i := 0; i < rows; i++ {
		southBorderLatitude := startLat + (float64(i) * latitudeStep)
		absCenterLatitude := math.Abs(southBorderLatitude + (latitudeStep / 2))
		squaresCount := int(
			math.Round(180-(0.00023*math.Pow(absCenterLatitude, 3)))) * density

		longitudeStep := absDiff(startLon, endLon) / float64(squaresCount)
		for j := 0; j < squaresCount; j++ {
			tile := NewTile(startLon+(float64(j)*longitudeStep), southBorderLatitude, longitudeStep, latitudeStep)
			tiles = append(tiles, *tile)
		}
	}
	return tiles
}

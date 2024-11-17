package game_map

import (
	"fmt"
	"math"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/basicutil"
)

type Tile struct {
	southWest coordinates.Geodesic
	northEast coordinates.Geodesic
	epicenter *coordinates.Geodesic
}

func MustNewTile(
	southLat, westLon,
	latStep, lonStep float64,
) *Tile {
	southWest := coordinates.MustNewGeodesic(southLat, westLon)
	northEast := coordinates.MustNewGeodesic(southLat+latStep, westLon+lonStep)
	return &Tile{
		southWest: southWest,
		northEast: northEast,
	}
}

func (t *Tile) GetBoundaries() (coordinates.Geodesic, coordinates.Geodesic) {
	return t.southWest, t.northEast
}

func (t *Tile) Epicenter() coordinates.Geodesic {
	if t.epicenter == nil {
		latitude := (t.southWest.Latitude() + t.northEast.Latitude()) / 2
		longitude := (t.southWest.Longitude() + t.northEast.Longitude()) / 2
		t.epicenter = basicutil.Pointer(coordinates.MustNewGeodesic(latitude, longitude))
	}
	return *t.epicenter
}

func (t *Tile) ID() string {
	return fmt.Sprintf("%0.3f;%0.3f", t.southWest.Latitude(), t.southWest.Longitude())
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
			tile := MustNewTile(southBorderLatitude, startLon+(float64(j)*longitudeStep), latitudeStep, longitudeStep)
			tiles = append(tiles, *tile)
		}
	}
	return tiles
}

type TilesChecker interface {
	CheckTile(tile string) bool
}

type TileStorage interface {
	Set(tile string, value string)
	Get() map[string]string
}

type TileReporter interface {
	Subscribe() <-chan TileUpdate
}

type TileUpdate struct {
	Tile  string
	Value string
}

type CountryChecker interface {
	CheckCountry(country string) bool
}

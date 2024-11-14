package coordinates

import (
	"errors"
	"math"
)

func MustNewGeodesic(latitude, longitude float64) Geodesic {
	c, err := NewGeodesic(latitude, longitude)
	if err != nil {
		panic(err)
	}
	return *c
}

func NewGeodesic(latitude, longitude float64) (*Geodesic, error) {
	if latitude < -90 || latitude > 90 {
		return nil, errors.New("latitude must be between -90 and 90")
	}

	if longitude < -180 || longitude > 180 {
		return nil, errors.New("longitude must be between -180 and 180")
	}

	return &Geodesic{
		latitude:  latitude,
		longitude: longitude,
	}, nil
}

type Geodesic struct {
	latitude  float64
	longitude float64
}

func (g Geodesic) Latitude() float64 {
	return g.latitude
}

func (g Geodesic) Longitude() float64 {
	return g.longitude
}

const earthRadius = 6371

func (c Geodesic) HaversineDistanceTo(other Geodesic) float64 {
	dLat := (other.latitude - c.latitude) * math.Pi / 180
	dLon := (other.longitude - c.longitude) * math.Pi / 180
	lat1 := c.latitude * math.Pi / 180
	lat2 := other.latitude * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	cVal := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * cVal // Distance en km
}

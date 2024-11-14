package coordinates

import "math"

type Cartesian3 struct {
	X float64
	Y float64
	Z float64
}

func (c Cartesian3) IsNormalized() bool {
	return math.Abs(c.X*c.X+c.Y*c.Y+c.Z*c.Z-1) < 0.0001
}

func (c Cartesian3) Normalize() Cartesian3 {
	length := math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z)
	return Cartesian3{X: c.X / length, Y: c.Y / length, Z: c.Z / length}
}

func (c Cartesian3) ToGeodesic() Geodesic {
	lon := math.Atan2(c.Y, c.X) * 180 / math.Pi
	lat := math.Asin(c.Z) * 180 / math.Pi
	return Geodesic{longitude: lon, latitude: lat}
}

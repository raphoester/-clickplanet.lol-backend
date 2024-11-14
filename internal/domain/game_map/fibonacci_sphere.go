package game_map

import (
	"math"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/coordinates"
)

func GenerateFibonacciSphere(pointsCount int) []coordinates.Cartesian3 {
	points := make([]coordinates.Cartesian3, 0, pointsCount)
	phi := (1 + math.Sqrt(5)) / 2

	for i := 0; i < pointsCount; i++ {
		y := 1 - (float64(i)/float64(pointsCount-1))*2
		radius := math.Sqrt(1 - y*y)
		theta := phi * float64(i)

		x := math.Cos(theta) * radius
		z := math.Sin(theta) * radius

		points = append(points, coordinates.Cartesian3{X: x, Y: y, Z: z})
	}

	return points
}

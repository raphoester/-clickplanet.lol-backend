package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetMapDensity(w http.ResponseWriter, _ *http.Request) {
	maxIndex := c.tilesChecker.MaxIndex()
	c.answerer.Data(w, &clicksv1.MapDensityResponse{
		Density: maxIndex,
	})
}

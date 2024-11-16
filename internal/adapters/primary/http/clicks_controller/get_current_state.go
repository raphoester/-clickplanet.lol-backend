package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetCurrentState(w http.ResponseWriter, _ *http.Request) {
	tiles := c.tilesStorage.Get()
	response := &clicksv1.State{Data: tiles}
	c.answerWithData(w, response)
}

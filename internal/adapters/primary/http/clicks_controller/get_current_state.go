package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetOwnerships(w http.ResponseWriter, _ *http.Request) {
	tiles := c.tilesStorage.Get()
	response := &clicksv1.Ownerships{Bindings: tiles}
	c.answerer.Data(w, response)
}

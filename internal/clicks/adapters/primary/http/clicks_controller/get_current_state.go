package clicks_controller

import (
	"fmt"
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetOwnerships(w http.ResponseWriter, r *http.Request) {
	tiles, err := c.tilesStorage.GetFullState(r.Context())
	if err != nil {
		c.answerer.Err(w, fmt.Errorf("failed to get full state: %w", err),
			"internal error", http.StatusInternalServerError)
		return
	}

	response := &clicksv1.Ownerships{Bindings: tiles}
	c.answerer.Data(w, response)
}

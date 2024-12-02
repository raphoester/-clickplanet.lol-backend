package clicks_controller

import (
	"fmt"
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) HandleClick(w http.ResponseWriter, r *http.Request) {
	req := &clicksv1.ClickRequest{}
	if err := c.reader.Read(r, req); err != nil {
		c.answerer.Err(w,
			fmt.Errorf("failed reading json req: %w", err),
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	if err := c.clickHandlerService.HandleClick(r.Context(), req.GetTileId(), req.GetCountryId()); err != nil {
		c.answerer.Err(w,
			fmt.Errorf("failed to handle click: %w", err),
			"internal error",
			http.StatusInternalServerError,
		)
		return
	}

	c.answerer.Empty(w)
}

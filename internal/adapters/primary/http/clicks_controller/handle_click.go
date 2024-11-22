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

	if !c.countryChecker.CheckCountry(req.GetCountryId()) {
		c.answerer.Err(w,
			fmt.Errorf("invalid country code %q", req.GetCountryId()),
			"invalid country",
			http.StatusBadRequest,
		)
		return
	}

	if !c.tilesChecker.CheckTile(req.GetTileId()) {
		c.answerer.Err(w,
			fmt.Errorf("invalid tile id %q", req.GetTileId()),
			"invalid tile",
			http.StatusBadRequest,
		)
		return
	}

	if err := c.tilesStorage.Set(r.Context(), req.GetTileId(), req.GetCountryId()); err != nil {
		c.answerer.Err(w, fmt.Errorf("failed to set tile: %w", err),
			"internal error", http.StatusInternalServerError)
		return
	}

	c.answerer.Empty(w)
}

package clicks_controller

import (
	"fmt"
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) HandleClick(w http.ResponseWriter, r *http.Request) {
	req := &clicksv1.ClickRequest{}

	if err := readExchangeMsg(r, req); err != nil {
		c.answerWithErr(w,
			fmt.Errorf("failed reading json req: %w", err),
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	if !c.countryChecker.Check(req.GetCountryId()) {
		c.answerWithErr(w,
			fmt.Errorf("invalid country code %q", req.GetCountryId()),
			"invalid country",
			http.StatusBadRequest,
		)
		return
	}

	if !c.tilesChecker.Check(req.GetTileId()) {
		c.answerWithErr(w,
			fmt.Errorf("invalid tile id %q", req.GetTileId()),
			"invalid tile",
			http.StatusBadRequest,
		)
		return
	}

	c.tilesStorage.Set(req.GetTileId(), req.GetCountryId())
	c.answerEmpty(w)
}

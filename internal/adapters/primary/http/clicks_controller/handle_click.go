package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
)

func (c *Controller) HandleClick(w http.ResponseWriter, r *http.Request) {
	req := &clicksv1.ClickRequest{}

	if err := readExchangeMsg(r, req); err != nil {
		c.logger.Info("failed to read request", lf.Err(err))
		answerWithErr(w, "invalid request", http.StatusBadRequest)
		return
	}

	if !c.countryChecker.Check(req.GetCountryId()) {
		answerWithErr(w, "invalid country", http.StatusBadRequest)
		return
	}

	if !c.tilesChecker.Check(req.GetTileId()) {
		c.logger.Info("invalid tile", lf.String("tile", req.GetTileId()))
		answerWithErr(w, "invalid tile", http.StatusBadRequest)
		return
	}

	c.tilesStorage.Set(req.GetTileId(), req.GetCountryId())
	answerEmpty(w)
}

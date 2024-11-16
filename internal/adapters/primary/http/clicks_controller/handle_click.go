package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) HandleClick(w http.ResponseWriter, r *http.Request) {
	req := &clicksv1.ClickRequest{}
	if err := readExchangeMsg(r, req); err != nil {
		answerWithErr(w, "invalid request", http.StatusBadRequest)
		return
	}

	if !c.countryChecker.Check(req.GetCountryId()) {
		answerWithErr(w, "invalid country", http.StatusBadRequest)
		return
	}

	if !c.tilesChecker.Check(req.GetTileId()) {
		answerWithErr(w, "invalid tile", http.StatusBadRequest)
		return
	}

	c.tilesStorage.Set(req.GetTileId(), req.GetCountryId())
}

package clicks_controller

import (
	"encoding/json"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
)

type TilesChecker interface {
	Check(tile string) bool
}

type CountryChecker interface {
	Check(country string) bool
}

type MapGetter interface {
	GetMap() *game_map.GameMap
}

type HandleClickRequest struct {
	CountryID string
	TileID    string
}

func (c *Controller) HandleClick(w http.ResponseWriter, r *http.Request) {
	var req HandleClickRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}

	if !c.countryChecker.Check(req.CountryID) {
		http.Error(w, "invalid country", http.StatusBadRequest)
		return
	}

	if !c.tilesChecker.Check(req.TileID) {
		http.Error(w, "invalid tile", http.StatusBadRequest)
		return
	}
}

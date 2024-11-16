package clicks_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"google.golang.org/protobuf/proto"
)

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
	mapGetter MapGetter,
	tilesStorage TileStorage,
) *Controller {
	return &Controller{
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
		tilesStorage:   tilesStorage,
	}
}

type Controller struct {
	tilesChecker   TilesChecker
	countryChecker CountryChecker
	mapGetter      MapGetter
	tilesStorage   TileStorage
}

type TilesChecker interface {
	Check(tile string) bool
	//GetAll() map[string]string
}

type CountryChecker interface {
	Check(country string) bool
}

type TileStorage interface {
	Set(tile string, value string)
	Get() map[string]string
}

type MapGetter interface {
	GetMap() *game_map.GameMap
}

type ExchangeFormat struct {
	Data []byte `json:"data"`
}

type ErrorFormat struct {
	Cause string `json:"cause"`
}

func answerWithErr(w http.ResponseWriter, cause string, status int) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorFormat{Cause: cause})
}

func answerWithData(w http.ResponseWriter, protoMsg proto.Message) {
	protoBytes, err := proto.Marshal(protoMsg)
	if err != nil {
		answerWithErr(w, "internal error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ExchangeFormat{Data: protoBytes}); err != nil {
		answerWithErr(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func readExchangeMsg(r *http.Request, protoMsg proto.Message) error {
	reqContent := &ExchangeFormat{}
	if err := json.NewDecoder(r.Body).Decode(reqContent); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	if err := proto.Unmarshal(reqContent.Data, protoMsg); err != nil {
		return fmt.Errorf("invalid proto data: %w", err)
	}

	return nil
}

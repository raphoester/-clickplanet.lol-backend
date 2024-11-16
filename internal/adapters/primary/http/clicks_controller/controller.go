package clicks_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
	"google.golang.org/protobuf/proto"
)

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
	mapGetter MapGetter,
	tilesStorage TileStorage,
	logger logging.Logger,
) *Controller {
	return &Controller{
		logger:         logger,
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
		tilesStorage:   tilesStorage,
	}
}

type Controller struct {
	logger         logging.Logger
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

func (c *Controller) answerWithErr(w http.ResponseWriter, logErr error, cause string, status int) {
	w.WriteHeader(status)
	c.logger.Error("error occurred", lf.Err(logErr))
	_ = json.NewEncoder(w).Encode(ErrorFormat{Cause: cause})
}

func (c *Controller) answerEmpty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

func (c *Controller) answerWithData(w http.ResponseWriter, protoMsg proto.Message) {
	protoBytes, err := proto.Marshal(protoMsg)
	if err != nil {
		c.answerWithErr(w,
			fmt.Errorf("failed marshalling proto msg: %w", err),
			"internal error",
			http.StatusInternalServerError,
		)
		return
	}

	if err := json.NewEncoder(w).Encode(ExchangeFormat{Data: protoBytes}); err != nil {
		c.logger.Error("failed to encode response", lf.Err(err))
		c.answerWithErr(w,
			fmt.Errorf("failed to encode wrapping json: %w", err),
			"internal error",
			http.StatusInternalServerError,
		)
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

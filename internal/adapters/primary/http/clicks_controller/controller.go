package clicks_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
	mapGetter MapGetter,
) *Controller {
	return &Controller{
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
	}
}

type Controller struct {
	tilesChecker   TilesChecker
	countryChecker CountryChecker
	mapGetter      MapGetter
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

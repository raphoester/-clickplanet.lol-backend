package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
	"google.golang.org/protobuf/proto"
)

type AnswerMode int

const (
	AnswerModeBinary AnswerMode = iota
	AnswerModeJSON
)

func NewAnswerer(
	logger logging.Logger,
	mode AnswerMode,
) *Answerer {
	answerMsg := answerJSON
	if mode == AnswerModeBinary {
		answerMsg = answerBinary
	}

	return &Answerer{
		logger:    logger,
		mode:      mode,
		answerMsg: answerMsg,
	}
}

type Answerer struct {
	logger    logging.Logger
	mode      AnswerMode
	answerMsg func(a *Answerer, w http.ResponseWriter, protoMsg proto.Message)
}

type ExchangeFormat struct {
	Data []byte `json:"data"`
}

type ErrorFormat struct {
	Cause string `json:"cause"`
}

func (a *Answerer) Err(w http.ResponseWriter, logErr error, cause string, status int) {
	w.WriteHeader(status)
	a.logger.Error("error occurred", lf.Err(logErr))
	_ = json.NewEncoder(w).Encode(ErrorFormat{Cause: cause})
}

func (a *Answerer) Empty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

func (a *Answerer) Data(w http.ResponseWriter, protoMsg proto.Message) {
	a.answerMsg(a, w, protoMsg)
}

func answerBinary(a *Answerer, w http.ResponseWriter, protoMsg proto.Message) {
	w.WriteHeader(http.StatusOK)
	protoBytes, err := proto.Marshal(protoMsg)
	if err != nil {
		a.Err(w, fmt.Errorf("failed marshalling proto msg: %w", err),
			"internal error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ExchangeFormat{Data: protoBytes}); err != nil {
		a.Err(w, fmt.Errorf("failed encoding exchange format: %w", err),
			"internal error", http.StatusInternalServerError)
	}
}

func answerJSON(a *Answerer, w http.ResponseWriter, protoMsg proto.Message) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(protoMsg); err != nil {
		a.Err(w, fmt.Errorf("failed to encode json: %w", err),
			"internal error", http.StatusInternalServerError)
	}
}

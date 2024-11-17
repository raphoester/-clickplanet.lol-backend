package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Reader interface {
	Read(r *http.Request, message proto.Message) error
}

type ProtoReader struct{}

func (ProtoReader) Read(r *http.Request, protoMsg proto.Message) error {
	reqContent := &ExchangeFormat{}
	if err := json.NewDecoder(r.Body).Decode(reqContent); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	if err := proto.Unmarshal(reqContent.Data, protoMsg); err != nil {
		return fmt.Errorf("invalid proto data: %w", err)
	}

	return nil
}

type JSONReader struct{}

func (JSONReader) Read(r *http.Request, protoMsg proto.Message) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed reading body: %w", err)
	}
	if err := protojson.Unmarshal(body, protoMsg); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	return nil
}

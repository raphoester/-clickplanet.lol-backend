package clicks_service_retriever

import (
	"context"

	"github.com/raphoester/clickplanet.lol-backend/internal/bookkeeping/runner"
)

func New() *Retriever {
	return &Retriever{}
}

type Retriever struct{}

func (r *Retriever) Retrieve(ctx context.Context) (*runner.Update, error) {
	return &runner.Update{}, nil
}

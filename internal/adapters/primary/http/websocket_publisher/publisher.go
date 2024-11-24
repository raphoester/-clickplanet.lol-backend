package websocket_publisher

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
	"google.golang.org/protobuf/proto"
)

func New(
	updates <-chan domain.TileUpdate,
	answerer *httpserver.Answerer,
) *Publisher {
	return &Publisher{
		clients:  make(map[*websocket.Conn]*clientMD),
		updates:  updates,
		answerer: answerer,
	}
}

type Publisher struct {
	mu       sync.RWMutex
	clients  map[*websocket.Conn]*clientMD
	updates  <-chan domain.TileUpdate
	answerer *httpserver.Answerer
}

type clientMD struct {
}

func (p *Publisher) Run() {
	for update := range p.updates {
		bin, err := proto.Marshal(&clicksv1.TileUpdate{
			TileId:            update.Tile,
			CountryId:         update.Value,
			PreviousCountryId: update.Previous,
		})

		if err != nil {
			continue
		}

		p.mu.RLock()
		for client := range p.clients {
			err := client.Write(context.Background(), websocket.MessageBinary, bin)
			if err == nil {
				continue
			}
		}
		p.mu.RUnlock()
	}
}

func (p *Publisher) Subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns:     []string{"*"},
		InsecureSkipVerify: true,
	})

	if err != nil {
		p.answerer.Err(w,
			fmt.Errorf("failed to accept websocket connection: %w", err),
			"cannot accept websocket connection",
			http.StatusInternalServerError,
		)
		return
	}

	p.mu.Lock()
	p.clients[conn] = &clientMD{}
	p.mu.Unlock()
}

package game_map_test

import (
	"testing"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("should generate a game map with a basic config", func(t *testing.T) {
		gameMap := game_map.Generate(game_map.GameMapConfig{
			RowsCount:              30,
			HorizontalDensityIndex: 2,
			RegionsCount:           10,
		})

		assert.Len(t, gameMap.Regions, 10)
		//assert.Len(t, gameMap.Tiles, )
	})
}

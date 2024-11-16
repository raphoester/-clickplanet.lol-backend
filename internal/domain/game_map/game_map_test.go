package game_map_test

import (
	"testing"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("should generate a game map with a basic config", func(t *testing.T) {
		gameMap, err := game_map.Generate(game_map.GameMapConfig{
			RowsCount:              30,
			HorizontalDensityIndex: 2,
			RegionsCount:           10,
		})

		require.NoError(t, err)
		assert.Len(t, gameMap.Regions, 10)
		var tilesCount int
		for _, region := range gameMap.Regions {
			tilesCount += len(region.Tiles())
		}
		t.Logf("tiles count: %d", tilesCount)
	})
}

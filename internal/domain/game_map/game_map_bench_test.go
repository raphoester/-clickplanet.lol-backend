package game_map

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkGenerate(b *testing.B) {
	for _, tc := range []struct {
		rows              int
		horizontalDensity int
		regions           int
	}{
		{300, 1, 150},
		{500, 1, 300},
		{1000, 2, 500},
		{2000, 3, 1000},
	} {
		b.Run(fmt.Sprintf(
			"rows: %d, horizontalDensity: %d, regions: %d",
			tc.rows, tc.horizontalDensity, tc.regions),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := Generate(GameMapConfig{
						RowsCount:              tc.rows,
						HorizontalDensityIndex: tc.horizontalDensity,
						RegionsCount:           tc.regions,
					})
					require.NoError(b, err)
				}
			},
		)
	}
}

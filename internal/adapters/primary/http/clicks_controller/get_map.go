package clicks_controller

import (
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetMap(w http.ResponseWriter, _ *http.Request) {
	theMap := c.mapGetter.GetMap()
	// TODO: avoid re-mapping on each call

	protoRegions := make([]*clicksv1.Region, 0, len(theMap.Regions))
	for _, region := range theMap.Regions {
		tiles := region.Tiles()
		protoTiles := make([]*clicksv1.Tile, 0, len(tiles))
		for _, tile := range tiles {
			southWest, northEast := tile.GetBoundaries()
			protoTiles = append(protoTiles, &clicksv1.Tile{
				SouthWest: &clicksv1.GeodesicCoordinates{
					Lat: southWest.Latitude(),
					Lon: southWest.Longitude(),
				},
				NorthEast: &clicksv1.GeodesicCoordinates{
					Lat: northEast.Latitude(),
					Lon: northEast.Longitude(),
				},
				Id: tile.ID(),
			})
		}
		epicenter := region.Epicenter()
		protoRegions = append(protoRegions, &clicksv1.Region{
			Epicenter: &clicksv1.GeodesicCoordinates{
				Lat: epicenter.Latitude(),
				Lon: epicenter.Longitude(),
			},
			Tiles: protoTiles,
		})
	}

	c.answerer.Data(w, &clicksv1.Map{
		Regions: protoRegions,
	})
}

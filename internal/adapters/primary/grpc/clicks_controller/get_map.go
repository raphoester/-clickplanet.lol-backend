package clicks_controller

import (
	"context"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
)

func (c *Controller) GetMap(
	ctx context.Context,
	req *clicksv1.GetMapRequest,
) (*clicksv1.GetMapResponse, error) {
	theMap := c.mapGetter.GetMap()
	
	// TODO: avoid re-mapping on each call

	protoRegions := make([]*clicksv1.Region, 0, len(theMap.Regions))
	for _, region := range theMap.Regions {
		tiles := region.GetTiles()
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
		protoRegions = append(protoRegions, &clicksv1.Region{
			Tiles: protoTiles,
		})
	}

	return &clicksv1.GetMapResponse{
		Regions: protoRegions,
	}, nil
}

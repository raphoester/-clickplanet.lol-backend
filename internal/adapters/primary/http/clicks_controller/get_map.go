package clicks_controller

import (
	"encoding/json"
	"net/http"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
	"google.golang.org/protobuf/proto"
)

func (c *Controller) GetMap(w http.ResponseWriter, r *http.Request) {
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

	protoBytes, err := proto.Marshal(&clicksv1.Map{Regions: protoRegions})
	if err != nil {
		http.Error(w, "failed to marshal map", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string][]byte{"data": protoBytes})

	return
}

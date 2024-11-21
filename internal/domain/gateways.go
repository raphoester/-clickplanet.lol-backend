package domain

type TilesChecker interface {
	CheckTile(tile uint32) bool
	MaxIndex() uint32
}

type TileStorage interface {
	Set(tile uint32, value string)
	Get() map[uint32]string
}

type TileReporter interface {
	Subscribe() <-chan TileUpdate
}

type CountryChecker interface {
	CheckCountry(country string) bool
}

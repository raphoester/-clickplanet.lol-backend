package clicks_controller

func New(
	tilesChecker TilesChecker,
	countryChecker CountryChecker,
	mapGetter MapGetter,
) *Controller {
	return &Controller{
		tilesChecker:   tilesChecker,
		countryChecker: countryChecker,
		mapGetter:      mapGetter,
	}
}

type Controller struct {
	tilesChecker   TilesChecker
	countryChecker CountryChecker
	mapGetter      MapGetter
}

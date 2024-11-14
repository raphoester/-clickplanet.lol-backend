package in_memory_country_checker

type Checker struct {
	countries map[string]struct{}
}

func New() *Checker {
	return &Checker{
		countries: countries,
	}
}

func (c *Checker) Check(s string) bool {
	_, ok := c.countries[s]
	return ok
}

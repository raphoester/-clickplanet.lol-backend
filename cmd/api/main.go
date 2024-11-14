package api

import "github.com/raphoester/clickplanet.lol-backend/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

package particles

import (
	"math/rand"
	"project-particles/config"
)

func randomPos() (float64, float64) {
	x := (-10) + rand.Float64()*(float64(config.General.WindowSizeX)+10)
	y := (-10) + rand.Float64()*(float64(config.General.WindowSizeY)+10)
	return x, y
}

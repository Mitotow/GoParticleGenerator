package particles

import (
	"math/rand"
	"project-particles/config"
)

func randomPos() (float64, float64) {
	return rand.Float64() * float64(config.General.WindowSizeX), rand.Float64() * float64(config.General.WindowSizeY)
}

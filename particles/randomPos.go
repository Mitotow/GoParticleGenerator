package particles

import (
	"math/rand"
	"project-particles/config"
)

func randomPos() (float64, float64) {
	var wX, wY = config.General.WindowSizeX, config.General.WindowSizeY
	x, y := rand.Float64()*float64(wX), rand.Float64()*float64(wY) // Create random coordonates
	return x, y
}

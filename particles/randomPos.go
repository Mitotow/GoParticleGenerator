package particles

import (
	"math/rand"
)

func randomPos() (float64, float64) {
	wX, wY := getWindowSize()                                      // Get the actual window size
	x, y := rand.Float64()*float64(wX), rand.Float64()*float64(wY) // Create random coordonates
	return x, y
}

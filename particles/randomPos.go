package particles

import (
	"math/rand"
)

var wX, wY = getWindowSize() // Get the actual window size (in this case configuration, it's the window size at the start of the app)

func randomPos() (float64, float64) {
	x, y := rand.Float64()*float64(wX), rand.Float64()*float64(wY) // Create random coordonates
	return x, y
}

package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"project-particles/config"
)

func (p Particle) InScreen() bool {
	var wX, wY int
	if ebiten.IsFullscreen() {
		wX, wY = ebiten.ScreenSizeInFullscreen()
	} else {
		wX, wY = config.General.WindowSizeX, config.General.WindowSizeY
	}

	if (-10 < p.PositionX && p.PositionX < float64(wX)+10) || (-10 < p.PositionY && p.PositionY < float64(wY)+10) {
		return true
	}
	return false
}

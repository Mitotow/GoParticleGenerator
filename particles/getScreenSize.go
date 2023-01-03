package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"project-particles/config"
)

func getScreenSize() (int, int) {
	if ebiten.IsFullscreen() {
		return ebiten.ScreenSizeInFullscreen()
	}
	return config.General.WindowSizeX, config.General.WindowSizeY
}

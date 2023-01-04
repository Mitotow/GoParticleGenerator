package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"project-particles/config"
)

func getWindowSize() (int, int) {
	if ebiten.IsFullscreen() {
		// Window is in fullscreen
		return ebiten.ScreenSizeInFullscreen() // Return the screen size of the user
	}
	return config.General.WindowSizeX, config.General.WindowSizeY // Return by default the Window size of the config.json
}
